package testingcommon

import (
	"database/sql"
	"fmt"

	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/jinzhu/gorm"
)

//Referred from https://medium.com/@jarifibrahim/using-gorm-hooks-to-clean-up-test-fixtures-in-golang-99b0fcb04354
func DeleteCreatedEntities() func() {
	db := dbconfig.GetDb()
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		if scope.PrimaryKey() == "" {
			return
		}
		// fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())

		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			// fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)

			var deleteValue string
			switch v := entry.key.(type) {
			case int:
				deleteValue = fmt.Sprint(v)

			case string:
				deleteValue = fmt.Sprintf("'%v'", v)

			default:
				logwrapper.Fatal("not implemented")
			}

			q := fmt.Sprintf(`DELETE FROM %v WHERE %v=%v`, entry.table, entry.keyname, deleteValue)
			db.Exec(q)
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}

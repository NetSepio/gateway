package models

import (
	"database/sql/driver"
)

type FlowIdType string

const (
	AUTH FlowIdType = "AUTH"
	ROLE FlowIdType = "ROLE"
)

func (fit *FlowIdType) Scan(value interface{}) error {
	*fit = FlowIdType([]byte(value.(string)))
	return nil
}

func (fit FlowIdType) Value() (driver.Value, error) {
	return string(fit), nil
}

type FlowId struct {
	FlowIdType    FlowIdType `sql:"flow_id_type"`
	WalletAddress string
	FlowId        string `gorm:"unique"`
}

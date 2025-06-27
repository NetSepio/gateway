package database

import (
	migrate "netsepio-gateway-v1.1/models/Migrate"
	"netsepio-gateway-v1.1/utils/load"
)

func Migrate() error {

	// db.Exec(`ALTER TABLE leader_boards DROP COLUMN IF EXISTS users;`)
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	for _, model := range []any{
		&migrate.User{},
		&migrate.Role{},
		&migrate.UserFeedback{},
		&migrate.FlowId{},
		&migrate.Report{},
		&migrate.ReportTag{},
		&migrate.ReportImage{},
		&migrate.ReportVote{},
		&migrate.Review{},
		&migrate.WaitList{},
		&migrate.Domain{},
		&migrate.DomainAdmin{},
		&migrate.DomainClaim{},
		&migrate.EmailAuth{},
		&migrate.SchemaMigration{},
		&migrate.SiteInsight{},
		&migrate.UserStripePi{},
		&migrate.Sotreus{},
		&migrate.Erebrus{},
		&migrate.Leaderboard{},
		&migrate.NftSubscription{},
		&migrate.DVPNNFTRecord{},
		&migrate.ScoreBoard{},
		&migrate.ActivityUnitXp{},
		&migrate.ReferralDiscount{},
		&migrate.ReferralAccount{},
		&migrate.ReferralSubscription{},
		&migrate.ReferralEarnings{},
		&migrate.Node{},
		&migrate.Organisation{},
		&migrate.OrganisationApp{},
		&migrate.Plan{},
		&migrate.SubscriptionPlan{},
		&migrate.SubscriptionRenewal{},
	} {
		if err := DB.AutoMigrate(model); err != nil {
			load.Logger.Sugar().Fatalf("failed to migrate %T: %v", model, err.Error())
			return err
		}
	}

	load.Logger.Info("Migrated all models")

	return nil
}

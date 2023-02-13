package orm

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  migrate,
}

var migrationUp bool
var migrationDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migrationUp, "up", "u", true, "running auto migration")
	MigrateCmd.Flags().BoolVarP(&migrationDown, "down", "d", false, "running auto reset migration")
}

func migrate(cmd *cobra.Command, args []string) error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	if migrationDown {
		return db.Migrator().DropTable(&model.User{}, &model.Vehicle{}, &model.Reservation{})
	}

	if migrationUp {
		return db.AutoMigrate(&model.User{}, &model.Vehicle{}, &model.Reservation{})
	}

	return nil

}

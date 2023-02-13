package command

import (
	"github.com/rfauzi44/vehicle-rental/database/orm"
	"github.com/rfauzi44/vehicle-rental/lib/server"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "vehicle rental backend",
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)

}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}

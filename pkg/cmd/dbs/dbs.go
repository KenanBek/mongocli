package dbs

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ListDatabaseNamesOp interface {
	ListDatabaseNames() ([]string, error)
}

// NewCmdDbs returns an instance of Dbs command
func NewCmdDbs(op ListDatabaseNamesOp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dbs",
		Short: "List database names",
		Long:  `long desc for dbs`,
		Run: func(cmd *cobra.Command, args []string) {
			RunDbs(op)
		},
	}

	return cmd
}

// RunDbs is exported.
func RunDbs(op ListDatabaseNamesOp) {
	names, err := op.ListDatabaseNames()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Dbs command error"))
		os.Exit(1)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

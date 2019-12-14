package colls

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCollectionNamesOp.
type ListCollectionNamesOp interface {
	ListCollectionNames(databaseName string) ([]string, error)
}

// CollsOptions is exported.
type CollsOptions struct {
	Database string
}

// NewCmdColls returns an instance of Colls command.
func NewCmdColls(op ListCollectionNamesOp) *cobra.Command {
	o := CollsOptions{}

	cmd := &cobra.Command{
		Use:   "colls",
		Short: "list collection names",
		Long:  `list collection names in the given database`,
		Run: func(cmd *cobra.Command, args []string) {
			o.Fill(cmd)
			o.Execute(op)
		},
	}

	return cmd
}

// Fill is exported.
func (o *CollsOptions) Fill(cmd *cobra.Command) {
	o.Database = viper.GetString("database")
}

// Execute is exported.
func (o *CollsOptions) Execute(op ListCollectionNamesOp) {
	names, err := op.ListCollectionNames(o.Database)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Dbs command error"))
		os.Exit(1)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

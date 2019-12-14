package count

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CountDocumentsOp is exported.
type CountDocumentsOp interface {
	CountDocuments(databaseName, collectionName string) (int64, error)
}

// CountOptions is exported.
type CountOptions struct {
	Database   string
	Collection string
}

// NewCmdCount returns an instance of Count command.
func NewCmdCount(op CountDocumentsOp) *cobra.Command {
	o := CountOptions{}

	cmd := &cobra.Command{
		Use:   "count",
		Short: "count documents",
		Long:  `count documents in the given collection`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a collection name argument")
			}

			o.Collection = args[0]
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			o.Fill(cmd)
			o.Execute(op)
		},
	}

	return cmd
}

// Fill is exported.
func (o *CountOptions) Fill(cmd *cobra.Command) {
	o.Database = viper.GetString("database")
}

// RunList is exported.
func (o *CountOptions) Execute(op CountDocumentsOp) {
	count, err := op.CountDocuments(o.Database, o.Collection)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error on count documents"))
	}
	fmt.Println(count)
}

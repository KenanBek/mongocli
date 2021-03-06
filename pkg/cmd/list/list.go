package list

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

// ListDocumentsOp is exported.
type ListDocumentsOp interface {
	ListDocuments(db, coll string) ([]bson.M, []error)
}

// ListOptions is exported.
type ListOptions struct {
	Database   string
	Collection string
}

// NewCmdList returns an instance of List command.
func NewCmdList(op ListDocumentsOp) *cobra.Command {
	o := ListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list documents",
		Long:  `list documents in the given collection`,
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
func (o *ListOptions) Fill(cmd *cobra.Command) {
	o.Database = viper.GetString("database")
}

// Execute is exported.
func (o *ListOptions) Execute(op ListDocumentsOp) {
	docs, errs := op.ListDocuments(o.Database, o.Collection)

	for _, err := range errs {
		fmt.Println(err)
	}

	for _, doc := range docs {
		fmt.Println(doc)
	}
}

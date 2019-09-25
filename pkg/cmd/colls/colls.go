/*
Copyright © 2019 KANAN RAHIMOV <mail@kenanbek.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

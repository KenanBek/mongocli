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

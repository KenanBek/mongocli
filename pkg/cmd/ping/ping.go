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
package ping

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// PingOp is exported.
type PingOp interface {
	Ping() error
}

// NewCmdPing returns an instance of Ping command
func NewCmdPing(op PingOp) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ping",
		Short: "List database names",
		Long:  `long desc for dbs`,
		Run: func(cmd *cobra.Command, args []string) {
			RunPing(op)
		},
	}

	return cmd
}

// RunPing is exported.
func RunPing(op PingOp) {
	err := op.Ping()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Database ping error"))
		os.Exit(1)
	}

	fmt.Println("Ping success!")
}

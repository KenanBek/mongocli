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

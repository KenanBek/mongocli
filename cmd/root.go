package cmd

import (
	"fmt"
	"os"

	"github.com/KenanBek/mongocli/pkg/mongo"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/KenanBek/mongocli/pkg/cmd/colls"
	"github.com/KenanBek/mongocli/pkg/cmd/count"
	"github.com/KenanBek/mongocli/pkg/cmd/dbs"
	"github.com/KenanBek/mongocli/pkg/cmd/list"
	"github.com/KenanBek/mongocli/pkg/cmd/ping"
)

var ro RootOptions

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongocli",
	Short: "MongoDB CLI",
	Long: `MongoCLI is a CLI application for MongoDB database that helps
to perform simple database operations just by using command line. It also
support file based configuration so you do not have to provide connection
details for each executed command.

Example:

  mongocli create col1 "{ title: 'title1', desc: 'desc1'}"

This command will use mongocli.yml file to connection to the database
and then create a document in collection col1.`,
}

// RootOptions is exported.
type RootOptions struct {
	config   string
	server   string
	port     int
	database string
}

// Execute is exported.
func Execute() {
	// Prepare configuration variables.
	initConfig()

	ro.config = viper.GetString("config")
	ro.server = viper.GetString("server")
	ro.port = viper.GetInt("port")

	mc := mongo.New(ro.server, ro.port)
	defer mc.Cancel()

	// Add sub-commands
	rootCmd.AddCommand(ping.NewCmdPing(mc))
	rootCmd.AddCommand(dbs.NewCmdDbs(mc))
	rootCmd.AddCommand(colls.NewCmdColls(mc))
	rootCmd.AddCommand(list.NewCmdList(mc))
	rootCmd.AddCommand(count.NewCmdCount(mc))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ro.config, "config", "", "config file (default is $HOME/mongocli.yml)")
	rootCmd.PersistentFlags().StringVarP(&ro.server, "server", "s", "localhost", "host name")
	rootCmd.PersistentFlags().IntVarP(&ro.port, "port", "p", 27017, "port number")
	rootCmd.PersistentFlags().StringVarP(&ro.database, "database", "d", "", "database name")

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")

	if ro.config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ro.config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "mongocli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("mongocli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}

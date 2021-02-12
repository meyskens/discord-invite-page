package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// to be overwritten in build
var revision = "dev"

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "discord-join-page",
		Short: "discord-join-page allows to dynamicly generate Discord join tokens",
		Long:  "discord-join-page allows to dynamicly generate Discord join tokens",
	}
)

func initConfig() {
	viper.AutomaticEnv()
}

func main() {
	flag.Parse()
	cobra.OnInitialize(initConfig)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	err := rootCmd.Execute()
	if err != nil {
		glog.Error(err)
	}
}

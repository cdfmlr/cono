package cmd

import (
	"conostudent/config"
	"conostudent/model"
	"conostudent/service"
	"conostudent/transport"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "conostudent",
	Short: "cono student server",
	Long:  `conostudent is a part of cono. It offers the student identity service for cono.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(cfgFile)
		model.Init()
		service.Init()
		transport.Init()
		transport.Serve(config.Serve.StudentRPCAddress)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	_ = rootCmd.MarkFlagRequired("config")
}

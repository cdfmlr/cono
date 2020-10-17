package cmd

import (
	"conocourse/config"
	"conocourse/model"
	"conocourse/service"
	"conocourse/transport"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "conocourse",
	Short: "cono course server",
	Long:  `conocourse is a part of cono. It offers the courses subscription service (including wechat serving, jw query & notice).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init
		config.Init(cfgFile)
		model.Init()
		transport.Init()
		service.Init()

		// Start services
		service.Run()
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

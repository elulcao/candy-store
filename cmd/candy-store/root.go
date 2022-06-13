/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	cs "github.com/elulcao/candy-store/pkg/candy-store"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "candy-store",
	Short: "Generate a JSON report with the users and snacks",
	Long:  "Generate a JSON report with the users and eaten snacks in sorted order",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := cs.ReadUsersFile(cmd.Flag("file").Value.String())

		if err != nil {
			os.Exit(1)
		}
		if _, err := cs.ReportUsers(users); err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.candy-store.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "", "", "CSV file with users, snacks and eaten snacks")
	rootCmd.MarkFlagRequired("file")
}

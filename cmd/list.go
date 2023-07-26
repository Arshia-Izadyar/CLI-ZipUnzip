/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigName("todos")
		v.AddConfigPath(".")
		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("Error reading todos.yml:", err)
			return
		}

		// Get the data from the configuration file
		todoList := v.AllSettings()
		td, err := json.MarshalIndent(todoList, "", "    ")
		if err != nil {
			fmt.Errorf("Error occured while reading todos", err)
			return
		}
		fmt.Println(string(td))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

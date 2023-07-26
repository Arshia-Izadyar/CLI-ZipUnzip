/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				fmt.Errorf("description adding failed!", err)
			}
			deadline, err := cmd.Flags().GetString("deadline")
			if err != nil {
				fmt.Errorf("Cant add deadline error happened", err)
			}
			todo := Todo{
				Name:        todoName,
				Description: description,
				Date:        deadline,
			}
			todos.TodoList = append(todos.TodoList, todo)
			viper.Set("todos", todos.TodoList)
			err = viper.WriteConfig()
			if err != nil {
				panic(err)
			}

		} else {
			fmt.Println("Add one todo")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	addCmd.Flags().String("description", "", "Todo description")
	addCmd.Flags().String("deadline", tomorrow.String(), "adds deadline")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"todoCli/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `Ad will create a new task into the task list.`,
	Run: addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {
	//fmt.Println("add called")

	//items := []todo.Item{}

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text:x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems("todos.json", items)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Println(items)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Allowed Priority:1,2,3")
}

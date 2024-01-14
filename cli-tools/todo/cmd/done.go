/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"todoCli/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Aliases: []string{"do"},
	Short: "Mark items as done",
	Long: `Mark items as done.`,
	Run: doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	//fmt.Println("done called")
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Errorf("%v", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked Done")

		todo.SaveItems(dataFile, items)
	} else {
		fmt.Println(i, "doesn't match any items")
	}

	
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

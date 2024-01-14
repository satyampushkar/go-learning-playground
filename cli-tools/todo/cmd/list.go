/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"strconv"
	"todoCli/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List all tasks.`,
	Run: listRun,
}

var (
	doneOpt bool
	allOpt bool
)	

func listRun(cmd *cobra.Command, args []string) {
	//fmt.Println("list called")
	items, err := todo.ReadItems("todos.json")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	//fmt.Println(items)
	pos := 1
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt{
			fmt.Fprintln(w, strconv.Itoa(pos) + ". "+ i.PrettyDone() +" (" + strconv.Itoa(i.Priority) + ")" + "\t", i.Text + "\t")
		}
		pos++
	}
	w.Flush()
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
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
}

package cmd

import (
	"fmt"
	"han109k/cliApp/todo"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long:  "Add will create a new todo item to the list",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	// x is index/key, y is value
	items, rErr := todo.ReadItems(dataFile)
	if rErr != nil {
		log.Printf("%v", rErr)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err := todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

var priority int

// special function, called after package variable declarations
// Each package may have multiple init(), order un-guaranteed
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority:1,2,3,4,5")
}

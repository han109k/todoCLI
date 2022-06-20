package cmd

import (
	"fmt"
	"han109k/cliApp/todo"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  "Listing of the todos",
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPrio(items))

	// tab writer
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		isDone := ""
		if i.PrettyDone() {
			isDone = "✔"
		}
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyP()+"\t"+i.Text+"\t"+isDone)
		}
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
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show 'Done' Todos")
}

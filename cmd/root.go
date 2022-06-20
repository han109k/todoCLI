package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cliApp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	/**
	* Errors are not exceptional, they are just values
	*	No exceptions in Go language
	* Errors should be handled when they occur
	 */
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(".cliApp")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cliApp")
	// export CLIAPP_DATAFILE=$HOME/Documents/todos.json
	// or put "datafile: $HOME/Documents/todos.json" in the config file -> .cliApp.yaml

	// if a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".todo.json", "data file to store todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cliApp.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

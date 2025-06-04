/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"snippet-manger/internal/snippet"
	"snippet-manger/internal/store"

	"github.com/spf13/cobra"
)

var (
	addName        string
	addContent     string
	addDescription string
	addTags        []string
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
		currentSnippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("error: unable to load snippet" + err.Error())
			os.Exit(1)
		}
		snippet.SetNextId(currentSnippets)

		newSnippet := snippet.New(addName, addContent, addTags, addDescription)
		store.SaveSnippet(append(currentSnippets, newSnippet))

		fmt.Println("Created a New Snippet")
		fmt.Println(newSnippet.String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Name of snippet (required")
	addCmd.Flags().StringVarP(&addContent, "content", "c", "", "Content of snippet (required)")
	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "Description of the snippet")
	addCmd.Flags().StringSliceVarP(&addTags, "tag", "t", []string{}, "comma-separated tags for the snippet (can be repeated)")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("content")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"snippet-manger/internal/snippet"
	"snippet-manger/internal/store"

	"github.com/spf13/cobra"
)

var deleteId string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentSnippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("error: " + err.Error())
		}
		newSnippets, err := snippet.Delete(deleteId, currentSnippets)
		if err != nil {
			fmt.Println("error: " + err.Error())
		}
		store.SaveSnippet(newSnippets)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	deleteCmd.Flags().StringVarP(&deleteId, "id", "i", "", "Id of snippet (required")
	deleteCmd.MarkFlagRequired("id")
}

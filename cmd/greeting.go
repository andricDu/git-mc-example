/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetingCmd represents the greeting command
var greetingCmd = &cobra.Command{
	Use:   "greeting",
	Short: "TODO: Add a short description here",               // TODO
	Long:  `TODO: Add a longer description here or remove it`, // TODO
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("greeting called")
	},
}

func init() {
	rootCmd.AddCommand(greetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

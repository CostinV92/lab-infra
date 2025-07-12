/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
	Long:  "start a service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		fmt.Println(Service)
	},
}

func init() {

}

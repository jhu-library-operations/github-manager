package cmd

import (
	"github.com/spf13/cobra"
	//"github.com/shurcooL/githubv4"
	"fmt"
)

var orgCmd = &cobra.Command {
	Use: "org [action]",
	Short: "Organization level functions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Incorrect number of arguments, accepts 1")
		}

		var action string
		if len(args) == 1 {
			action = args[0]
		}

		switch action {
		case "get":
			fmt.Println("get an organization")
		case "new":
			fmt.Println("create an organization")
		case "delete":
			fmt.Println("destroy an organization")
		default:
			fmt.Println("you're confused")
		}
	},
}

func init() {
	rootCmd.AddCommand(orgCmd)

	// orgCmd.Flags().StringVarP(&OrgName, "name", "n", nil, "The name of the organization")
}

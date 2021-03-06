package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish API definitions from a Git repo to a gateway or dashboard",
	Long:  `Publish API definitions from a Git repo to a gateway or dashboard, this
	will not update existing APIs, and if it detects a collision, will stop.`,
	Run: func(cmd *cobra.Command, args []string) {
		gwString, _ := cmd.Flags().GetString("gateway")
		dbString, _ := cmd.Flags().GetString("dashboard")

		if gwString == "" && dbString == "" {
			fmt.Println("Publish requires either gateway or dashboard target to be set")
			return
		}

		if gwString != "" && dbString != "" {
			fmt.Println("Publish requires either gateway or dashboard target to be set, not both")
			return
		}

		err := processPublish(cmd, args)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.
	publishCmd.Flags().StringP("gateway", "g", "", "Fully qualified gateway target URL")
	publishCmd.Flags().StringP("dashboard", "d", "", "Fully qualified dashboard target URL")
	publishCmd.Flags().StringP("key", "k", "", "Key file location for auth (optional)")
	publishCmd.Flags().StringP("branch", "b", "refs/heads/master", "Branch to use (defaults to refs/heads/master)")
	publishCmd.Flags().StringP("secret", "s", "", "Your API secret")
	publishCmd.Flags().Bool("test", false, "Use test publisher, output results to stdio")
}

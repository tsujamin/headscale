package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(aclCmd)
	aclCmd.AddCommand(getAclPolicyCmd)
	aclCmd.AddCommand(setAclPolicyCmd)

	// TODO arguments
}

var aclCmd = &cobra.Command{
	Use:     "aclpolicy",
	Short:   "ACL policy",
	Aliases: []string{"acl"},
}

var getAclPolicyCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the ACL policy", // TODO
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unimplemented") // TODO
	},
}

var setAclPolicyCmd = &cobra.Command{
	Use:     "set",
	Short:   "Sets the ACL policy", // TODO
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unimplemented") // TODO
	},
}

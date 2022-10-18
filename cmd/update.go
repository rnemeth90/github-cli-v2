/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// const url = "https://"

		// flags := cmd.Flags()

		// repoData := Repository{
		// 	Name:                      mustString(flags, "name"),
		// 	Owner:                     mustString(flags, "owner"),
		// 	Description:               mustString(flags, "description"),
		// 	Private:                   mustBool(flags, "private"),
		// 	Homepage:                  mustString(flags, "homepage"),
		// 	HasIssues:                 mustBool(flags, "hasissues"),
		// 	HasWiki:                   mustBool(flags, "haswiki"),
		// 	IsTemplate:                mustBool(flags, "istemplate"),
		// 	DefaultBranch:             mustBool(flags, "defaultbranch"),
		// 	AllowSquashMerge:          mustBool(flags, "allowsquashmerge"),
		// 	AllowMergeMerge:           mustBool(flags, "allowmergemerge"),
		// 	AllowRebaseMerge:          mustBool(flags, "allowrebasemerge"),
		// 	AllowAutoMerge:            mustBool(flags, "allowautomerge"),
		// 	DeleteBranchOnMerge:       mustBool(flags, "deletebranchonmerge"),
		// 	AllowUpdateBranch:         mustBool(flags, "allowupdatebranch"),
		// 	UseSquashPRTitleAsDefault: mustBool(flags, "usesquashprtitleasdefault"),
		// 	SquashMergeCommitTitle:    mustString(flags, "squashmergecommittitle"),
		// 	SquashMergeCommitMessage:  mustString(flags, "squashmergecommitmessage"),
		// 	MergeCommitTitle:          mustString(flags, "mergecommittitle"),
		// 	MergeCommitMessage:        mustString(flags, "mergecommitmessage"),
		// 	Archived:                  mustBool(flags, "archived"),
		// 	AllowForking:              mustBool(flags, "allowforking"),
		// 	WebCommitSignoffRequired:  mustBool(flags, "webcommitsignoffrequired"),
		// }

		// jsonData, err := json.Marshal(repoData)
		// if err != nil {
		// 	log.Fatal(err)
		// }

	},
}

func init() {
	updateCmd.Flags().SortFlags = true
	repoCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/
package repo

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rnemeth90/github-cli-v2/cmd/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		owner := utils.ParseString(flags, "owner")
		repo := utils.ParseString(flags, "name")
		token := utils.ParseString(flags, "token")

		url := "https://api.github.com/repos/" + owner + "/" + repo

		client := &http.Client{}
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/vnd.github+json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			log.Printf("Unable to delete repo %s\n", utils.ParseString(flags, "name"))
			log.Printf("Response code is %d\n", resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(utils.JsonPrettyPrint(string(body)))
			log.Fatal(err)
		}

		log.Printf("%s deleted successfully", utils.ParseString(flags, "name"))

	},
}

func init() {
	deleteCmd.Flags().SortFlags = true
	repoCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("owner", "o", "", "The owner")
	deleteCmd.MarkFlagRequired("owner")
	deleteCmd.Flags().StringP("token", "t", "", "The token")
	deleteCmd.MarkFlagRequired("token")
	deleteCmd.Flags().StringP("name", "n", "", "The name")
	deleteCmd.MarkFlagRequired("name")
}

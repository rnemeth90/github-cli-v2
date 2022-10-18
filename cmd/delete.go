/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		owner := mustString(flags, "owner")
		repo := mustString(flags, "name")
		token := mustString(flags, "token")

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
			log.Printf("Unable to delete repo %s\n", mustString(flags, "name"))
			log.Printf("Response code is %d\n", resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(jsonPrettyPrint(string(body)))
			log.Fatal(err)
		}

		log.Printf("%s deleted successfully", mustString(flags, "name"))

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

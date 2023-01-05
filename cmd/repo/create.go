/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/
package repo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rnemeth90/github-cli-v2/cmd/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		const url = "https://api.github.com/user/repos"

		flags := cmd.Flags()

		log.Printf("Creating repo %s...\n", utils.ParseString(flags, "name"))

		repoData := Repository{
			Name:        utils.ParseString(flags, "name"),
			Description: utils.ParseString(flags, "description"),
			Private:     utils.ParseBool(flags, "private"),
		}

		jsonData, err := json.Marshal(repoData)
		if err != nil {
			log.Fatal(err)
		}

		t := utils.ParseString(flags, "token")

		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+t)
		req.Header.Set("Accept", "application/vnd.github+json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Printf("Unable to create repo %s\n", utils.ParseString(flags, "name"))
			log.Printf("Response code is %d\n", resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(utils.JsonPrettyPrint(string(body)))
			log.Fatal(err)
		}

		log.Printf("%s created successfully", utils.ParseString(flags, "name"))
	},
}

func init() {
	createCmd.Flags().SortFlags = true
	repoCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the repo")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringP("description", "d", "", "Description of the repo")
	createCmd.Flags().StringP("token", "t", "", "A token")
	createCmd.MarkFlagRequired("token")
	createCmd.Flags().BoolP("private", "p", false, "Create a private repo")
}

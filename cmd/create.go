/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

		log.Printf("Creating repo %s...\n", mustString(flags, "name"))

		repoData := Repository{
			Name:        mustString(flags, "name"),
			Description: mustString(flags, "description"),
			Private:     mustBool(flags, "private"),
		}

		jsonData, err := json.Marshal(repoData)
		if err != nil {
			log.Fatal(err)
		}

		t := mustString(flags, "token")

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
			log.Printf("Unable to create repo %s\n", mustString(flags, "name"))
			log.Printf("Response code is %d\n", resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(jsonPrettyPrint(string(body)))
			log.Fatal(err)
		}

		log.Printf("%s created successfully", mustString(flags, "name"))
	},
}

func init() {
	createCmd.Flags().SortFlags = true
	repoCmd.AddCommand(createCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("name", "n", "", "Name of the repo")
	createCmd.Flags().StringP("description", "d", "", "Description of the repo")
	createCmd.Flags().StringP("token", "t", "", "A token")
	createCmd.Flags().BoolP("private", "p", false, "Create a private repo")
}

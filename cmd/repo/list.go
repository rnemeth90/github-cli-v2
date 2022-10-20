/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/

package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repos for an authenticated user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		const url = "https://api.github.com/user/repos"

		flags := cmd.Flags()
		t := mustString(flags, "token")

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
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

		body, _ := ioutil.ReadAll(resp.Body)

		var jsonData []Repository
		err = json.Unmarshal([]byte(body), &jsonData) // here!
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("Unable to create repo %s\n", mustString(flags, "name"))
			log.Printf("Response code is %d\n", resp.StatusCode)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(jsonPrettyPrint(string(body)))
			log.Fatal(err)
		}
		fmt.Println(structPrettyPrintToJSON(jsonData))
	},
}

func init() {
	listCmd.Flags().SortFlags = true
	repoCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("token", "t", "", "Authentication token")
	listCmd.MarkFlagRequired("token")
}

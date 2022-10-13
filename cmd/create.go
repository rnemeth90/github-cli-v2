/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Repository defines a github repo
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func mustString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}
	return v
}

func mustBool(fs *pflag.FlagSet, name string) bool {
	v, err := fs.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// structPrettyPrint to print struct in a readable way
func structPrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func init() {
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

/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/

package cmd

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/pflag"
)

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

// structPrettyPrintToJson to print struct in a readable way
func structPrettyPrintToJSON(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

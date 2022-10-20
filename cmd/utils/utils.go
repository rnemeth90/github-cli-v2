/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/

package utils

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/pflag"
)

func ParseString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}
	return v
}

func ParseBool(fs *pflag.FlagSet, name string) bool {
	v, err := fs.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

// JsonPrettyPrint will return JSON parsed in a readable format
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// StructPrettyPrintToJSON will return a parsed struct in a readable format
func StructPrettyPrintToJSON(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

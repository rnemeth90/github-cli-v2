/*
Copyright © 2022 Ryan Nemeth ryannemeth(at)live(dot)com
*/
package main

import (
	"github.com/rnemeth90/github-cli-v2/cmd"
	_ "github.com/rnemeth90/github-cli-v2/cmd/repo"
)

func main() {
	cmd.Execute()
}

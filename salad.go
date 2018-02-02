// +build go1.6

// Copyright 2018 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/urfave/cli"
	"github.com/vsoch/salad/cmd"
	"github.com/vsoch/salad/config"
	"os"
)

const Version = "0.0.1"
const Name = "Salad Fork"
const Github = "https://www.github.com/vsoch/salad"

func init() {
	config.Version = Version
	config.Name = Name
	config.Github = Github
}

func main() {
	app := cli.NewApp()
	app.Name = "Salad Fork"
	app.Usage = "Stick this in your salad, with caution, grasshopper!"
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.Fork,
		cmd.Spoon,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}

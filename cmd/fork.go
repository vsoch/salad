// Copyright 2018 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"time"
)

// -----------------------------------------------------------------------------
//
// Command Entrypoint
//
// -----------------------------------------------------------------------------

var Fork = cli.Command{
	Name:        "fork",
	Usage:       "Fork",
	Description: `Create a fork`,
	Action:      printFork,
	Flags: []cli.Flag{
		stringFlag("message, m", "", "Give wisdom for your fork."),
	},

	//TODO: add subcommands for color, etc.
	// https://github.com/gogits/gogs/blob/master/cmd/hook.go#L40
}

// -----------------------------------------------------------------------------
//
// Selection Functions
// Choose a pun randomly from the list
//
// -----------------------------------------------------------------------------

func selectRandom(choices ...string) string {

	// Seed random number generator, won't be random without it
	rand.Seed(time.Now().Unix())

	// Make random selection
	return choices[rand.Intn(len(choices))]
}

func selectPun() string {

	// Create a "slice" (array) of puns
	puns := []string{" You're done!",
		" Cut it out.",
		" Fork off, already.",
		" You want a peas of me?",
		" Have a knife day? Fork you!",
		" My life purpose: I cut butter.",
		" Forkin' repos, that's what I do.",
		" What the F*rk are you asking me for?!",
		" You think you have problems? I'm a fork.",
		" Take the one less traveled, they said...",
		" In Go an array is a slice. Utensil discrimination!",
		" I'm not much of a traveler, but people take me anyway.",
		" I can't help with yo' Momma, I'm not that kind of fork."}

	// Randomly select one
	return selectRandom(puns...)
}

func selectColor() string {

	colors := []string{"",
		"\033[95m", // purple
		"\033[93m", // yellow
		"\033[91m", // red
		"\033[31m", // darkred
		"\033[36m"} // cyan

	return selectRandom(colors...)
}

func selectFork() string {

	forks := make([]string, 5)

	// source
	// ascii.co.uk/art/fork

	forks[0] = `

                   ________  .====
                  [________>< :===
                             '====`

	forks[1] = `

          _________________  .========
         [_________________>< :======
                             '========`

	forks[2] = `

                            _
                           / )
                     |||| / /
                     ||||/ /
                     \__(_/
                      ||//
                      ||/
                      ||
                     (||
                      ""`

	forks[3] = `

                       /\
                      //\\
                     //  \\
                 ^   \\  //   ^
                / \   )  (   / \ 
                ) (   )  (   ) (
                \  \_/ /\ \_/  /
                 \__  _)(_  __/
                    \ \  / /
                     ) \/ (
                     | /\ |
                     | )( |
                     | )( |
                     | \/ |
                     )____(
                    /      \
                    \______/`

	forks[4] = `

                                   ⎯⎯∈

               `

	// Randomly select one
	return selectRandom(forks...)
}

// -----------------------------------------------------------------------------
//
// Main Function to Print Fork
// If the user provides message, print instead of pun
//
// -----------------------------------------------------------------------------

func printFork(ctx *cli.Context) {

	pun := selectPun()
	if ctx.IsSet("message") { // 1. Select a random pun
		pun = ctx.String("message")
	}

	fork := selectFork() // 2. Select a colored fork
	color := selectColor()

	fmt.Println()
	fmt.Println(pun, color, fork, "\033[0m") // off sequence to end color
	fmt.Println()

	os.Exit(0)

}

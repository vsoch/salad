// Copyright 2018 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
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
		stringFlag("color, c", "", "color your spoon"),
	},
}

// -----------------------------------------------------------------------------
//
// Selection Functions
// Choose a pun randomly from the list
//
// -----------------------------------------------------------------------------


func selectForkPun() string {

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

        // 1. Select a random pun

	pun := selectForkPun()
	if ctx.IsSet("message") {
		pun = ctx.String("message")
	}

        
        // 2. Select a fork

	fork := selectFork()

        // 3. Select a color

        color := ctx.String("color")
        color = selectColor(color)

	fmt.Println()
	fmt.Println(pun, color, fork, "\033[0m") // off sequence to end color
	fmt.Println()

	os.Exit(0)

}

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

var Spoon = cli.Command{
	Name:        "spoon",
	Usage:       "Spoon",
	Description: `Create a spoon`,
	Action:      printSpoon,
	Flags: []cli.Flag{
		stringFlag("message, m", "", "Give wisdom for your spoon."),
		stringFlag("color, c", "", "color your spoon"),
	},
}

// -----------------------------------------------------------------------------
//
// Selection Functions
// Choose a pun randomly from the list
//
// -----------------------------------------------------------------------------

func selectSpoonPun() string {

	// Create a "slice" (array) of puns
	puns := []string{" See you spoon!",
		" They say I'm rather spoontaneous!",
		" Did you hear about... oh, too spoon?",
		" They call me the cereal kiler.",
		" Wanna spoon?",
		" My favorite actress? Reese Witherspoon!",
		" *singing* The dark side of... the spooooon!",
		" I'm a spoon. How would I know?",
		" If you want to be sharp, I'm the wrong utensil.",
		" I don't have a point. Go talk to Fork."}

	// Randomly select one
	return selectRandom(puns...)
}

func selectSpoon() string {

	spoons := make([]string, 5)

	// source
	// spoon.ascii.uk/

	spoons[0] = `

           ________   .==.
          [________>c((_  )
                      '=='

         `

	spoons[1] = `


        __________        .-"""-.
       /          ''''---' .'    \
       \__________....---. '.    /
                          '-...-'
        `

	spoons[2] = `

                              _
                             / \
          _..--"""""--.._    \_/
         /,_..-------.._,\    |
        |  ''''-----''''  |   |
         \               /   / \
          '.           .'    | |
            '--.....--'      \_/
        `

	spoons[3] = `

          ___           .-""-.
         /   '''---...-'.'  '\\
         \___...---"""-._-.__//
                         '---'

        `

	spoons[4] = `


           $$$$$$$$$
          $$$$$$$$$$$
         $$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
         $$$$$$$$$$$$$$
          $$$$$$$$$$$$
           $$$$$$$$$$
            $$$$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
            $$$$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
              $$$$

        `

	// Randomly select one
	return selectRandom(spoons...)
}

// -----------------------------------------------------------------------------
//
// Main Function to Print Spoon
// If the user provides message, print instead of pun
//
// -----------------------------------------------------------------------------

func printSpoon(ctx *cli.Context) {

	// 1. Select a random pun

	pun := selectSpoonPun()
	if ctx.IsSet("message") {
		pun = ctx.String("message")
	}

	spoon := selectSpoon()

	// 2. Select a colored spoon

	color := ctx.String("color")
	color = selectColor(color)

	fmt.Println()
	fmt.Println(pun, color, spoon, "\033[0m") // off sequence to end color
	fmt.Println()

	os.Exit(0)

}

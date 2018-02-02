// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/urfave/cli"
	"math/rand"
	"time"
)



func selectRandom(choices ...string) string {

	// Seed random number generator, won't be random without it
	rand.Seed(time.Now().Unix())

	// Make random selection
	return choices[rand.Intn(len(choices))]
}


func selectRandomMap(choices map[string]string) string {

        // If given a map, convert to array of values first
        values := make([]string, len(choices))
        i := 0
        for k := range choices {
            values[i] = choices[k]
            i++
        } 

        return selectRandom(values...)
}



func selectColor(color string) string {

	colors := map[string]string{
		"purple":  "\033[95m",
		"yellow":  "\033[93m",
		"red":     "\033[91m",
		"darkred": "\033[31m",
		"cyan":    "\033[36m"}

	if color != "" {
		if val, ok := colors[color]; ok {
			return colors[color]
		} else {
                        _ = val // declared and not used error
                }
	}

	return selectRandomMap(colors)
}


func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

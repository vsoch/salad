package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)



// 
// Selection Functions
// Choose a pun randomly from the list
//

func select_random(choices ...string) string {

    // Seed random number generator, won't be random without it
    rand.Seed(time.Now().Unix())

    // Make random selection
    return choices[rand.Intn(len(choices))]
}

func select_pun() string {

    // Create a "slice" (array) of puns
    puns := []string{"  You're done!",
                     "  Cut it out.",
                     "  Fork off, already.",
                     "  You want a peas of me?",
                     "  Have a knife day? Fork you!",
                     "  My life purpose: I cut butter.",
                     "  Forkin' repos, that's what I do.",
                     "  What the F*rk are you asking me for?!",
                     "  You think you have problems? I'm a fork.",
                     "  Take the one less traveled, they said...",
                     "  In Go an array is a slice. Utensil discrimination!",
                     "  I'm not much of a traveler, but people take me anyway.",
                     "  I can't help with yo' Momma, I'm not that kind of fork."}

   // Randomly select one 
   return select_random(puns...)
}


func select_fork() string {

    
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
    return select_random(forks...)
}


// 
// Print Functions
//



func main() {

    pun := select_pun()              // 1. Select a random pun
    fork := select_fork()            // 2. Select a colored fork

    fmt.Println()
    fmt.Println(pun, fork)
    fmt.Println()

    os.Exit(0)

}

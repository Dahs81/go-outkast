package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-outkast"
	app.Usage = "test outkast"
	app.Action = func(c *cli.Context) {
		// Grab all the arguments
		ca := c.Args()

		// Get the total number of arguments
		l := len(c.Args())

		// Select a random item from args
		removedItem := random(0, l)

		// Remove random item from ca
		ca = append(ca[:removedItem], ca[removedItem+1:]...)

		for i := range ca {
			print(c.Args()[i], " ")
		}
	}

	app.Run(os.Args)
}

// Random number generator function 2
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-outkast"
	app.Usage = "A fun commandline utility that allows you to randomly remove item(s) from a list of args."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rand, r",
			Value: "1",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) {
		// Grab all the arguments
		ca := c.Args()

		// Variables to use later
		var (
			rc  int
			err error
		)

		if c.String("rand") != "" {
			r := c.String("rand")
			rc, err = strconv.Atoi(r)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}

		// Get the total number of arguments
		l := len(c.Args())

		if rc > 1 && rc < len(c.Args()) {
			for i := 0; i < rc; i++ {
				ca, l = Remove(ca, l)
			}
		} else {
			ca, l = Remove(ca, l)
		}

		for i := range ca {
			print(c.Args()[i], " ")
		}
	}

	app.Run(os.Args)
}

// Remove -
func Remove(ca []string, l int) ([]string, int) {
	// Select a random item from args
	removedItem := random(0, l)

	// Remove random item from ca
	ca = append(ca[:removedItem], ca[removedItem+1:]...)
	// Reset the length of the slice in order to return it
	l = len(ca)
	return ca, l
}

// Random number generator function 2
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

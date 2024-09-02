package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "Health Checker",
		Usage: "Small tool to check if a domain + port is up or down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "The domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "The port of the domain to check\nDefault is 80",
				Required: false,
			},
			&cli.IntFlag{
				Name: "timeout",
				Aliases: []string{"t"},
				Usage: "The number of seconds that should be used before the app times out\nDefault is 5",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}

			timeout := c.Int("timeout")
			if c.Int("timeout") == 0 {
				timeout = 5
			}

			status := Check(c.String("domain"), port, timeout)

			fmt.Println(status)

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
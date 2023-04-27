package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// return cli app ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "cli"
	app.Usage = "Find IP and DNS information about a domain"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find IP information",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: FindIps,
		},
	}

	return app
}

func FindIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

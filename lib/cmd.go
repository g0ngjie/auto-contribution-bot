package lib

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	flags []cli.Flag
	host  string
	port  int
)

func init() {
	flags = []cli.Flag{
		cli.StringFlag{
			Name:        "t, host, ip-address",
			Value:       "127.0.0.1",
			Usage:       "Server host",
			Destination: &host,
		},
		cli.IntFlag{
			Name:        "p, port",
			Value:       8000,
			Usage:       "Server port",
			Destination: &port,
		},
	}

	app := cli.NewApp()
	app.Name = "贡献度小机器人"
	app.Usage = ""
	app.HideVersion = true
	app.Version = "1.2"
	app.Flags = flags
	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	if c.Int("port") < 1024 {
		cli.ShowAppHelp(c)
		return cli.NewExitError("Ports below 1024 is not available", 2)
	}

	return nil
}

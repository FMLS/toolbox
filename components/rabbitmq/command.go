package rabbitmq

import "github.com/urfave/cli/v2"

func GetCommand() *cli.Command {
	return &cli.Command{
		Name: "rmq",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"i"},
				Value:   "127.0.0.1",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 5672,
			},
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:    "vhost",
				Aliases: []string{"v"},
			},
			&cli.StringFlag{
				Name:    "exchange",
				Aliases: []string{"e"},
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:    "publish",
				Aliases: []string{"p"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "key",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "data",
						Aliases:  []string{"d"},
						Required: true,
					},
				},
				Action: Publish,
			},
			{
				Name:    "consume",
				Aliases: []string{"c"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "queue",
						Aliases:  []string{"q"},
						Required: true,
					},
				},
				Action: Consume,
			},
		},
	}
}

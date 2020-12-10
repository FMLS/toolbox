package tat

import "github.com/urfave/cli/v2"

const (
	FlagApiHost     = "ApiHost"
	FlagForwardHost = "ForwardHost"
	FlagApiPort     = "ApiPort"
	FlagWsPort      = "WsPort"

	FlagAppID  = "AppID"
	FlagUin    = "Uin"
	FlagRegion = "Region"
	FlagInsID  = "InsID"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name: "tat",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  FlagApiHost,
				Value: "127.0.0.1",
			},
			&cli.StringFlag{
				Name:  FlagForwardHost,
				Value: "127.0.0.1",
			},
			&cli.IntFlag{
				Name:  FlagApiPort,
				Value: 8520,
			},
			&cli.IntFlag{
				Name:  FlagWsPort,
				Value: 8900,
			},
			&cli.IntFlag{
				Name:  FlagAppID,
				Value: 1251783334,
			},
			&cli.StringFlag{
				Name:  FlagUin,
				Value: "3205597606",
			},
			&cli.StringFlag{
				Name:     FlagRegion,
				Required: true,
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:    "TestConn",
				Aliases: []string{"tc"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     FlagInsID,
						Required: true,
					},
				},
				Action: TestConn,
			},
		},
	}
}

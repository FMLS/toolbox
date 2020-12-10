package main

import (
	"log"
	"os"

	"github.com/FMLS/toolbox/components/rabbitmq"
	"github.com/FMLS/toolbox/components/tat"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			rabbitmq.GetCommand(),
			tat.GetCommand(),
		},
	}

	log.Println(app.Run(os.Args))
}

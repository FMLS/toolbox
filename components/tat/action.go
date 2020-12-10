package tat

import (
	"fmt"

	"github.com/FMLS/toolbox/components/tat/sessionmgr"
	"github.com/urfave/cli/v2"
)

func TestConn(ctx *cli.Context) error {
	sm := sessionmgr.NewSessionManager(ctx.String(FlagRegion), ctx.String(FlagApiHost),
		ctx.String(FlagForwardHost), ctx.Int(FlagApiPort), ctx.Int(FlagWsPort), ctx.Int(FlagAppID),
		ctx.String(FlagUin), ctx.String(FlagInsID),
	)
	if err := sm.TestConn(); err != nil {
		fmt.Println(err)
	}
	return nil
}

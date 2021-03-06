package cmd

import (
	"github.com/rancher/kontainer-engine/plugin"
	"github.com/rancher/kontainer-engine/types"
	"github.com/urfave/cli"
)

// runRPCDriver runs the rpc server and returns
func runRPCDriver(driverName string) (types.Driver, string, error) {
	// addrChan is the channel to receive the server listen address
	addrChan := make(chan string)
	plugin.Run(driverName, addrChan)

	addr := <-addrChan
	rpcClient, err := types.NewClient(driverName, addr)
	if err != nil {
		return nil, "", err
	}
	return rpcClient, addr, nil
}

// getDriverOpts get the flags and value and generate DriverOptions
func getDriverOpts(ctx *cli.Context) types.DriverOptions {
	driverOptions := types.DriverOptions{
		BoolOptions:        make(map[string]bool),
		StringOptions:      make(map[string]string),
		IntOptions:         make(map[string]int64),
		StringSliceOptions: make(map[string]*types.StringSlice),
	}
	for _, flag := range ctx.Command.Flags {
		switch flag.(type) {
		case cli.StringFlag:
			driverOptions.StringOptions[flag.GetName()] = ctx.String(flag.GetName())
		case cli.BoolFlag:
			driverOptions.BoolOptions[flag.GetName()] = ctx.Bool(flag.GetName())
		case cli.Int64Flag:
			driverOptions.IntOptions[flag.GetName()] = ctx.Int64(flag.GetName())
		case cli.StringSliceFlag:
			driverOptions.StringSliceOptions[flag.GetName()] = &types.StringSlice{
				Value: ctx.StringSlice(flag.GetName()),
			}
		}
	}
	return driverOptions
}

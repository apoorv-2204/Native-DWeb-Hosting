package cli

import (
	"strconv"

	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdHostManifest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host-manifest [version] [hash-function]",
		Short: "Broadcast message HostManifest",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVersion := args[0]
			argHashFunction := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgHostManifest(
				clientCtx.GetFromAddress().String(),
				argVersion,
				argHashFunction,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

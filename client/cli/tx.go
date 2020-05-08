package cli

import (
	"bufio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"

	"github.com/irismod/record/types"
)

// GetTxCmd returns the transaction commands for the record module.
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Record transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(flags.PostCommands(
		GetCmdCreateRecord(cdc),
	)...)
	return txCmd
}

// GetCmdCreateRecord implements the create record command.
func GetCmdCreateRecord(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [digest] [digest-algo]",
		Short: "Create a new record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			fromAddr := cliCtx.GetFromAddress()

			content := types.Content{
				Digest:     args[0],
				DigestAlgo: args[1],
				URI:        viper.GetString(FlagURI),
				Meta:       viper.GetString(FlagMeta),
			}

			msg := types.NewMsgCreateRecord([]types.Content{content}, fromAddr)
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsCreateRecord)
	return cmd
}

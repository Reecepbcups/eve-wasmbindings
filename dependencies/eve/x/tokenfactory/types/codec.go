package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "tokenfactory/CreateDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "tokenfactory/UpdateDenom", nil)
	cdc.RegisterConcrete(&MsgMintAndSendTokens{}, "tokenfactory/MintAndSendTokens", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "tokenfactory/UpdateOwner", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgUpdateDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintAndSendTokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateOwner{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

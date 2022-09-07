// manually done bc I cant get proto files to work...
// TODO:

use osmosis_std_derive::CosmwasmExt;
/// MsgCreateDenom is the sdk.Msg type for allowing an account to create
/// a new denom. It requires a sender address and a subdenomination.
/// The (sender_address, sub_denomination) pair must be unique and cannot be
/// re-used. The resulting denom created is `factory/{creator
/// address}/{subdenom}`. The resultant denom's admin is originally set to be the
/// creator, but this can be changed later. The token denom does not indicate the
/// current admin.
#[derive(
    Clone,
    PartialEq,
    ::prost::Message,
    serde::Serialize,
    serde::Deserialize,
    schemars::JsonSchema,
    CosmwasmExt,
)]
#[proto_message(type_url = "/eve.tokenfactory.v1beta1.MsgCreateDenom")]
pub struct MsgCreateDenom {
    #[prost(string, tag = "1")]
    pub owner: ::prost::alloc::string::String,

    /// TOKEN NAME
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    
    #[prost(string, tag = "3")]
    pub denom: ::prost::alloc::string::String,

    #[prost(string, tag = "4")]
    pub description: ::prost::alloc::string::String,

    #[prost(int32, tag = "5")]
    pub precision: i32,

    #[prost(int32, tag = "6")]
    pub max_supply: i32,

    #[prost(int32, tag = "7")]
    pub can_change_max_supply: i32,
}
/// MsgCreateDenomResponse is the return value of MsgCreateDenom
/// It returns the full string of the newly created denom
// #[derive( Clone, PartialEq, ::prost::Message, serde::Serialize, serde::Deserialize, schemars::JsonSchema, CosmwasmExt)]
// #[proto_message(type_url = "/eve.tokenfactory.v1beta1.QueryGetDenomResponse")]
// pub struct QueryGetDenomResponse {
//     #[prost(denom, tag = "1")]
//     pub new_token_denom: Denom,
// }

// /// MsgMint is the sdk.Msg type for allowing an admin account to mint
// /// more of a token.  For now, we only support minting to the sender account
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.MsgMint")]
// pub struct MsgMint {
//     #[prost(string, tag = "1")]
//     pub sender: ::prost::alloc::string::String,
//     #[prost(message, optional, tag = "2")]
//     pub amount: ::core::option::Option<super::super::super::cosmos::base::v1beta1::Coin>,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.MsgMintResponse")]
// pub struct MsgMintResponse {}
// /// MsgBurn is the sdk.Msg type for allowing an admin account to burn
// /// a token.  For now, we only support burning from the sender account.
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]

// /// MsgChangeAdmin is the sdk.Msg type for allowing an admin account to reassign
// /// adminship of a denom to a new account
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.MsgChangeAdmin")]
// pub struct MsgChangeAdmin {
//     #[prost(string, tag = "1")]
//     pub sender: ::prost::alloc::string::String,
//     #[prost(string, tag = "2")]
//     pub denom: ::prost::alloc::string::String,
//     #[prost(string, tag = "3")]
//     pub new_admin: ::prost::alloc::string::String,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.MsgChangeAdminResponse")]
// pub struct MsgChangeAdminResponse {}
// /// DenomAuthorityMetadata specifies metadata for addresses that have specific
// /// capabilities over a token factory denom. Right now there is only one Admin
// /// permission, but is planned to be extended to the future.
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata")]
// pub struct DenomAuthorityMetadata {
//     /// Can be empty for no admin, or a valid osmosis address
//     #[prost(string, tag = "1")]
//     pub admin: ::prost::alloc::string::String,
// }
// /// Params holds parameters for the tokenfactory module
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.Params")]
// pub struct Params {
//     #[prost(message, repeated, tag = "1")]
//     pub denom_creation_fee:
//         ::prost::alloc::vec::Vec<super::super::super::cosmos::base::v1beta1::Coin>,
// }
// /// QueryParamsRequest is the request type for the Query/Params RPC method.
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.QueryParamsRequest")]
// #[proto_query(
//     path = "/osmosis.tokenfactory.v1beta1.Query/Params",
//     response_type = QueryParamsResponse
// )]
// pub struct QueryParamsRequest {}
// /// QueryParamsResponse is the response type for the Query/Params RPC method.
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.QueryParamsResponse")]
// pub struct QueryParamsResponse {
//     /// params defines the parameters of the module.
//     #[prost(message, optional, tag = "1")]
//     pub params: ::core::option::Option<Params>,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataRequest")]
// #[proto_query(
//     path = "/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata",
//     response_type = QueryDenomAuthorityMetadataResponse
// )]
// pub struct QueryDenomAuthorityMetadataRequest {
//     #[prost(string, tag = "1")]
//     pub denom: ::prost::alloc::string::String,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataResponse")]
// pub struct QueryDenomAuthorityMetadataResponse {
//     #[prost(message, optional, tag = "1")]
//     pub authority_metadata: ::core::option::Option<DenomAuthorityMetadata>,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]

// /// GenesisState defines the tokenfactory module's genesis state.
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.GenesisState")]
// pub struct GenesisState {
//     /// params defines the paramaters of the module.
//     #[prost(message, optional, tag = "1")]
//     pub params: ::core::option::Option<Params>,
//     #[prost(message, repeated, tag = "2")]
//     pub factory_denoms: ::prost::alloc::vec::Vec<GenesisDenom>,
// }
// #[derive(
//     Clone,
//     PartialEq,
//     ::prost::Message,
//     serde::Serialize,
//     serde::Deserialize,
//     schemars::JsonSchema,
//     CosmwasmExt,
// )]
// #[proto_message(type_url = "/osmosis.tokenfactory.v1beta1.GenesisDenom")]
// pub struct GenesisDenom {
//     #[prost(string, tag = "1")]
//     pub denom: ::prost::alloc::string::String,
//     #[prost(message, optional, tag = "2")]
//     pub authority_metadata: ::core::option::Option<DenomAuthorityMetadata>,
// }
pub struct TokenfactoryQuerier<'a> {
    querier: cosmwasm_std::QuerierWrapper<'a, cosmwasm_std::Empty>,
}
impl<'a> TokenfactoryQuerier<'a> {
    pub fn new(querier: cosmwasm_std::QuerierWrapper<'a, cosmwasm_std::Empty>) -> Self {
        Self { querier }
    }
    // pub fn params(&self) -> Result<QueryParamsResponse, cosmwasm_std::StdError> {
    //     QueryParamsRequest {}.query(self.querier)
    // }
    // pub fn denom_authority_metadata(
    //     &self,
    //     denom: ::prost::alloc::string::String,
    // ) -> Result<QueryDenomAuthorityMetadataResponse, cosmwasm_std::StdError> {
    //     QueryDenomAuthorityMetadataRequest { denom }.query(self.querier)
    // }
    // pub fn denoms_from_creator(
    //     &self,
    //     creator: ::prost::alloc::string::String,
    // ) -> Result<QueryDenomsFromCreatorResponse, cosmwasm_std::StdError> {
    //     QueryDenomsFromCreatorRequest { creator }.query(self.querier)
    // }
}

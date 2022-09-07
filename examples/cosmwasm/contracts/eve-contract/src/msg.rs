use cosmwasm_std::Coin;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum ExecuteMsg {
    // CreateDenom {
    //     subdenom: String,
    //     initial_mint: Option<String>,
    //     initial_pool: Option<InitPoolCfg>,
    // },

    CreateDenom {
        owner: String,
        name: String,
        denom: String,
        description: String, 
        precision: i32, 
        max_supply: i32, 
        can_change_max_supply: bool,         
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub struct InitPoolCfg {
    pub swap_fee: String,
    pub exit_fee: String,
    pub pairing_denom: String,
    pub pool_assets: PoolAssests,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub struct PoolAssests {
    pub new_token_amount: String,
    pub new_token_weight: String,
    pub pairing_token_amount: String,
    pub pairing_token_weight: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    QueryTokenCreationFee {},
    QueryCreatorDenoms {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub struct QueryTokenCreationFeeResponse {
    pub fee: Vec<Coin>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub struct QueryCreatorDenomsResponse {
    pub denoms: Vec<String>,
}

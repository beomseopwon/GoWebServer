package dto

type ReqBalanceOfDTO struct {
	ChainName    string `json:"chainName"`
	ContractName string `json:"contractName"`
	Address      string `json:"address"`
}

type ReqTokenDTO struct {
	ChainName    string `json:"chainName"`
	ContractName string `json:"contractName"`
	TokenId      string `json:"tokenId"`
}

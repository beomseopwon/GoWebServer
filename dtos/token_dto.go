package dtos

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

type ReqContractAddressDTO struct {
	ChainName    string `json:"chainName"`
	ContractName string `json:"contractName"`
}

type ResContractAddressDTO struct {
	ContractAddress string `json:"contractAddress"`
}

type ResContractRecoverMessageDTO struct {
	SignHash string `json:"signHash"`
	Address  string `json:"address"`
}

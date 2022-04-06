package v1

import (
	"GoWebServer/client"
	"GoWebServer/config"
	"GoWebServer/dtos"
	c "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Status godoc
// @Summary      nft 컨트랙트 주소
// @Description  nft 컨트랙트 주소
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/contractaddress [post]
func ContractAddress(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data dtos.ReqContractAddressDTO
	json.Unmarshal([]byte(value), &data)
	response, err := (*client.BinderClient()).Call(c.Background(), "contract_contractAddress", data.ChainName, data.ContractName)
	if err != nil {
		fmt.Println("contract_contractAddress err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		resData := &dtos.ResContractAddressDTO{ContractAddress: result}
		// resJson, err := json.Marshal(resData)
		// if err != nil {
		// 	fmt.Println(err)
		// 	context.IndentedJSON(http.StatusInternalServerError, err.Error())
		// }
		// fmt.Println(string(resJson))
		// context.JSON(http.StatusOK, gin.H{
		// 	"contractAddress": result,
		// })
		context.IndentedJSON(http.StatusOK, resData)
	}
}

func MintNFT(context *gin.Context) {
	ContractRecoverMessage(context)
}

type Message struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func ContractRecoverMessage(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var reqData dtos.ResContractRecoverMessageDTO
	json.Unmarshal([]byte(value), &reqData)
	messages := []Message{
		Message{"address", reqData.ContractAddress},
		Message{"uint", "0"},
		Message{"address", reqData.Address},
		Message{"uint", reqData.Nonce},
		Message{"bool", "false"},
	}
	response, err := (*client.BinderClient()).Call(c.Background(), "contract_recoverMessage", messages, reqData.SignHash)
	if err != nil {
		fmt.Println("contract_recoverMessage err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("contract_recoverMessage result", result)
		fmt.Println("contract_recoverMessage Address", reqData.Address)
		if strings.Compare(result, reqData.Address) == 0 {
			fmt.Println("validate address Compare", "true")
			ContractSoliditySHA3(context, reqData.Address, reqData.SignHash, reqData.Nonce)
			return
		}
		context.IndentedJSON(http.StatusBadRequest, "")
	}
}

func ContractSoliditySHA3(context *gin.Context, userAddress string, userSignHash string, userNonce string) {
	messages := [1][2]Message{}
	messages[0][0].Type = "string"
	messages[0][0].Value = "/ftfish0.json"
	messages[0][1].Type = "bytes"
	messages[0][1].Value = userSignHash
	response, err := (*client.BinderClient()).Call(c.Background(), "contract_soliditySHA3", messages)
	if err != nil {
		fmt.Println("contract_soliditySHA3 err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			fmt.Println("contract_soliditySHA3 err ", err.Error())
			context.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("contract_soliditySHA3 sh3", result)
		fmt.Println("contract_soliditySHA3 userSignHash", userSignHash)
		VksSignMessage(context, userAddress, userSignHash, userNonce, result)
	}
}

// func VksSignMessage(context *gin.Context, userAddress string, userSignHash string, soliditySHA3 string) {
// 	response, err := (*client.VKSClient()).Call(c.Background(), "sign_signMessage", soliditySHA3)
// 	if err != nil {
// 		fmt.Println("VksSignMessage err ", err.Error())
// 	} else {
// 		// result, err :=
// 		// if err != nil {
// 		// 	fmt.Println("VksSignMessage err ", err.Error())
// 		// 	context.IndentedJSON(http.StatusInternalServerError, err.Error())
// 		// 	return
// 		// }
// 		fmt.Println("VksSignMessage success ", response.Result)
// 		// var vksResult struct {
// 		// 	Result string `json:"result"`
// 		// }
// 		// json.Unmarshal(res.Body(), &result)
// 		// fmt.Println("VksSignMessage", vksResult.Result)
// 	}
// }

func VksSignMessage(context *gin.Context, userAddress string, userSignHash string, userNonce string, soliditySHA3 string) {

	restyClient := resty.New()
	res, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(config.Config().GetString("wemix.vks.jwt")).
		SetBody(map[string]string{
			"messageHash": soliditySHA3,
		}).
		Post(config.Config().GetString("wemix.vks.url") + "sign/message")

	if err != nil {
		fmt.Println("VksSignMessage err", err.Error())
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var vksResult struct {
		Result string `json:"result"`
	}
	json.Unmarshal(res.Body(), &vksResult)
	fmt.Println("VksSignMessage", vksResult.Result)
	TxSendUnsignedTx(context, "0", userAddress, userSignHash, vksResult.Result, userNonce)
}

// {
//     "jsonrpc": "2.0",
//     "id": "",
//     "method": "tx_sendUnsignedTx",
//     "params": [
//         "tornado",
//         "SampleERC721",
//         "mint",
//         [
//             "0",
//             "0xc32e7a0ba3922fafd516b2cb9fd3ef17e539f8fd",
//             "0",
//             "false",
//             "nft/comm/1900101.json",
//             "0x8c33e0adf7765bd8f58170bc15d3aadfaa4d0b279ef5276362a4ab83ff5e0a383a4ab31e7bd3bed22eb43f17dc42fc28f5e79b512737cb516ebc97f4fd408f8a1b",
//             "0xed8ea4a7d53f23eb9c98af6d1984abae62dcb84e2773d46baed117cc931c1a097733d1e459581224bae082d1ce7a18f45d118b70f10800488aa81b289a18056d1c"
//         ]
//     ]
// }

type MintResult struct {
	Status uint   `json:"status"`
	TxHash string `json:"txhash"`
	Revert Revert `json:"revert"`
}

type Revert struct {
	Message string `json:"message"`
}

func TxSendUnsignedTx(context *gin.Context, fee string, userAddress string, userSign string, vksSign string, nonce string) {
	strParam := []string{
		fee,
		userAddress,
		nonce,
		"false",
		"/ftfish0.json",
		userSign,
		vksSign,
	}
	a, _ := json.Marshal(strParam)
	fmt.Println(string(a))
	response, err := (*client.BinderClient()).Call(c.Background(), "tx_sendUnsignedTx", "tornado", "SampleERC721", "mint", strParam)
	if err != nil {
		fmt.Println("TxSendUnsignedTx err ", err.Error())
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {

		var ma *MintResult
		aerr := response.GetObject(&ma)
		if aerr != nil {
			fmt.Println("TxSendUnsignedTx aerr ", aerr.Error())
		} else {

			fmt.Println("TxSendUnsignedTx Result ", response.Result)
			fmt.Println("TxSendUnsignedTx TxHash ", ma.TxHash)
			fmt.Println("TxSendUnsignedTx Revert.Message ", ma.Revert.Message)
		}

		// result, err := response.GetString()
		// if err != nil {
		// 	fmt.Println("TxSendUnsignedTx err ", err.Error())
		// 	context.IndentedJSON(http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// fmt.Println("TxSendUnsignedTx ssc", result)
		// context.IndentedJSON(http.StatusOK, map[string]string{
		// 	"result": "success",
		// })
	}
}

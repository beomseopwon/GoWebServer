package v1

import (
	"GoWebServer/client"
	"GoWebServer/dtos"
	c "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
		Message{"address", "0x428bf6d6644a57eb4cc393ab8c643e33c9421106"},
		Message{"uint256", "0"},
		Message{"address", "0xf3e2467c29a3d316d3270dd9bfa89acc1878ee84"},
		Message{"uint256", "0"},
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
		if strings.Compare(result, reqData.Address) == 0 {
			fmt.Println("validate address", "true")
			VksSignMessage(context)
		}
		context.IndentedJSON(http.StatusBadRequest, "")
	}
}

func VksSignMessage(context *gin.Context) {

}

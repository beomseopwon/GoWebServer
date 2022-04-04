package v1

import (
	"GoWebServer/client"
	"GoWebServer/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	c "context"

	"github.com/gin-gonic/gin"
)

// Status godoc
// @Summary      서버 상태 체크
// @Description  서버가 살아 있는 확인
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/balanceof [post]
func BalanceOf(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data dto.ReqBalanceOfDTO
	json.Unmarshal([]byte(value), &data)
	response, err := (*client.RPCClient()).Call(c.Background(), "token_balanceOf", data.ChainName, data.ContractName, data.Address)
	if err != nil {
		fmt.Println("token_balanceOf err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		context.String(http.StatusOK, result)
		fmt.Println("token_balanceOf response ", result)
	}
}

// Status godoc
// @Summary      보유중인 모든 토큰
// @Description  보유중인 모든 토큰
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/tokesnofowner [post]
func TokensOfOwner(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data dto.ReqBalanceOfDTO
	json.Unmarshal([]byte(value), &data)
	response, err := (*client.RPCClient()).Call(c.Background(), "token_tokensOfOwner", data.ChainName, data.ContractName, data.Address)
	if err != nil {
		fmt.Println("token_tokensOfOwner err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("token_tokensOfOwner response ", result)
		context.String(http.StatusOK, result)
	}
}

// Status godoc
// @Summary      nft 토큰의 uri 정보
// @Description  nft 토큰의 uri 정보
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/tokenuri [post]
func TokenURI(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data dto.ReqTokenDTO
	json.Unmarshal([]byte(value), &data)
	response, err := (*client.RPCClient()).Call(c.Background(), "token_tokenURI", data.ChainName, data.ContractName, data.TokenId)
	if err != nil {
		fmt.Println("token_tokenURI err ", err.Error())
	} else {
		result, err := response.GetString()
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("token_tokenURI response ", result)
		context.String(http.StatusOK, result)
	}
}

package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BalanceOfData struct {
	ChainName    string `json:"chainName"`
	ContractName string `json:"contractName"`
	Address      string `json:"address"`
}

// Status godoc
// @Summary      서버 상태 체크
// @Description  서버가 살아 있는 확인
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/health [get]
func BalanceOf(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data BalanceOfData
	json.Unmarshal([]byte(value), &data)
	//doc, _ := json.Marshal(body)
	fmt.Println("BalanceOf", string(value))
	context.String(http.StatusOK, "BalanceOf")
}

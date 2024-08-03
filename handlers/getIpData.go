package handlers

import (
	"fmt"
	"ip-data-api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	goip "github.com/jpiontek/go-ip-api"
)

func GetIpData(c *gin.Context) {
	var ipData model.Location
	givenIP := c.Request.URL.Query().Get("ip")
	h := time.Now()
	client := goip.NewClient()

	result, err := client.GetLocationForIp(givenIP)
	if err != nil {
		fmt.Println("Implementar lógica: ", err)
	}

	ipData = model.SetIpData(result, h)

	c.IndentedJSON(http.StatusOK, ipData)
}

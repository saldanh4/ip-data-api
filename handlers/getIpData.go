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
	ipData = model.Location{
		As:          result.As,
		City:        result.City,
		Country:     result.Country,
		CountryCode: result.CountryCode,
		Isp:         result.Isp,
		Lat:         result.Lat,
		Lon:         result.Lon,
		Org:         result.Org,
		Query:       result.Query,
		Region:      result.Region,
		RegionName:  result.RegionName,
		Status:      result.Status,
		Timezone:    result.Timezone,
		Zip:         result.Zip,
		TimeStamp:   h.Format("02/Jan/2006 15:04:05"),
	}

	c.IndentedJSON(http.StatusOK, ipData)
}

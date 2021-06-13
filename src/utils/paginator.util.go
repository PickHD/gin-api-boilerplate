package utils

import (
	"strconv"
	"math"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	//*Initializing default
	limit := 5
	page := 1

	query := c.Request.URL.Query()

	for key, value := range query {

		queryValue := value[len(value)-1]

		switch key {

		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
			
		}
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}

}

func CountTotalPage(totalData int,limit *Pagination) int {
	return int(math.Ceil(float64(totalData)/float64(limit.Limit)))
}
package examples

import (
	"net/http"

	"gin-api-boilerplate/src/utils"

	"github.com/gin-gonic/gin"
)

//*Example creating a controllers
func (e *ExampleBase) HelloWorldController(c *gin.Context) {

	//*Call services here
	getHello:=PrintHelloWorld()

	utils.ResponseFormatter(http.StatusOK,getHello,nil,nil,c)

}
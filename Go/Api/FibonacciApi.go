package api

import (
	"Fibonacci/Common"
	"Fibonacci/Loop"
	"Fibonacci/Recursion"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const loopMethod string = "loop"
const recursiveMethod string = "recursion"

func GetNumbers(ctx *gin.Context) {
	var count int = getCount(ctx)
	var method string = getMethod(ctx)

	var service Common.IFibonacciService = getService(method)
	var numbers []int = service.GetNumbers(count)

	ctx.IndentedJSON(http.StatusOK, numbers)
}

func getCount(ctx *gin.Context) int {
	qsValue := ctx.Query("count")
	if qsValue == "" {
		return 5
	} else {
		var count int
		var err error

		count, err = strconv.Atoi(qsValue)
		if err == nil {
			return count
		} else {
			return 5
		}
	}
}

func getMethod(ctx *gin.Context) string {
	qsValue := ctx.Query("method")

	if qsValue == "" {
		return recursiveMethod
	}

	return qsValue
}

func getService(method string) Common.IFibonacciService {
	formattedMethod := strings.ToLower(method)

	if formattedMethod == loopMethod {
		return Loop.LoopFibonacciService{}
	} else {
		return Recursion.RecursiveFibonacciService{}
	}
}
package middlewares

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var stored float64 = 0 // Represent request stored at bucket

const bucket_size = 10 // Represent volume bucket capacity

var l_time = time.Now() // It's used to measure elapsed time

const tokenPerPeriod = 4 // Represent quantity throttle allows per second

func LeakyBucket(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//Verify if bucket is full
		tm := time.Now()
		elapsed := tm.Sub(l_time)
		l_time = tm
		var consumes = elapsed.Seconds() * tokenPerPeriod
		fmt.Println("Token consumed by request ", consumes)
		fmt.Println("Token stored ", stored)
		stored = math.Max(0, stored-consumes)
		fmt.Println("Elapsed time ", elapsed)
		if (stored <= bucket_size) && (consumes <= tokenPerPeriod) {
			fmt.Println("Accept with tokens", stored)
			stored++
			return next(c)

		} else {
			fmt.Println("Denied 429")
			return c.JSON(http.StatusTooManyRequests, 429)
		}

	}
}

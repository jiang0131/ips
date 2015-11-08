package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

func loggerMiddleware() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {

		start := time.Now().UTC()
		logData := make(map[string]interface{})

		logData["reqMethod"] = req.Method
		logData["reqUrl"] = req.URL.Path

		if len(req.URL.Query()) != 0 {
			logData["reqQuery"] = req.URL.Query()
		}

		c.Next()

		rw := res.(martini.ResponseWriter)

		logData["resTimeMs"] = float64(int(time.Since(start).Seconds()*1000000)) / 1000
		logData["httpStatus"] = rw.Status()

		if rw.Status() >= http.StatusBadRequest {
			logger.error(logData)
			return
		}

		logger.info(logData)
	}
}

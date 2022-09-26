package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/logger"
	"io"
	"io/ioutil"
	"time"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logger.GetLogger()
		blw := &responseBodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		start := time.Now()
		c.Next()

		// Excluding health check request from logs.
		if c.Request.RequestURI == "/hc" {
			return
		}

		headers, _ := json.Marshal(c.Request.Header)
		requestBody := readRequestBody(c)
		responseBody := blw.body.String()

		if c.Writer.Status() > 299 {
			l.Error().
				Dur("duration", time.Since(start)).
				Str("path", c.Request.RequestURI).
				Str("method", c.Request.Method).
				Str("request_headers", string(headers)).
				Str("request_body", requestBody).
				Str("response_body", responseBody).
				Int("status", c.Writer.Status()).
				Str("client_ip", c.ClientIP()).
				Str("user_agent", c.Request.UserAgent()).
				Str("use-case", "infra/http-request").
				Msg("Unsuccessful request at" + c.Request.RequestURI)
		} else {
			l.Error().
				Dur("duration", time.Since(start)).
				Str("path", c.Request.RequestURI).
				Str("method", c.Request.Method).
				Str("request_headers", string(headers)).
				Str("request_body", requestBody).
				Str("response_body", responseBody).
				Int("status", c.Writer.Status()).
				Str("client_ip", c.ClientIP()).
				Str("user_agent", c.Request.UserAgent()).
				Str("use-case", "infra/http-request").
				Msg("Successful request at" + c.Request.RequestURI)
		}
	}
}

func readRequestBody(c *gin.Context) (rb string) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	rb = string(body)
	c.Request.Body = ioutil.NopCloser(&buf)
	return
}

type responseBodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

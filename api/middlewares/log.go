package middlewares

import (
	"bytes"
	"net/http"
	"time"

	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/logging"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//Logging ...
func Logging(restricteds ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer recovery()
		defer c.Request.Body.Close()

		requestID, _ := uuid.NewV4()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		blw.Header().Set("Request-Id", requestID.String())
		c.Writer = blw

		now := time.Now()
		req := request(c, restricteds)

		c.Next()

		elapsed := time.Since(now)
		fields := make(map[string]interface{})

		fields["request"] = req
		fields["claims"] = c.Keys
		fields["errors"] = c.Errors
		fields["ip"] = c.ClientIP()
		fields["latency"] = elapsed.Seconds()
		fields["request_id"] = requestID.String()
		fields["response"] = response(blw, restricteds)

		logging.LogWith(fields).Info(
			"Request incoming from %s elapsed %s completed with %d",
			c.ClientIP(),
			elapsed.String(),
			c.Writer.Status(),
		)
	}
}

func request(context *gin.Context, restricteds []string) interface{} {
	r := make(map[string]interface{})

	headers := sanitizeHeader(context.Request.Header)

	r["headers"] = headers
	r["host"] = context.Request.Host
	r["form"] = context.Request.Form
	r["path"] = context.Request.URL.Path
	r["method"] = context.Request.Method
	r["url"] = context.Request.URL.String()
	r["body"] = context.Request.Body
	r["post_form"] = context.Request.PostForm
	r["remote_addr"] = context.Request.RemoteAddr
	r["query_string"] = context.Request.URL.Query()

	return r
}

func response(writer *bodyLogWriter, restricteds []string) interface{} {
	r := make(map[string]interface{})

	r["body"] = writer.body
	r["status"] = writer.Status()
	r["headers"] = writer.Header()

	return r
}

func sanitizeHeader(header http.Header) http.Header {
	cp := make(http.Header)

	for k, v := range header {
		cp[k] = v
	}

	delete(cp, "Authorization")

	return cp
}

func recovery() {
	if err := recover(); err != nil {
		logging.LogWith(err).Error("Error on logging middleware")
	}
}

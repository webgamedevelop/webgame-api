package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func InspectRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if klog.V(klog.Level(inspectLevel)).Enabled() {
			var out bytes.Buffer
			if c.Request.Method == http.MethodPost {
				if c.Request.Body != nil {
					body, err := c.GetRawData()
					if err != nil {
						klog.Error(err, "read request body failed")
					}

					if len(body) > 0 {
						c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
					}

					if err := json.Compact(&out, body); err != nil {
						klog.Error(err, "read request body failed")
					}
				}
			}
			klog.V(klog.Level(inspectLevel)).Info("inspect request", "method", c.Request.Method, "uri", c.Request.RequestURI, "body", out.String())
		}
		c.Next()
	}
}

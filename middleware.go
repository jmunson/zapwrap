package zapwrap

//inspired by https://github.com/echo-contrib/echo-logrus but with zap instead
import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/uber-go/zap"
)

func NewMiddleware() echo.MiddlewareFunc {
	return NewMiddlewareWithLogger(zap.NewJSON())
}

// NewMiddlewareWithLogger returns a new logging middleware
func NewMiddlewareWithLogger(l zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			entry := l.With(
				zap.String("request", c.Request().URI()),
				zap.String("method", c.Request().Method()),
				zap.String("remote", c.Request().RemoteAddress()),
			)

			if reqID := c.Request().Header().Get("X-Request-Id"); reqID != "" {
				entry = entry.With(zap.String("request_id", reqID))
			}

			entry.Info("started handling request")
			c.Set("log", entry) //Experiment.. storing contextual logger in context

			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)
			entry = c.Get("log").(zap.Logger) //Get log back out of context, so that we include any fields they've included.

			entry.With(
				zap.Int("status", c.Response().Status()),
				zap.String("text_status", http.StatusText(c.Response().Status())),
				zap.Duration("took", latency)).
				Info("completed handling request")

			return nil
		}
	}
}

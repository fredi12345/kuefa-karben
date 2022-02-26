package extensions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type LoggerConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper middleware.Skipper

	// Logger is the underlying zap.Logger.
	Logger *zap.Logger
}

func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}
	if config.Logger == nil {
		config.Logger, _ = zap.NewProduction()
	}
	config.Logger.Named("request")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			res := c.Response()

			contentLength, _ := strconv.ParseInt(req.Header.Get(echo.HeaderContentLength), 10, 64)
			msg := fmt.Sprintf(
				"%s %s %d: %s in %s",
				req.Method,
				req.RequestURI,
				res.Status,
				byteCountSI(res.Size),
				stop.Sub(start).String(),
			)
			config.Logger.Info(msg,
				zap.Int64("latency", stop.Sub(start).Milliseconds()),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.Int("status", res.Status),
				zap.Int64("bytes_in", contentLength),
				zap.Int64("bytes_out", res.Size),
			)

			return err
		}
	}
}

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

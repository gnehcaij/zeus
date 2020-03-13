package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gnehcaij/zeus/constant"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LogMiddleware() gin.HandlerFunc {

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	apiLogPath := "output/log/info.log"
	errorLogPath := "output/log/error.log"
	infoLogWriter, _ := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	errorLogWriter, _ := rotatelogs.New(
		errorLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(errorLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  infoLogWriter,
		logrus.FatalLevel: errorLogWriter,
		logrus.ErrorLevel: errorLogWriter,
		logrus.WarnLevel:  errorLogWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{})
	logrus.AddHook(lfHook)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP, _ := c.Get(constant.LOCAL_IP_KEY)
		logId, _ := c.Get(constant.LOG_ID)

		logrus.WithFields(logrus.Fields{
			"logId":       logId,
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIp":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
	}
}

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotlog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	filePath := "log/log"
	//linkName := "latest_log.log"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = file // 记录日志至文件夹中log/log.log
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotlog.New(
		filePath+"%Y%m%d.log",
		rotlog.WithMaxAge(7*24*time.Hour),
		rotlog.WithRotationTime(24*time.Hour),
		//rotlog.WithLinkName(linkName), // 创建软链接（需要管理员权限），而且是最新文件链接。避免报错，这里注释
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.WarnLevel:  logWriter,
	}
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 执行下一个中间件，执行完所有的中间件之后会继续执行next后面的函数
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.URL

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIP,
			"Method":    method,
			"path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

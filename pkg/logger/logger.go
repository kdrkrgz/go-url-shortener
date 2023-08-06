package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init initializes the logger.
func init() {
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout"}
	conf.Level.SetLevel(zap.InfoLevel)
	if os.Getenv("Environment") == "development" {
		conf.Level.SetLevel(zap.DebugLevel)
	}
	conf.Level.SetLevel(zap.DebugLevel)
	conf.EncoderConfig.TimeKey = "timestamp"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := conf.Build()
	if err != nil {
		log.Fatalf("Error initializing logger: %s", err)
	}
	logger = zapLogger.With(zap.String("service", "go-url-shortener"))
}

func Logger() *zap.Logger {
	return logger
}

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
func twoSum(nums []int, target int) []int {
	nmap := map[int]int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		nmap[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		complement := target - nums[i]
		for k, _ := range nmap {
			if k == complement && nmap[complement] != i {
				return []int{i, nmap[complement]}
			}
		}
	}
	return []int{}
}

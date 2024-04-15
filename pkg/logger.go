package pkg

import (
	"fmt"
	"log"
	"os"

	c "github.com/edB96/isaac/cmd/constants"
)

type Level int8

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

var logger *log.Logger
var level = c.DefaultLogLevel

func SetupLogger() {
	logger = log.New(os.Stdout, "", log.Ltime)
}

// DEBUG (not colorized)
func Debug(format string, vars ...interface{}) {
	if level <= "DEBUG" {
		logger.Printf(c.LogPrefixDebug+format, vars...)
	}
}

// INFO (cyan)
func Info(format string, vars ...interface{}) {
	if level <= "INFO" {
		logger.Printf(c.TextColorCyan, c.LogPrefixInfo+fmt.Sprintf(format, vars...))
	}
}

// WARNING (yellow)
func Warn(format string, vars ...interface{}) {
	if level <= "WARNING" {
		logger.Printf(c.TextColorYellow, c.LogPrefixWarning+fmt.Sprintf(format, vars...))
	}
}

// ERROR (red)
func Error(format string, vars ...interface{}) {
	if level <= "ERROR" {
		logger.Printf(c.TextColorRed, c.LogPrefixError+fmt.Sprintf(format, vars...))
	}
}

package logger

import (
	"fmt"
)

type LogrusLogger struct {
}

func NewLogrusLogger() *LogrusLogger {
	_logger := &LogrusLogger{}
	return _logger
}

func (log *LogrusLogger) Info(msg string) {
	fmt.Println("log Info:", msg)
}

func (log *LogrusLogger) Debug(msg string) {
	fmt.Println("log Debug:", msg)
}
func (log *LogrusLogger) Warn(msg string) {
	fmt.Println("log Warn:", msg)
}
func (log *LogrusLogger) Error(msg string) {
	fmt.Println("log Error:", msg)

}
func (log *LogrusLogger) Fatal(msg string) {
	fmt.Println("log Fatal:", msg)
}

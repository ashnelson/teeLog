package teeLog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	logs "github.com/sirupsen/logrus"
)

type jsonFormatter struct{}

func (this *jsonFormatter) Format(entry *logs.Entry) ([]byte, error) {
	type MyEntry struct {
		Level      string    `json:"level"`
		Timestamp  time.Time `json:"time"`
		StackTrace []string  `json:"message"`
	}

	myEntry := MyEntry{
		Level:      entry.Level.String(),
		Timestamp:  entry.Time,
		StackTrace: strings.Split(entry.Message, "\n"),
	}

	serialized, err := json.Marshal(myEntry)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

type StdOutFormatter struct{}

func (this *StdOutFormatter) Format(entry *logs.Entry) ([]byte, error) {
	formatStr := fmt.Sprintf("level: %s; time: %d\n\t%s\n", entry.Level, entry.Time.UnixNano(), strings.Replace(entry.Message, "\n", "\n\t", -1))
	return []byte(formatStr), nil
}

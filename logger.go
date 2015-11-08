package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"time"
)

type StenoJSONFormatter struct{}

type logWrapper struct{}

func (f *StenoJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := entry.Data
	entry.Data = logrus.Fields{}

	entry.Data["data"] = data
	entry.Data["time"] = time.Now().UTC()
	entry.Data["level"] = entry.Level.String()
	entry.Data["name"] = SERVICE_NAME

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marchal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

var log = func() *logrus.Logger {
	log := logrus.New()

	log.Formatter = new(StenoJSONFormatter)

	level, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	log.Level = level
	return log
}()

var logger = new(logWrapper)

func (l *logWrapper) info(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Info("") })
}

func (l *logWrapper) warn(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Warn("") })
}

func (l *logWrapper) error(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Error("") })
}

func (l *logWrapper) debug(data map[string]interface{}) {
	l.log(data, func(e *logrus.Entry) { e.Debug("") })
}

func (l *logWrapper) log(data map[string]interface{}, fn func(*logrus.Entry)) {
	f, err := os.OpenFile(LOG_FILE_PATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	log.Out = f

	if err != nil {
		panic(err)
	}

	fn(log.WithFields(data))
}
package heartbeat

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

const (
  FileKey     = "HEARTBEAT_FILE"
  IntervalKey = "HEARTBEAT_INTERVAL"
)

// RunFileHeartbeat loops heartbeat to file emition in
// an infinite loop, with a specified interval.
func RunFileHeartbeat(filepath string, sleepSeconds int) {
  for {
    EmitToFile(filepath)
    time.Sleep(sleepSeconds * time.Second)
  }
}

// EmitToFile Changes the time modified on a specified file,
// creates the file if not present.
func EmitToFile(filepath string) {
  now := getNow()
  err := os.Chtimes(filepath, now, now)
  if err != nil {
    log.Println(err)
    createFile(filepath)
  }
}

// Config configuration for running file heartbeat.
type Config struct {
  File     string `json:"file"`
  Interval int    `json:"interval"`
}

// NewConfigFromEnv reads heartbeat configuration from environment variables
func NewConfigFromEnv() (Config, error) {
  config := Config{File: os.Getenv(FileKey)}
  if config.File == "" {
    return config, fmt.Error("Heartbeat Config.File not set")
  }
  interval, err := strconv.Atoi(os.Getenv(IntervalKey))
  if err != nil {
    return config, err
  }
  config.Interval = interval
  return config, nil
}

// createFile creates a new file named as specified by the filepath.
func createFile(filepath string) {
  f, err := os.Create(filepath)
  if err != nil {
    log.Fatal(err)
  }
  f.Close()
}

// getNow gets the current UTC timestamp.
func getNow() time.Time {
  return time.Now().UTC()
}

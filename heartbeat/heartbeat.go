package heartbeat

import (
  "os"
  "time"
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

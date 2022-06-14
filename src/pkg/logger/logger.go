package logger

import (
  "log"
  "os"
)

var (
  WarningLogger *log.Logger
  InfoLogger *log.Logger
  ErrorLogger *log.Logger
)

func init() {
  InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
  ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
  InfoLogger.Println("Starting Logger...")
}
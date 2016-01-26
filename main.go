package main

import (
  "fmt"
  "os"
  "io"
  "errors"
  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/saintfish/chardet"
)

var (
  app = kingpin.New("icu", "Detect whatsoever charset")
  bufferSize = app.Flag("buffer", "Buffer size of the file.").Default("100000").Int()
  forText = app.Flag("text", "Enable text more, default is HTML mode.").Bool()
  fileName = app.Arg("file", "Input file name, or you can use pipe or redirection.").String()
)

func main() {
  app.Parse(os.Args[1:])
  app.Version("0.1.0")

  inputFile := os.Stdin
  err := errors.New("")

  inputFileName := *fileName
  if inputFileName != "" {
    inputFile, err = os.Open(inputFileName)
    if err != nil {
      fmt.Println(err)
      os.Exit(-1)
    }
  }

  buffer := make([]byte, *bufferSize)
  size, _ := io.ReadFull(inputFile, buffer)

  input := buffer[:size]
  detector := chardet.NewHtmlDetector()
  if *forText == true {
    detector = chardet.NewTextDetector()
  }
  detectResult, err := detector.DetectBest(input)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  fmt.Print(detectResult.Charset)
}

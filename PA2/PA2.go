package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
  inputFilename := ""
  fmt.Printf("Please input the filename you want to read: ")
  fmt.Scanf("%s", &inputFilename)

  outputFilename := ""
  fmt.Printf("Please input the filename you want to output: ")
  fmt.Scanf("%s", &outputFilename)
	
  inputFile, err := os.Open(inputFilename)
	check(err)
	scanner := bufio.NewScanner(inputFile)

  outputFile, err := os.Create(outputFilename)
  check(err)
  writer := bufio.NewWriter(outputFile)
	
  defer inputFile.Close()
  defer outputFile.Close()

  strBuffer := ""
  i := 1
  for scanner.Scan() {
    strBuffer = scanner.Text()
    writer.WriteString(strconv.Itoa(i)+ " " + strBuffer +"\n")
    i++
	}
  writer.Flush()

}

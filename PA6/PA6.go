package main
import "fmt"
import "bufio"
import "net"
import "strconv"
import "os"
import "time"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func handleConnection (c net.Conn) {
    outputFile, err := os.Create("whatever.txt")
    check(err)
    writer := bufio.NewWriter(outputFile)
    defer outputFile.Close()


   reader := bufio.NewReader(c) // define a reader
   message, errr := reader.ReadString('\n') // get file size from client
   check(errr)


   fmt.Printf("Upload file size: %s", message)
   fileSize, err := strconv.Atoi(message[:len(message) - 1])
   check(err)


   remainBytes := fileSize
   i := 1
   for{
    message, errr := reader.ReadString('\n') // get file size
    check(errr)

    // fmt.Printf(message)
    remainBytes -= len(message)
       _, errs := writer.WriteString(strconv.Itoa(i) + " " + message)
    check(errs)
    i+=1
    if(remainBytes <= 0){
        break
    }
   }
    writer.Flush()
   newFile, err := os.Open("whatever.txt")
   check(err)
   fi, _ := newFile.Stat()
   defer newFile.Close()

   fmt.Printf("Output file size: %s\n", strconv.Itoa(int(fi.Size())))

   writerToClient := bufio.NewWriter(c)
    _, errw := writerToClient.WriteString("original file size: " + strconv.Itoa(fileSize) + " / new file size: " + strconv.Itoa(int(fi.Size())) + "\n")
    writerToClient.Flush()
    check(errw)

    time.Sleep(5 * time.Second)
}


func main() {
    fmt.Println("Launching server...")
    ln, _ := net.Listen("tcp", ":12005")
    defer ln.Close()
    for {
        conn, _ := ln.Accept()
        defer conn.Close()
        handleConnection(conn) 
    } 
}



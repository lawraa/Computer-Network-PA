package main
import "fmt"
import "bufio"
import "net"
import "net/http"
import "strconv"
import "os"
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
} 

func handleConnection (c net.Conn) {
	reader := bufio.NewReader(c)
	req, err := http.ReadRequest(reader)
	check(err)
	filePath := "." + req.URL.String()
	fileExist, _ := exists(filePath)
	if(!fileExist){
		fmt.Println("file not found")
	}else{
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()
	fi, _ := f.Stat()
	fmt.Println(strconv.Itoa(int(fi.Size())))
	}
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



package main 
import "fmt"
import "os"
import "strconv"
import "net"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	//connects to the server
	conn, errc := net.Dial("tcp", "127.0.0.1:1000")
	check(errc)
	defer conn.Close()
	
	
	// prompts the user for the upload filename
	inputFilename := ""
	fmt.Printf("Please input the filename you want to read: ")
	fmt.Scanf("%s", &inputFilename)
	
	inputFile, err := os.Open(inputFilename)
	check(err)
	defer inputFile.Close()
	
	// start up an output buffer for the tcp socket("conn")
	writer := bufio.NewWriter(conn)
	// open input buffer for file io
	scanner := bufio.NewScanner(inputFile)
	
	
	// send file size
	fi, _ := inputFile.Stat()
	fmt.Printf("from client: " + strconv.Itoa(int(fi.Size())) + "\n")
	_, errw := writer.WriteString(strconv.Itoa(int(fi.Size())) + "\n")
	check(errw)
	
	// read file
	for scanner.Scan() {	
		_, errs := writer.WriteString(scanner.Text() +"\n")
		check(errs)
	}
	// end of buffer writing
	writer.Flush()
	
	// receive a msg from server
	scanner = bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server replies: %s\n", scanner.Text())
	}
}

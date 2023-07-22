package main 
import "fmt"
import "os"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	//connects to the server


	// prompts the user for the upload filename
	inputFilename := ""
	fmt.Printf("Please input the filename you want to read: ")
	fmt.Scanf("%s", &inputFilename)

	inputFile, err := os.Open(inputFilename)
	check(err)
	defer inputFile.Close()




	// send file size
	fi, _ := inputFile.Stat()
	print(strconv.Itoa(int(fi.Size())))


	// read file

}
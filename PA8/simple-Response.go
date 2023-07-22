package main
import "fmt"
import "net/http"
func main() {
 fmt.Println("Launching server...")
 http.ListenAndServe(":12005", http.FileServer(http.Dir(".")))
}
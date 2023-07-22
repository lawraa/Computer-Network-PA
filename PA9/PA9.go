package main
import "fmt"
import "net/http"
import "os"
import "strings"
import "net/url"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func fileExist(p string)(bool){
	filePath := "./" + p
	fileExist, _ := exists(filePath)
	if(!fileExist){
		return false
	}else{
		return true
	}
}

func myStripPrefix(prefix string, h Handler) Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			if fileExist(p){
				r2.URL.RawPath = rp
				r2.URL.Path = p
				h.ServeHTTP(w, r2)
				}else{
					fmt.Fprintln(w, "File not found")
				}
		}else {
			fmt.Fprintln(w, "File not found")
		}
	})
}

func main(){
// conn for replying "file not found" 
 fmt.Println("Launching server...")

 fs := http.FileServer(http.Dir("."))
 http.Handle("/", myStripPrefix("/", fs))
 http.ListenAndServeTLS(":12005","server.cer", "server.key", nil)

} 
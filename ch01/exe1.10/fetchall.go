// fetchall; get URL using goroutine
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	filename := `./fatchlog_` + fmt.Sprint(time.Now().Unix())
	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("err %v : cannot create file:%s \n", err, filename)
		return
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start goroutine
	}
	for range os.Args[1:] {
		io.WriteString(out, <-ch)
		//fmt.Println(<-ch) // reserve form channel
	}
	io.WriteString(
		out,
		fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()),
	)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // sent channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // resource release
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s \n", secs, nbytes, url)
}

package main

import (
	"fmt"
	"net/http"
	"path"
	"os"
	"io"
)

func main() {
	fmt.Println(fetch("https://github.com/Yuicon"))
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	var closeErr error
	defer func() { closeErr = f.Close() }()
	// Close file, but prefer error from Copy, if any.
	if err == nil {
		err = closeErr
	}
	return local, n, err
}

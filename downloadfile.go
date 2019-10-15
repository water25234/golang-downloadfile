package main

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

var (
    fileName    string
    fullUrlFile string
    theArray [5]string
)

func theArrayList() {
    theArray[0] = "http://www.golangprograms.com/skin/frontend/base/default/logo.jpg"
    theArray[1] = "https://golangcode.com/images/avatar.jpg"
}

func main() {

    theArrayList()

	for i := 0; i < len(theArray); i++ {

        if i%100 == 0 && i > 0 {
            fmt.Println("機器累了休息中")
            time.Sleep(time.Duration(2)*time.Second)
            fmt.Println("機器上工囉～")
        }

		fullUrlFile = theArray[i]
		// Build fileName from fullPath
		buildFileName()

		// Create blank file
		file := createFile()

		// Put content on file
		putFile(file, httpClient())
	}
}

func putFile(file *os.File, client *http.Client) {
    resp, err := client.Get(fullUrlFile)

    checkError(err)

    defer resp.Body.Close()

    size, err := io.Copy(file, resp.Body)

    defer file.Close()

    checkError(err)

    fmt.Println("Just Downloaded a file", fileName, "with size", size)
}

func buildFileName() {
    fileUrl, err := url.Parse(fullUrlFile)
    checkError(err)

    path := fileUrl.Path
    segments := strings.Split(path, "/")

    fileName = segments[len(segments)-1]

    var fileNameSplit = strings.Split(fileName, ".")

    if len(fileNameSplit) == 1 {
        fileName = strings.Join([]string{fileName,"png"}, ".")
    }
}

func httpClient() *http.Client {
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }

    return &client
}

func createFile() *os.File {
    file, err := os.Create(fileName)

    checkError(err)
    return file
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

package upload

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

const maxUploadSize = 2 * 1024 * 1024   // 2 MB 

func UpdateBlockHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(maxUploadSize)
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    dst, err := os.Create(handler.Filename)
	defer dst.Close()
    // dst, err := CreateFile()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
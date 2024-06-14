package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Create directory if not exists
    uploadDir := "./uploads"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.Mkdir(uploadDir, os.ModePerm)
    }

    // Create a new file in the uploads directory
    filePath := filepath.Join(uploadDir, handler.Filename)
    destFile, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Unable to create the file", http.StatusInternalServerError)
        return
    }
    defer destFile.Close()

    // Copy the uploaded file data to the destination file
    if _, err := io.Copy(destFile, file); err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Successfully uploaded file: %s\n", handler.Filename)
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

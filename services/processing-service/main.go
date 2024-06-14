package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "path/filepath"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
    fileID := r.URL.Query().Get("file_id")
    if fileID == "" {
        http.Error(w, "Missing file_id parameter", http.StatusBadRequest)
        return
    }

    // Dummy processing: just rename the file to add a processed prefix
    srcPath := filepath.Join("./uploads", fileID)
    destPath := filepath.Join("./processed", "processed-"+fileID)

    // Create directory if not exists
    if _, err := exec.Command("mkdir", "-p", "./processed").Output(); err != nil {
        http.Error(w, "Unable to create processed directory", http.StatusInternalServerError)
        return
    }

    // Rename the file to simulate processing
    if err := exec.Command("mv", srcPath, destPath).Run(); err != nil {
        http.Error(w, "Error processing the file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Successfully processed file: %s\n", fileID)
}

func main() {
    http.HandleFunc("/process", processHandler)
    fmt.Println("Server started at :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

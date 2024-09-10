package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Struct Tugas
type Tugas struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}

// Data Tugas dengan Title "Hello World"
var tugas = []Tugas{
    {ID: "1", Title: "Hello World"},
    {ID: "2", Title: "Hello World"},
}

// Fungsi untuk mengambil semua tugas
func getTugas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, tugas)
}

func main() {
    // Membuat router gin
    router := gin.Default()

    // Endpoint GET untuk mengambil semua tugas
    router.GET("/tugas", getTugas)

    // Menjalankan server pada port 8080
    router.Run("localhost:8080")
}

package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "html/template"
    "strconv"
)

var numeroGenerato int
var count int

func main() {
    rand.Seed(int64(rand.Intn(100)))
    numeroGenerato = rand.Intn(100)
    count = 0

    // Imposta il routing
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/tentativo", tentativoHandler)

    // Prendi la porta da Render
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Imposta una porta di fallback per il locale
    }

    // Avvia il server
    fmt.Printf("Il gioco è avviato su http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, nil)
}

// Gli altri handler (homeHandler, tentativoHandler) restano invariati

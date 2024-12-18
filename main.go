package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var prices = map[string]int{
	"patatine": 2,
	"hamburger": 5,
	"cocacola": 2,
}

// Struttura per gestire l'ordine
type Order struct {
	Patatine int
	Hamburger int
	Cocacola int
	Total int
	PatatineTimesPrice int
	HamburgerTimesPrice int
	CocacolaTimesPrice int
}

func main() {
	// Prendi la porta dalla variabile di ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Imposta la porta di default se non trovata
	}

	// Route per servire la pagina principale
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="it">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Fast Food Ordinazione</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f0f0f0;
            color: #333;
            text-align: center;
        }
        h1 {
            color: #FF5733;
        }
        .menu {
            margin-top: 20px;
        }
        .menu-item {
            background-color: #fff;
            border: 1px solid #ddd;
            margin: 10px;
            padding: 15px;
            border-radius: 8px;
            box-shadow: 0 0 5px rgba(0,0,0,0.1);
        }
        .menu-item h3 {
            color: #333;
        }
        .menu-item button {
            background-color: #FF5733;
            color: white;
            border: none;
            padding: 10px;
            margin: 5px;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
        }
        .menu-item button:hover {
            background-color: #ff4500;
        }
        .order-summary {
            margin-top: 30px;
            background-color: #fff;
            border: 1px solid #ddd;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 5px rgba(0,0,0,0.1);
            font-size: 18px;
        }
        .total-price {
            font-size: 24px;
            color: #FF5733;
            margin-top: 15px;
        }
    </style>
</head>
<body>

    <h1>Benvenuto nel nostro Fast Food!</h1>
    <p>Seleziona gli articoli e aggiungi la quantità</p>

    <form action="/order" method="POST">
        <div class="menu">
            <div class="menu-item">
                <h3>Patatine - 2€</h3>
                <input type="number" name="patatine" value="0" min="0">
            </div>
            <div class="menu-item">
                <h3>Hamburger - 5€</h3>
                <input type="number" name="hamburger" value="0" min="0">
            </div>
            <div class="menu-item">
                <h3>Coca-Cola - 2€</h3>
                <input type="number" name="cocacola" value="0" min="0">
            </div>
        </div>

        <button type="submit">Ordina</button>
    </form>

    {{if .Total}}
        <div class="order-summary">
            <h3>Riepilogo ordine:</h3>
            <p>Patatine: {{.Patatine}} x 2€ = {{.PatatineTimesPrice}}€</p>
            <p>Hamburger: {{.Hamburger}} x 5€ = {{.HamburgerTimesPrice}}€</p>
            <p>Coca-Cola: {{.Cocacola}} x 2€ = {{.CocacolaTimesPrice}}€</p>
            <p class="total-price">Totale: {{.Total}}€</p>
        </div>
    {{end}}

</body>
</html>
`)

		if err != nil {
			http.Error(w, "Errore nel caricamento della pagina", http.StatusInternalServerError)
			return
		}

		// Imposta l'ordine a zero
		order := Order{}

		// Verifica se è presente un ordine (POST)
		if r.Method == "POST" {
			// Prendi i dati dalla form
			patatineQty, _ := strconv.Atoi(r.FormValue("patatine"))
			hamburgerQty, _ := strconv.Atoi(r.FormValue("hamburger"))
			cocacolaQty, _ := strconv.Atoi(r.FormValue("cocacola"))

			// Calcola il totale
			order.Patatine = patatineQty
			order.Hamburger = hamburgerQty
			order.Cocacola = cocacolaQty
			order.Total = (order.Patatine * prices["patatine"]) + (order.Hamburger * prices["hamburger"]) + (order.Cocacola * prices["cocacola"]) + 2
		}

		// Calcola i valori per la pagina
		order.PatatineTimesPrice = order.Patatine * prices["patatine"]
		order.HamburgerTimesPrice = order.Hamburger * prices["hamburger"]
		order.CocacolaTimesPrice = order.Cocacola * prices["cocacola"]

		// Esegui il template con i dati dell'ordine
		tmpl.Execute(w, order)
	})

	// Avvia il server sulla porta dinamica
	fmt.Println("Server in esecuzione sulla porta", port)
	http.ListenAndServe(":"+port, nil)
}

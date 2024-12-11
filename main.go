package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var numeroGenerato int
var count int

func main() {
	// Inizializza il numero casuale
	rand.Seed(int64(rand.Intn(100)))
	numeroGenerato = rand.Intn(100)
	count = 0

	// Imposta il routing
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tentativo", tentativoHandler)

	// Avvia il server
	fmt.Println("Il gioco è avviato! Vai su http://localhost:8080 per iniziare.")
	http.ListenAndServe(":8080", nil)
}

// Home Handler che mostra il form iniziale
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("home").Parse(`
	<!DOCTYPE html>
	<html lang="it">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Gioco del Numero</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f4f4f9;
				color: #333;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
			}
			.container {
				text-align: center;
				background-color: #fff;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				width: 100%;
				max-width: 400px;
			}
			h1 {
				font-size: 2em;
				color: #2c3e50;
			}
			form {
				margin-top: 20px;
			}
			input[type="number"] {
				padding: 10px;
				font-size: 1.2em;
				width: 80%;
				border: 1px solid #ccc;
				border-radius: 5px;
			}
			button {
				padding: 10px 20px;
				font-size: 1.2em;
				border: none;
				border-radius: 5px;
				background-color: #3498db;
				color: white;
				cursor: pointer;
				margin-top: 10px;
			}
			button:hover {
				background-color: #2980b9;
			}
			a {
				display: inline-block;
				margin-top: 20px;
				color: #3498db;
				text-decoration: none;
				font-size: 1.1em;
			}
			a:hover {
				text-decoration: underline;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Indovina il Numero!</h1>
			<p>Prova a indovinare il numero tra 0 e 99.</p>
			<form action="/tentativo" method="post">
				<input type="number" id="tentativo" name="tentativo" required>
				<button type="submit">Invia</button>
			</form>
		</div>
	</body>
	</html>
	`)
	tmpl.Execute(w, nil)
}

// Tentativo Handler che gestisce il gioco
func tentativoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tentativo, err := strconv.Atoi(r.FormValue("tentativo"))
	if err != nil {
		http.Error(w, "Inserisci un valore valido.", http.StatusBadRequest)
		return
	}

	count++

	var messaggio string
	if tentativo < numeroGenerato {
		messaggio = "Troppo basso! Riprova!"
	} else if tentativo > numeroGenerato {
		messaggio = "Troppo alto! Riprova!"
	} else {
		messaggio = fmt.Sprintf("Hai indovinato in %d tentativi!", count)
		numeroGenerato = rand.Intn(100) // Nuovo numero generato per una nuova partita
		count = 0                       // Resetta il contatore dei tentativi
	}

	// Risposta al cliente con il risultato
	tmpl, _ := template.New("result").Parse(`
	<!DOCTYPE html>
	<html lang="it">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Gioco del Numero</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f4f4f9;
				color: #333;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
			}
			.container {
				text-align: center;
				background-color: #fff;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				width: 100%;
				max-width: 400px;
			}
			h1 {
				font-size: 2em;
				color: #2c3e50;
			}
			p {
				font-size: 1.2em;
				margin-top: 20px;
			}
			a {
				display: inline-block;
				margin-top: 20px;
				color: #3498db;
				text-decoration: none;
				font-size: 1.1em;
			}
			a:hover {
				text-decoration: underline;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Gioco del Numero</h1>
			<p>{{.}}</p>
			<a href="/">Prova un altro numero!</a>
		</div>
	</body>
	</html>
	`)
	tmpl.Execute(w, messaggio)
}

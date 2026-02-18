package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	MovieID int
	Title   string
	Genres  string
}

// Global variable to store movies in memory
var movies []Movie

func main() {
	// 1. Load Data
	loadData("./data/movies.csv")

	// 2. Define Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/search", searchHandler)

	// 3. Start Server
	fmt.Println("🚀 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func loadData(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, _ := reader.ReadAll()

	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		movies = append(movies, Movie{MovieID: id, Title: row[1], Genres: row[2]})
	}
}

// homeHandler renders the initial page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

// searchHandler processes the search and returns results
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var results []Movie

	if query != "" {
		for _, m := range movies {
			if strings.Contains(strings.ToLower(m.Title), query) ||
				strings.Contains(strings.ToLower(m.Genres), query) {
				results = append(results, m)
			}
		}
	}
	renderTemplate(w, "index", results)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// We define the HTML directly in the code for simplicity in one file
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go Movie Toolkit</title>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
		<style>
			body { background-color: #f8f9fa; padding-top: 50px; }
			.search-container { max-width: 800px; margin: auto; background: white; padding: 30px; border-radius: 10px; shadow: 0 4px 6px rgba(0,0,0,0.1); }
		</style>
	</head>
	<body>
		<div class="container search-container">
			<h2 class="text-center mb-4">🎬 Go Movie Toolkit</h2>
			<form action="/search" method="get" class="d-flex mb-4">
				<input type="text" name="q" class="form-control me-2" placeholder="Search Title or Genre...">
				<button type="submit" class="btn btn-primary">Search</button>
			</form>
			
			{{if .}}
				<table class="table table-striped">
					<thead><tr><th>ID</th><th>Title</th><th>Genres</th></tr></thead>
					<tbody>
						{{range .}}
						<tr><td>{{.MovieID}}</td><td>{{.Title}}</td><td>{{.Genres}}</td></tr>
						{{end}}
					</tbody>
				</table>
			{{else}}
				<p class="text-muted text-center">Enter a keyword to see results.</p>
			{{end}}
		</div>
	</body>
	</html>`

	t, _ := template.New(tmpl).Parse(html)
	t.Execute(w, data)
}

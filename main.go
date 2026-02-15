package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	MovieID int
	Title   string
	Genres  string
}

func main() {
	// 1. Open the file
	file, err := os.Open("./data/movies.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// 2. Read the CSV data
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	var movies []Movie

	// 3. Map raw text into our Movie list (The "Loop")
	for i, row := range records {
		if i == 0 {
			continue // Skip headers
		}

		id, _ := strconv.Atoi(row[0])

		currentMovie := Movie{
			MovieID: id,
			Title:   row[1],
			Genres:  row[2],
		}

		movies = append(movies, currentMovie)
	}

	fmt.Printf("✅ Success! Loaded %d movies.\n", len(movies))
	fmt.Println("--------------------------------------------------")

	// 4. Advanced Interactive Search (Title OR Genre)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n[Search Mode]")
		fmt.Print("Enter a Title or Genre (e.g., 'Action', 'Matrix') or 'exit': ")

		if !scanner.Scan() {
			break
		}
		input := strings.ToLower(scanner.Text())

		if input == "exit" {
			fmt.Println("Closing program. Goodbye!")
			break
		}

		matchCount := 0
		fmt.Printf("Searching for: %s...\n", input)
		fmt.Println("--------------------------------------------------")

		for _, m := range movies {
			// Check if input is in Title OR in Genres
			titleMatch := strings.Contains(strings.ToLower(m.Title), input)
			genreMatch := strings.Contains(strings.ToLower(m.Genres), input)

			if titleMatch || genreMatch {
				fmt.Printf("[%d] %s | %s\n", m.MovieID, m.Title, m.Genres)
				matchCount++
			}
		}

		if matchCount > 0 {
			fmt.Printf("\n✨ Found %d matches.\n", matchCount)
		} else {
			fmt.Println("❌ No matches found for that keyword.")
		}
	}
}

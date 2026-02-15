OVERVIEW
For my capstone project, I decided to use into Go (Golang). I wanted to build a toolkit that demonstrates how to handle large-scale data processing without the overhead of heavy frameworks. Using a dataset of over 9,000 movies, I built this tool to show how Go's standard library can efficiently read, parse, and search through data.

I chose Go for this project because:

It's incredibly fast: I wanted to see how it handled thousands of rows compared to interpreted languages.

Type Safety: The static typing helped me catch bugs during development rather than at runtime.

Simplicity: The syntax is clean, making it a perfect "bridge" language for someone coming from Python.

System Requirements
To run my project, you’ll need:

Operating System: macOS, Windows, or Linux.

Go Compiler: Version 1.18 or higher.

Text Editor: I used VS Code, but any editor works.

Dataset: A movies.csv file 

How to Set This Up
Install Go: I downloaded mine from golang.org.

Clone my Repository:

Bash
git clone https://github.com/mandykwamboka/MORINGA-PROJECT.git
cd MORINGA-PROJECT

Drop your movies.csv file.

Run the Code:

Bash
go run main.go

   How the Code Functions


File Initialization: The program opens the movies.csv file using Go's os package.

Data Parsing: We use the encoding/csv library. By setting reader.LazyQuotes = true, the code is able to handle movie titles that contain commas or unconventional quotation marks without crashing.

Data Structuring: Each row from the CSV is converted into a Movie struct. This allows us to access data using logical names like Movie.Title and Movie.Genres instead of just array indices.

Memory Storage: All movie objects are stored in a Slice (a dynamic array in Go), which allows for lightning-fast searching through the 9,000+ records.

⌨️ How to Navigate the Output

Once you run the program using go run main.go, follow these steps to interact with the data:

Initialization: The terminal will display a success message:  Loaded 9742 movies into memory.

Searching by Title: Type a movie name (e.g., Toy Story). The program will perform a case-insensitive search and return all matches with their IDs.

Searching by Genre: You can also type a specific genre (e.g., Adventure or Comedy). The tool will filter the list and show you every movie categorized under that genre.

Combined Search: If you type a keyword like 1995, it will show you all movies with that year in the title.

Exiting: To close the program, simply type exit at the prompt.

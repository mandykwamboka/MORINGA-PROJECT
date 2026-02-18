

Overview

For my capstone project, I built a high-performance movie search application using Go (Golang). While I initially started with a terminal-based tool, I iterated on the design to create a Web-based Graphical User Interface (GUI). This application demonstrates how Go can handle large-scale data processing (over 9,000 records) and serve it instantly to a browser using only the standard library.

I chose Go for this project because:

Speed: It processes thousands of rows and serves web requests with nearly zero latency.

The Standard Library: I built the entire web server using the net/http package without needing heavy external frameworks.

Type Safety: Go's strict typing ensured that the movie data was validated before ever reaching the user's screen.

 System Requirements
 
To run this project, you’ll need:

Go Compiler: Version 1.18 or higher.

Web Browser: Chrome, Firefox, Safari, or Edge.

Dataset: The movies.csv file .

How to Set This Up
Clone the Repository:

Bash
git clone https://github.com/mandykwamboka/MORINGA-PROJECT.git
cd MORINGA-PROJECT


Launch the Web Server:

Bash
go run main.go
Open the Interface: Once the terminal says "Server started," open your browser and go to:
http://localhost:8080

 How the Code Functions (The Backend)
The application architecture is split into a professional data pipeline:

Data Ingestion: The program uses the os and encoding/csv packages to read the dataset. I implemented LazyQuotes = true to prevent crashes caused by unconventional punctuation in movie titles.

Memory Management: Each movie is stored in a Slice of Structs, allowing the search engine to filter thousands of records in milliseconds.

Web Server Logic: I used Go's net/http package to create a local server. The html/template package is used to dynamically inject search results into a clean, modern HTML frontend.

How to Navigate the Web GUI
Instead of a text-only prompt, you now have a visual dashboard:

The Dashboard: When you load the page, you'll see a clean search bar.

Instant Search: Enter a Title (e.g., Matrix) or a Genre (e.g., Sci-Fi) and click "Search."

Dynamic Results: The app will instantly generate a table showing the Movie ID, Title, and Genres.

Reset: To start over, simply clear the search bar or refresh the page.


Initial Research

Prompt: "How do I read a CSV file in Go and store it in a struct?"

AI Response Summary: Provided the encoding/csv package basics and showed how to define a struct with tags.

Evaluation: This was helpful for scaffolding the basic code, but I didn't know how to handle files that weren't perfectly formatted yet.

2. Debugging Data Issues

Prompt: "I'm getting 'parse error on line 1: bare " in non-quoted-field' while running my Go program. What does this mean?"

AI Response Summary: Explained that Go's CSV parser is strict about quotes. Suggested using reader.LazyQuotes = true.

Evaluation:  It saved me hours of manually editing a 9,000-line file. It taught me that real-world data is often "messy."

3. Identifying File Format Errors

Prompt: "My CSV looks binary meta data when I print the first few lines in Go. It shows things like 'Index/Document.iwa'."

AI Response Summary: Identified that the file was an Apple Numbers binary file renamed to .csv and instructed me to export it as a true UTF-8 CSV.

Evaluation: This taught me about file encodings and binary vs. text formats, which is a fundamental concept in software development.

4. Adding Features

Prompt: "How can I make a terminal search tool that looks through both Title and Genre fields and counts the matches?"

AI Response Summary: Suggested using a for loop with strings.Contains and a counter variable.

Evaluation: This allowed me to add "Creativity" to the project, turning a static list into an interactive tool.

Reflection: How AI Improved Productivity

Using AI as a "pair programmer" allowed me to jump into a complex language like Go and have a working prototype in a few hours.

Clarity: Instead of just getting code, the AI explained why the LazyQuotes setting existed.

Problem Solving: The AI acted as a senior engineer when I hit the binary file error, diagnosing a problem I didn't even know existed (file encoding).

Feedback: Refining my prompts from "It doesn't work" to "I have this specific error code" made the solutions much more accurate.

COMMON ISSUES AND FIXES 
 My Debugging Journey

The development process wasn't exactly a straight line. I hit a few major roadblocks that forced me to rethink how Go handles data. Here are the main issues I ran into and how I fought through them:

The "Bare Quote" Headache Early on, my program kept crashing with a parse error on line 1: bare " in non-quoted-field. After some digging, I realized that movie titles in the CSV often have complex punctuation—like American President, The (1995)—that confuses the standard Go parser.

The Fix: I had to explicitly enable reader.LazyQuotes = true. This taught me that Go is very strict by default, and you have to tell it when to be "forgiving" with messy, real-world data.

The "Gibberish" File Encoding Issue At one point, my terminal was just printing weird symbols and binary gibberish. I eventually realized the movies.csv file I was using wasn't actually a plain text file; it was an Apple Numbers file that had been renamed.

The Fix: I had to perform a clean export from the spreadsheet software to CSV (UTF-8). It was a huge lesson in verifying your data source before writing code—if the input is "garbage," the output will be too.

The "Vanishing" Terminal Initially, the program would run, find the movies, and then immediately close before I could read anything. I needed the program to stay "alive" so I could actually interact with it.

The Fix: I implemented a bufio scanner inside an infinite for loop. This keeps the program waiting for user input, allowing for a continuous search experience until I type "exit."

Missing Package Imports I kept getting "undefined" errors for basic functions like strings.Contains or strconv.Atoi. Coming from other languages, I wasn't used to Go's strictness regarding imports.

The Fix: I learned to manually manage the import block at the top of the file. Every time I added a new feature (like searching or converting strings to numbers), I had to ensure the corresponding library was declared.
References

Official Go Documentation: pkg.go.dev/encoding/csv – The source of truth for the CSV library.

Go By Example: gobyexample.com – A great hands-on resource for learning Go syntax (especially Structs and Slices).

MovieLens Dataset: GroupLens Research – The source of the movies.csv data.

Moringa School Canvas: Project requirements and Python-to-Go transition guides

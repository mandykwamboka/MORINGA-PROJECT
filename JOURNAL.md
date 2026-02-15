Initial Research

Prompt: "How do I read a CSV file in Go and store it in a struct?"

AI Response Summary: Provided the encoding/csv package basics and showed how to define a struct with tags.

Evaluation: This was helpful for scaffolding the basic code, but I didn't know how to handle files that weren't perfectly formatted yet.

2. Debugging Data Issues

Prompt: "I'm getting 'parse error on line 1: bare " in non-quoted-field' while running my Go program. What does this mean?"

AI Response Summary: Explained that Go's CSV parser is strict about quotes. Suggested using reader.LazyQuotes = true.

Evaluation: This was the "Aha!" moment. It saved me hours of manually editing a 9,000-line file. It taught me that real-world data is often "messy."

3. Identifying File Format Errors

Prompt: "My CSV looks like gibberish when I print the first few lines in Go. It shows things like 'Index/Document.iwa'."

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

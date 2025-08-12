# Challenge 1: User Input

## Problem

Write a Go program that:

1. Prints a welcome message
2. Asks the user for their name
3. Reads input from the console
4. Prints a greeting message with their name

## Approach

- Used `bufio.NewReader` to read input from `os.Stdin`
- Used `ReadString('\n')` to capture user input
- Trimmed newline characters using `strings.TrimSpace`

## Example Output

Welcome! Please enter your name:
Tirth
Hello, Tirth !

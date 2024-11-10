# CC-Grep - Another Command-Line Grep Utility

CC-Grep is a custom command-line grep utility developed as part of a coding challenge. It supports case-insensitive searches, recursive file matching, and inverted matches to display lines that do not contain the specified pattern.

Challenge URL: [https://codingchallenges.fyi/challenges/challenge-grep](https://codingchallenges.fyi/challenges/challenge-grep)

## Features

- **Case-Insensitive Search**: Use the `-i` flag to perform case-insensitive pattern matching.
- **Recursive Search**: Use the `-r` flag to search through all files within a directory and its subdirectories.
- **Inverted Match**: Use the `-v` flag to return lines that do **not** contain the specified pattern.
- **Command-line Interface (CLI)**: Simple and efficient CLI for flexible text searching.

## Getting Started

These instructions will help you set up and use `cc-grep` on your local machine.

### Prerequisites

- **Go**: Ensure that Go (version 1.15 or later) is installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-grep.git
cd cc-grep
```

## Building

To compile the project, run:

go build -o ccgrep

## Testing

Run tests to verify the application’s functionality:

```bash
go test ./...
```

## Usage

To use the utility, execute the compiled binary with the desired options:

### Basic Usage

Search for a pattern in a file:

```bash
./ccgrep "Rock" test_data/rockbands.txt
```

### Flags

- **Case-Insensitive Search**: Use -i or --case-insensitive to ignore case.
- **Recursive Search**: Use -r or --recursive to search through all files in a directory and its subdirectories.
- **Inverted Match**: Use -v or --invert to return lines that do not contain the specified pattern.

## Examples

### Case-Insensitive Search

Perform a case-insensitive search for “symbol” in symbols.txt:

```bash
./ccgrep -i "symbol" test_data/symbols.txt
```

### Recursive Search

Search recursively for the word “1985” in all files within the test_data directory and its subdirectories:

```bash
./ccgrep -r "1985" test_data
```

### Inverted Match

Find all lines in test.txt that do not contain the word “example”:

```bash
./ccgrep -v "example" test_data/test.txt
```

### Combine Flags

Perform a recursive, case-insensitive search for “band” in all files within test_data, returning only lines that do not contain the pattern:

```bash
./ccgrep -r -i -v "band" test_data
```

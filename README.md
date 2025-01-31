# Monkey Interpreter

This project is an implementation of an interpreter for the Monkey programming language. The goal is to learn about the inner workings of interpreters and gain hands-on experience by building one from scratch using Go.

## Getting Started

To run the interpreter, ensure that you have Go installed on your system. Clone this repository and navigate to the project directory. Then, run the following command:

```bash
go run main.go
```

This will start the REPL, and you can begin entering Monkey code. The interpreter will parse and evaluate the code, providing the results in the REPL.

## Why Go?

- Closely maps to C (or other lower level languages)
- Readable and easy to understand
- Built-in tooling
- Fast

## Project Structure

### `main.go`
The entry point of the interpreter.

### `repl/repl.go`
Implementation of the REPL.

### `lexer/lexer.go`
Defines the `Lexer` struct and its methods for tokenizing input.

### `lexer/lexer_test.go`
Unit tests for the lexer.

### `token/token.go`
Defines the `Token` struct and token-related constants.

### `ast/ast.go`
Defines the Abstract Syntax Tree (AST) nodes and their associated interfaces.

### `parser/parser.go`
Implementation of the parser for constructing an AST from tokens.

## TODO
- [ ] Implement Parser

## Acknowledgements
This project follows the concepts taught in the book “Writing an Interpreter in Go” by Thorsten Ball. Thorsten's book serves as an excellent resource for learning about interpreters, and the subsequent project has been incredibly enlightening while also helping me become more familiar with Golang!
# Claude Code Instructions

## Code Assistance
- Do not provide code snippets unless explicitly asked for them
- Only provide ideas on modules, functions, and algorithmic approaches

## Testing
- Always run test suites after every refactor
- For Gleam projects, use `gleam test` to run tests
- Ensure all tests pass before considering a refactor complete

## Project Structure
This is an Exercism Gleam track workspace containing multiple exercise directories.

## Gleam Style Guidelines
- Avoid piping into anonymous functions
- Instead, bind results with `let` and pattern match, or start a new pipe chain
- This is the idiomatic Gleam style preferred by the language creator
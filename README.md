# Tic-Tac-Toe — Low Level System Design (Go)

A hands-on project to learn Go and LLD principles by building a tic-tac-toe game from scratch.

## Project Structure

```
tic-tac-toe/
├── cmd/                    # Application entry point (main.go)
├── internal/               # Private application code
│   ├── board/              # Board representation & logic
│   ├── game/               # Game loop, rules, state management
│   └── player/             # Player abstraction (human, AI later)
├── go.mod
└── README.md
```

## Learning Goals
- Go project layout & packages
- Structs, interfaces, methods
- Error handling idioms
- Clean separation of concerns (LLD)

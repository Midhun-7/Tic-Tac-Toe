# Tic-Tac-Toe — Low Level System Design (Go)

A hands-on project to learn Go and LLD principles by building a tic-tac-toe game from scratch.

## Project Structure

```
tic-tac-toe-go/
├── cmd/                    # Application entry point (main.go)
├── internal/               # Private application code
│   ├── board/              # Board representation & logic
│   ├── game/               # Game loop, rules, state management
│   └── player/             # Player abstraction (human, AI later)
├── pkg/
│   └── utils/              # Shared utilities (input helpers, etc.)
├── go.mod                  # (you'll create this with `go mod init`)
├── progress.md             # Track your learning journey
├── syntax_to_remember.md   # Go syntax quick-reference
└── README.md               # This file
```

## Learning Goals
- Go project layout & packages
- Structs, interfaces, methods
- Error handling idioms
- Clean separation of concerns (LLD)

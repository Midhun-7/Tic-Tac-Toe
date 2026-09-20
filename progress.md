# Progress Tracker 🚀

> Tracks milestones, what was learned, and current status.

---

## Current Phase: **Phase 3 — Polish & Edge Cases**

| # | Milestone | Status | Date |
|---|-----------|--------|------|
| 1 | Project folder structure created | ✅ | 2026-09-19 |
| 2 | Understand the design & plan components | ✅ | 2026-09-19 |
| 3 | Define the Board (data structure + methods) | ✅ | 2026-09-19 |
| 4 | Define the Player (struct + constructor) | ✅ | 2026-09-19 |
| 5 | Define the Game (struct) | ✅ | 2026-09-19 |
| 6 | Display the board | ✅ | 2026-09-20 |
| 7 | Win detection logic | ✅ | 2026-09-20 |
| 8 | Implement Game loop | ✅ | 2026-09-20 |
| 9 | **GAME IS PLAYABLE!** 🎉 | ✅ | 2026-09-20 |
| 10 | Draw detection | ⬜ | — |
| 11 | Input validation (out of bounds, invalid input) | ⬜ | — |
| 12 | Refactor & clean up | ⬜ | — |
| 13 | (Stretch) AI opponent | ⬜ | — |
| 14 | (Stretch) Web frontend with Go HTTP server | ⬜ | — |

---

## Design Decisions

| Question | Answer | Reasoning |
|----------|--------|-----------|
| Core entities | Player, Board, Game | Game is the orchestrator/referee |
| Board data structure | `[3][3]string` | Fixed-size, dense grid — array over map/slice |
| Cell values | `""`, `"X"`, `"O"` | 3 states need string, not bool |
| Player type | struct with constructor | Clean, only holds Name & Mark |
| Turn tracking | `CurrentPlayer int` in Game | Index into Players array; Game owns turn logic |
| Two players | `[2]player.Player` array | Fixed count, array over slice |
| Win check ownership | Board method | Board knows its own grid (encapsulation) |

---

## Lessons Learned

1. **Value vs Pointer receiver** — use `*` when modifying struct data
2. **Unused imports** — Go won't compile with them
3. **Case sensitivity** — `string` not `String`, `bool` not `boolean`
4. **Early return** — no `else` needed after `return`
5. **Separation of concerns** — Turn belongs in Game, not Player
6. **Can't define methods on types from other packages**
7. **Constructor pattern** — `NewTypeName()` returning a pointer
8. **Code must be inside functions** — no floating statements
9. **0-indexed arrays** — Go arrays start at 0, not 1
10. **Empty string gotcha** — `"" == ""` is true, always check `!= ""` before comparing
11. **Program entry point** — must have `package main` + `func main()`
12. **Import path** — starts with module name: `"tic-tac-toe/internal/board"`

---

## Session Log

### Session 1 — 2026-09-19
- Created project folder structure
- Designed and implemented all 3 core entities: Board, Player, Game
- Linked project to GitHub

### Session 2 — 2026-09-20
- Implemented Display with grid formatting
- Implemented CheckWin (rows, columns, both diagonals)
- Fixed empty cell matching bug
- Implemented game loop with turn switching
- Winner announcement with correct player name
- **First successful game run! 🎉**
- Next: Draw detection, input validation

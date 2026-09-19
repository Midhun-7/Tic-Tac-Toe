# Progress Tracker 🚀

> Tracks milestones, what was learned, and current status.

---

## Current Phase: **Phase 2 — Methods & Game Logic**

| # | Milestone | Status | Date |
|---|-----------|--------|------|
| 1 | Project folder structure created | ✅ | 2026-09-19 |
| 2 | Understand the design & plan components | ✅ | 2026-09-19 |
| 3 | Define the Board (data structure + methods) | ✅ | 2026-09-19 |
| 4 | Define the Player (struct + constructor) | ✅ | 2026-09-19 |
| 5 | Define the Game (struct) | ✅ | 2026-09-19 |
| 6 | Display the board | ⬜ | — |
| 7 | Implement Game loop | ⬜ | — |
| 8 | Win / Draw detection logic | ⬜ | — |
| 9 | Input validation & error handling | ⬜ | — |
| 10 | Refactor & clean up | ⬜ | — |
| 11 | (Stretch) AI opponent | ⬜ | — |

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

---

## Session Log

### Session 1 — 2026-09-19
- Created project folder structure
- Designed and implemented all 3 core entities: Board, Player, Game
- Board: `[3][3]string` with `PlaceMark` method (pointer receiver)
- Player: `Name` + `Mark` with `NewPlayer` constructor
- Game: holds Board, 2 Players, CurrentPlayer index
- Next: Display the board, then build the game loop

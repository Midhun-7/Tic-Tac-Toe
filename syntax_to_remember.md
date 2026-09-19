# Go Syntax to Remember 📝

> This file captures Go syntax patterns encountered during the tic-tac-toe project.
> Updated as we go — use this as a quick reference.

---

## 1. Struct Definition

```go
type StructName struct {
    FieldName  FieldType
    AnotherField AnotherType
}
```

---

## 2. Package Declaration

Every `.go` file must start with a package declaration:

```go
package packagename
```

The package name should match the folder name (e.g., file in `internal/board/` → `package board`).

---

## 3. Method Syntax

```go
func (receiverName ReceiverType) MethodName(params) returnType {
    // body
}
```

### Value receiver vs Pointer receiver

```go
func (b Board) Display()                    // gets a COPY — cannot modify the original
func (b *Board) PlaceMark(row, col int)     // gets a POINTER — can modify the original
```

**Rule of thumb**: Use `*` (pointer receiver) when the method needs to **change** the struct's data.

---

## 4. Early Return Pattern

Instead of `if-else`, return early on error/invalid cases:

```go
// ✅ Idiomatic Go
if condition {
    return
}
// normal code continues here
```

---

## 5. If-Else Brace Rule

In Go, `else` must be on the **same line** as the closing `}`:

```go
// ✅ Correct
if x > 0 {
    // ...
} else {
    // ...
}
```

---

## 6. Go is Case-Sensitive for Types

Built-in types are **always lowercase**:

```go
string    // ✅
String    // ❌ not a built-in type

bool      // ✅
boolean   // ❌ not a Go type

int       // ✅
Int       // ❌
```

---

## 7. Unused Imports

Go **won't compile** if you import a package and don't use it:

```go
import "fmt"   // ❌ error if fmt is never used in the file
```

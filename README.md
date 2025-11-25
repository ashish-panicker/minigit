# MiniGit – A Minimal Git-like Source Control System (Learning Project)

MiniGit is a simplified version-control system inspired by Git, built in Go for educational purposes.
The goal is to learn Go deeply by implementing core concepts of a DVCS (Distributed Version Control System), step-by-step, without relying on external Git internals.

This project does **not** aim to fully replicate Git.
Instead, it focuses on *core plumbing* commands and internal data structures.

# Project Goals

1. Learn Go by building a meaningful systems-level project
2. Understand how Git works under the hood
3. Implement essential concepts:

   * Object storage (blobs, trees, commits)
   * Staging area (index)
   * Object hashing
   * Persisting metadata
4. Build a CLI tool using idiomatic Go
5. Practice Go modules, packages, error handling, structs, interfaces, testing, and file I/O

# Scope (What you **will** build)

## 1. Repository Initialization

Command:

```
mygit init
```

Creates:

```
.mygit/
    objects/
    refs/
    HEAD
```

Concepts:

* Directory creation
* Minimal repository metadata
* Go filesystem operations

## 2. Blob Objects (File Snapshots)

Command:

```
mygit hash-object <file>
mygit cat-file <hash>
```

You will:

* Read file contents
* Compute SHA-1 hash
  *(You will implement hashing using Go's crypto libraries — not write SHA-1 yourself.)*
* Store compressed blob objects
* Retrieve objects by hash

## 3. Index (Staging Area)

Command:

```
mygit add <file>
```

You will:

* Build an index file containing tracked files
* Track file state (path, hash, size, timestamp)
* Understand how Git's staging area acts as an intermediate layer

## 4. Tree Objects (Directory Snapshots)

Command:

```
mygit write-tree
mygit ls-tree <hash>
```

You will:

* Collect blobs from index
* Build directory hierarchy
* Create tree objects referencing blob hashes
* Store trees under `objects/`

## 5. Commit Objects

Command:

```
mygit commit -m "message"
```

You will:

* Create commit objects
* Point to tree objects and parent commits
* Update `refs/heads/main`
* Update `HEAD`

---

## 6. Checkout (Very Minimal)

Command:

```
mygit checkout <commit>
```

You will:

* Read tree object
* Restore working directory files
  (*No branch switching or complex merges — just basic restore*)

# Scope (What you **will NOT** build)

This protects the project from getting too large.

- Full branching system
- Merge algorithms
- Diffs
- Network operations (push/pull)
- Packfiles
- Rebase / cherry-pick
- Advanced index features (conflicts, flags)

# Project Structure (Recommended)

```
mygit/
│
├── cmd/
│   └── mygit/
│        └── main.go
│
├── internal/
│   ├── repo/
│   ├── objects/
│   ├── index/
│   ├── plumbing/
│   └── commands/
│
├── go.mod
└── README.md
```

# Learning Outcomes

By completing MiniGit, you will improve your Go skills in:

* Structs, interfaces, and package organization
* File system operations
* Hashing and binary encoding
* CLI application structure
* Error handling patterns
* Testing with Go’s built-in test framework
* Using Go Modules
* Designing maintainable, modular code

# Development Roadmap

### **Phase 1 — Basics**

* Set up project structure
* Implement `init` command
* Build utility functions (file operations, hashing)

### **Phase 2 — Blob Layer**

* Implement object storage
* Implement blob read/write
* Test with `hash-object` + `cat-file`

### **Phase 3 — Staging Area**

* Build index file
* Implement `add` command

### **Phase 4 — Tree Layer**

* Convert index into tree objects
* Implement `write-tree` and `ls-tree`

### **Phase 5 — Commit Layer**

* Implement commit creation
* Manage refs and HEAD

### **Phase 6 — Checkout**

* Restore working files from commit tree

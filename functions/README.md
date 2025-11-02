# Lem-in   
<a id="readme-top"></a>

[](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNTguNzUiIGhlaWdodD0iMzUiIHZpZXdCb3g9IjAgMCAxNTguNzUgMzUiPjxyZWN0IHdpZHRoPSI4NC4yMTY2NjcxNzUyOTI5NyIgaGVpZ2h0PSIzNSIgZmlsbD0iIzRhOTBlMiIvPjxyZWN0IHg9Ijg0LjIxNjY2NzE3NTI5Mjk3IiB3aWR0aD0iNzQuNTMzMzMyODI0NzA3MDMiIGhlaWdodD0iMzUiIGZpbGw9IiNjY2Q4ZTciLz48dGV4dCB4PSI0Mi4xMDgzMzM1ODc2NDY0ODQiIHk9IjIxLjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInUm9ib3RvJywgc2Fucy1zZXJpZiIgZmlsbD0iI0ZGRkZGRiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgbGV0dGVyLXNwYWNpbmc9IjIiPkdPTEFORzwvdGV4dD48dGV4dCB4PSIxMjEuNDgzMzMzNTg3NjQ2NDgiIHk9IjIxLjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInTW9udHNlcnJhdCcsIHNhbnMtc2VyaWYiIGZpbGw9IiMwZTBlMGUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSI5MDAiIGxldHRlci1zcGFjaW5nPSIyIj4xLjIyLjM8L3RleHQ+PC9zdmc+)



<a href="https://gitpoint.co/">
    <img alt="GitPoint" title="GitPoint" src="https://cdn.creazilla.com/cliparts/24564/1524593971-clipart-xl.png" width="350">
  </a>

## Description

This project is designed to enable us to create a digital version of an ant farm. We need to help the ants find their way through a colony of connected rooms, from the starting point to the exit. The goal is to find the shortest path and move all the ants in the minimum number of turns.

**Repository**: [lem-in](https://learn.zone01oujda.ma/git/obenmbar/lem-in)






## Features

- Efficient pathfinding algorithm using BFS (Breadth-First Search)
- Multiple path finding and optimization
- Handles various graph configurations and edge cases
- Validates input format and detects invalid farms
- Optimizes ant distribution across multiple paths

## Project Structure

```
lem-in/
├── Check_helpers.go   # Helper functions for input validation
├── Check_utils.go     # Utility functions for checking
├── check.go          # Main input validation logic
├── moove.go          # Ant movement logic
├── Solve_helpers.go  # Helper functions for solving
├── Solve.go          # Main pathfinding algorithm
├── go.mod            # Go module file
└── main.go           # Entry point
```

## Algorithm

The project implements:
1. **BFS Algorithm**: Finds all possible paths from start to end
2. **Path Optimization**: Selects the best combination of paths to minimize total turns
3. **Ant Distribution**: Distributes ants across paths for optimal movement

## Installation

```bash
git clone https://learn.zone01oujda.ma/git/obenmbar/lem-in.git
cd lem-in
go build
```

## Usage

```bash
go run main.go test_file.txt
```

### Input Format

```
number_of_ants
##start
start_room x y
room1 x y
room2 x y
##end
end_room x y
start_room-room1
room1-room2
room2-end_room
```

### Output Format

```
[Input content]

L1-room1 L2-start_room
L1-room2 L2-room1 L3-start_room
L1-end_room L2-room2 L3-room1
...
```

Where `Lx-roomy` represents ant number `x` moving to room `y`.

### Example : 
```bash
$ go run . level-2.txt
4
##start
start 0 0
1 1 0
2 2 0
##end
end 3 0
start-1
start-2
1-end
2-end

L1-1 L2-2
L1-end L2-end L3-1 L4-2
L3-end L4-end
```

## Rules

- Ants start at `##start` room and must reach `##end` room
- Only one ant per room at a time (except start and end)
- Each turn, each ant can move to an adjacent room
- Goal: Move all ants in minimum number of turns
- Output must display the original input followed by ant movements

## Instructions

- We must use Go's standard library only
- The project must compile without errors
- Imports must be formatted correctly
- Functions, variables, constants, types must follow good practices
- Code must be well-structured and properly tested
- It is recommended to have test files with `_test` suffix for unit testing

## Allowed Packages

Only the following Go standard library packages are allowed:
- `os`
- `fmt`
- `strconv`
- `strings`
- Other standard library packages as needed

## Learning Objectives

This project helps us learn about:
- Graph theory and pathfinding algorithms
- Algorithm optimization
- Input parsing and validation
- Efficient data structures
- Problem-solving and algorithmic thinking

---



## Authors

- **Othmane Benmbarek** - [obenmbar](https://learn.zone01oujda.ma/git/obenmbar)
- **Mohamed Nouri** - [mohnouri](https://learn.zone01oujda.ma/git/mohnouri)
- **BEMAMORY Nomenjanahary Luciano Loic** - [bnomenja](https://learn.zone01oujda.ma/git/bnomenja)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

---
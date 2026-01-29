# Smart Layout Engine

A high-performance 2D layout optimization engine built with Go and WebAssembly. The system automatically arranges furniture/items within a room space while avoiding collisions using backtracking algorithms.

## Overview

Smart Layout Engine is a full-stack web application that solves the 2D bin packing problem with spatial constraints. It combines:

- **Go Backend**: Core layout solver with collision detection and backtracking algorithms
- **WebAssembly Frontend**: High-performance in-browser computation via Go compiled to WASM
- **Gin Web Server**: REST API for layout persistence and frontend serving
- **HTML5 Canvas**: Interactive visualization of layout results

## Features

- **Automatic Layout Generation**: Intelligently positions items to maximize space utilization
- **Collision Detection**: AABB (Axis-Aligned Bounding Box) algorithm prevents overlapping
- **Backtracking Solver**: Systematic search algorithm to find valid arrangements
- **Fixed Position Support**: Allow certain items to remain stationary while others are arranged
- **WebAssembly Performance**: Run computationally intensive layout algorithms in the browser
- **Interactive UI**: Add items dynamically and visualize results in real-time
- **RESTful API**: Save and retrieve layout configurations

## Architecture

```
smart-layout-engine/
├── core/           # Domain logic and algorithms
│   ├── domain.go       # Core types (Entity, Point, Dimension)
│   ├── solver.go       # Backtracking layout solver
│   └── validator.go    # Collision detection (AABB)
├── server/         # Gin web server
│   └── main.go         # HTTP endpoints and static file serving
├── wasm/           # WebAssembly entry point
│   └── main.go         # JavaScript bridge for Go Wasm functions
├── frontend/       # Web UI
│   ├── index.html      # Main application page
│   ├── app.js          # Client-side logic
│   └── wasm_exec.js    # Go WebAssembly runtime
└── test files
```

## Getting Started

### Prerequisites

- Go 1.25.4 or later
- Modern web browser with WebAssembly support

### Installation

1. Clone the repository:
```bash
git clone https://github.com/nguyentuio1/smart-layout-engine.git
cd smart-layout-engine
```

2. Install dependencies:
```bash
go mod download
```

3. Build the WebAssembly module:
```bash
GOOS=js GOARCH=wasm go build -o frontend/main.wasm ./wasm
```

### Running the Application

1. Start the server:
```bash
go run server/main.go
```

2. Open your browser:
```
http://localhost:8080/app
```

### Running Tests

```bash
go test ./core -v
```

## Usage

### Via Web Interface

1. Enter item name, width, and height
2. Click "Add" to add items to the list
3. Click "Run Algorithm" to generate the layout
4. View the arranged items on the canvas

### Via API

**Health Check**
```http
GET /ping
```

**Save Layout** (placeholder for future database integration)
```http
POST /save-layout
Content-Type: application/json

{
  "items": [...]
}
```

### Programmatic Usage (Go)

```go
import "github.com/nguyentuio1/smart-layout-engine/core"

solver := core.SimpleSolver{
    Step: 10,
    Room: core.Dimension{W: 500, H: 500},
}

items := []core.Entity{
    {ID: "Table", Size: core.Dimension{W: 100, H: 60}},
    {ID: "Chair", Size: core.Dimension{W: 40, H: 40}},
}

result, err := solver.Solve(items)
if err != nil {
    // Handle error
}
```

## Algorithm Details

### Backtracking Solver

The solver uses a recursive backtracking approach:

1. Sort items by area (largest first for better pruning)
2. For each item, try every valid position in the room
3. Check collision with already-placed items using AABB
4. Recursively place the next item
5. Backtrack if no valid position found

### Collision Detection (AABB)

Two entities overlap if:
```
a.X < b.X + b.W &&
a.X + a.W > b.X &&
a.Y < b.Y + b.H &&
a.Y + a.H > b.Y
```

## Configuration

- **Step Size**: Controls the granularity of position attempts (default: 10px)
- **Room Size**: Defined in `wasm/main.go` (default: 500x500)
- **Server Port**: Configured in `server/main.go` (default: 8080)

## Performance Considerations

- Larger step sizes trade precision for speed
- Complex layouts with many items may require longer computation time
- WebAssembly performance varies by browser (Chrome/Firefox recommended)
- Items are sorted by area to optimize the search space

## Future Enhancements

- [ ] Database integration for layout persistence
- [ ] Support for rotation constraints
- [ ] Advanced constraint system (adjacency, distance, alignment)
- [ ] Multiple solver algorithms (genetic, simulated annealing)
- [ ] Export layouts to SVG/PNG
- [ ] Undo/redo functionality
- [ ] Room boundary constraints (walls, doors, windows)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Author

nguyentuio1

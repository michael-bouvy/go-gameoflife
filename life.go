package main

import "time"
import "fmt"

type cell struct {
    x int
    y int
}

func main() {
    oldGrid := initGrid (20, 10)

    // Initialize grid data
    oldGrid[3][10] = 1
    oldGrid[3][11] = 1
    oldGrid[4][10] = 1
    oldGrid[4][11] = 1

    oldGrid[5][12] = 1
    oldGrid[5][13] = 1
    oldGrid[6][12] = 1
    oldGrid[6][13] = 1

    var duration int64 = 10000000000
    var step int64 = 500000000
    var last int64 = 0
    begin := time.Now().UnixNano()

    for {
        now := time.Now().UnixNano()
        if (now >= (last + step)) {
            last = now
            grid    := initGrid (20, 10)

            for y := range oldGrid {
                fmt.Println(oldGrid[y]);
            }
            fmt.Print("\033[10A"); // Go back 10 lines
            for y := range oldGrid {
                for x := range oldGrid[y] {
                    grid[y][x] = getNewCellState(oldGrid, x, y, oldGrid[y][x])
                }
            }

            copy (oldGrid, grid)

        }

        if (now > (begin + duration)) {
            break
        }
    }
}

func isCellAlive (slice [][]byte, x int, y int) bool {
    return slice[y][x] == 1
}

func getNewCellState (slice [][]byte, x int, y int, cur byte) byte {
    aliveNeighbors := countAliveNeighbors(slice, x, y)
    if (isCellAlive(slice, x, y)) {
        // Alive cell: if 2 or 3 neighbors alive, return 1, else 0
        if (2 == aliveNeighbors || 3 == aliveNeighbors) {
            return 1
        } else {
            return 0
        }
    } else {
        // Dead cell: 3 alive neighbors => return 1, else 0
        if (3 == aliveNeighbors) {
            return 1
        } else {
            return 0
        }
    }
}

func countAliveNeighbors(slice [][]byte, x int, y int) int {
    neighbors := getNeighbors()
    countAlive := 0
    for i := range neighbors {
        for j := range neighbors[i] {
            if (neighbors[i][j] == 0) {
                continue
            }
            xToCheck := x + j - 1
            yToCheck := y + i - 1
            if (xToCheck < 0 || xToCheck >= 20 || yToCheck < 0 || yToCheck >= 10) {
                continue
            }
            if (slice[yToCheck][xToCheck] == 1) {
                countAlive += 1
            }
        }
    }
    return countAlive
}

func initGrid (width int, height int) [][]byte {
    twoD := make([][]byte, height)
    for i := 0; i < height; i++ {
        twoD[i] = make([]byte, width)
        for j := 0; j < width; j++ {
            twoD[i][j] = 0
        }
    }
    return twoD
}

func initGridFromFile() {
    // TODO
}

func getNeighbors() [][]int {
    neighbors := make([][]int, 3)
    neighbors[0] = make([]int, 3)
    neighbors[0][0]   = 1;
    neighbors[0][1]   = 1;
    neighbors[0][2]   = 1;
    neighbors[1] = make([]int, 3)
    neighbors[1][0]   = 1;
    neighbors[1][1]   = 0;
    neighbors[1][2]   = 1;
    neighbors[2] = make([]int, 3)
    neighbors[2][0]   = 1;
    neighbors[2][1]   = 1;
    neighbors[2][2]   = 1;
    return neighbors
}

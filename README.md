# Go ASCII Grid Renderer 🗺️

> A simple, efficient command-line tool written in Go to render character-based maps into detailed ASCII art.

This program reads a grid blueprint from standard input and draws it to the console with borders and specific symbols for different cell types. It's built to be robust, with error handling for common invalid input scenarios.

-----

## Features ✨

  * **Simple Blueprint:** Define your map using a simple grid of numbers.
  * **Detailed Rendering:** Translates the simple blueprint into a bordered grid with distinct multi-character cells.
  * **Multiple Cell Types:** Supports walls, empty space, a player character, and other entities.
  * **Robust Input Validation:** Gracefully handles errors like incorrect dimensions, mismatched row lengths, or invalid characters in the map.
  * **Lightweight & Fast:** Written in Go with no heavy dependencies, ensuring efficient processing.

-----

## How to Use

### Prerequisites

1.  **Go:** You must have the Go programming language installed.
2.  **`ap` Package:** The project depends on the `alem-platform/ap` package. You can install it with the following command:
    ```bash
    go get github.com/alem-platform/ap
    ```

### Running the Program

1.  Save the code as `main.go` in a new directory.
2.  Open your terminal in that directory.
3.  Run the program using:
    ```bash
    go run main.go
    ```
4.  The program will wait for you to provide input directly in the terminal.

### Input Format

The program expects input in the following format:

1.  **First Line:** Two numbers separated by a space: `rows cols`.
      * `rows`: The number of rows in the grid.
      * `cols`: The number of columns in the grid.
2.  **Subsequent Lines:** The map data, one line per row. Each line must contain `cols` number of characters, corresponding to the cell types.

-----

## Example

Here is an example of providing input to the running program and the resulting output.

#### **Input:**

```
3 5
11211
10001
11311
```

#### **Output:**

```
 _______________________________________
|       |       |   >   |       |       |
|       |       |       |       |       |
|_______|_______|_______|_______|_______|
|       |XXXXXXX|XXXXXXX|XXXXXXX|       |
|       |XXXXXXX|XXXXXXX|XXXXXXX|       |
|_______|XXXXXXX|XXXXXXX|XXXXXXX|_______|
|       |       |   @   |       |       |
|       |       |       |       |       |
|_______|_______|_______|_______|_______|

```

-----

## Cell Types

The numbers in your input blueprint correspond to the following rendered objects:

  * `0` - **Wall**: A solid block of `X`'s.
    ```
    |XXXXXXX|
    |XXXXXXX|
    |XXXXXXX|
    ```
  * `1` - **Empty Space**: A clear, open cell.
    ```
    |       |
    |       |
    |_______|
    ```
  * `2` - **Player**: An entity represented by `>`.
    ```
    |   >   |
    |       |
    |_______|
    ```
  * `3` - **Item/Enemy**: An entity represented by `@`.
    ```
    |   @   |
    |       |
    |_______|
    ```

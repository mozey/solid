package main

// Clients should not be forced to depend on interfaces they do not use,
// the interfaces should be designed to be as small and specific as possible.
// This keeps the code flexible and avoids unnecessary coupling between classes.

func main() {
	// TODO Make an example using io.Writer
	// - Specific Purpose: write a slice of bytes to some destination
	// - Minimal Contract: only requires a single method
	// - Client Focus: don't need to depend on any other methods
	// that might be related to other I/O operations e.g.
	// reading, seeking, or closing
	// - Flexibility and Composability: easily combine different io.Writer
	// implementations to create more complex I/O pipelines.
	// For example, you can chain a gzip.Writer to a bufio.Writer
	// to an os.File to compress data before writing it to a file,
	// all while working with the simple io.Writer interface
}

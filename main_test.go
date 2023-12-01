package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainWithNoArgs(t *testing.T) {

	// Save the original command-line arguments and restore them when the test is done
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set the command-line arguments for the test
	os.Args = []string{"programName"} // Assuming program name is always present

	// Use a test function to capture the output and check if it matches the expected output
	output := captureOutput(main)

	expectedOutput := "Usage: mycli <parameter>\n"
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
}

func TestMainWithBytesFlag(t *testing.T) {
	// Save the original command-line arguments and restore them when the test is done
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set the command-line arguments for the test
	os.Args = []string{"programName", "-c", "testfile.txt"} // Replace with a valid file path

	// Use a test function to capture the output and check if it contains the expected output
	output := captureOutput(main)

	expectedOutput := "Bytes:"
	if !contains(output, expectedOutput) {
		t.Errorf("Expected output to contain: %s, but got: %s", expectedOutput, output)
	}
}

// captureOutput captures the standard output of a function
func captureOutput(f func()) string {
	// Save the original stdout
	oldStdout := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture the output in a goroutine
	capturedOutput := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		capturedOutput <- buf.String()
	}()

	// Call the function
	f()

	// Close the write end of the pipe and restore the original stdout
	w.Close()
	os.Stdout = oldStdout

	// Wait for the goroutine to finish capturing output
	output := <-capturedOutput

	return output
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}

package main

import (
	"testing"
)

// Checks whether the program is able to successfully check for suffixes in the data file + benign list.
func TestCheckData(t *testing.T) {

	test_data := []string{"test.microsoft.com", "m", "test.multiple.domains.microsoft.com", "other.site.com", "other.site.org", "netflix.com"}

	test_benign_list := [][]string{{"1", "microsoft.com"}, {"2", "netflix.com"}}

	expected_output := []string{"m", "other.site.com", "other.site.org"}

	test_output := check_data(test_data, test_benign_list)

	if !check_test_results(test_output, expected_output) {
		// Test Failed
		t.Error("Test Failed. Expected: ", expected_output, " but recieved: ", test_output)
	}
}

// Verify that the readFiles function is able to read the appropriate files and return them correctly.
// Can use default files, data/test_data_file.txt and data/test_benign_list.csv, or specified files in flags
func TestReadFiles(t *testing.T) {
	input_test_data_file := "data/test_data_file.txt"
	input_test_benign_file := "data/test_benign_list.csv"

	data_results, benign_results := readFiles(input_test_data_file, input_test_benign_file)

	// If no data exists or nothing is read, test failed
	if data_results == nil || benign_results == nil {
		t.Error("Testing failed to read data from files")
	}
}

// Returns True if the test output matches the expected list
func check_test_results(output []string, expected []string) bool {
	if len(output) != len(expected) {
		return false
	}
	for i, output_line := range output {
		if output_line != expected[i] {
			return false
		}
	}
	return true
}

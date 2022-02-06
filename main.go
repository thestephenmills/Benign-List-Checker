package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	var input_data_file string
	var input_benign_file string

	flag.StringVar(&input_data_file, "d", "data/test_data_fie.txt", "Text file containing list of DNS names (newline separated)")
	flag.StringVar(&input_benign_file, "b", "data/test_benign_list.csv", "CSV file containing the benign list")

	flag.Parse()

	dns_data, benign_list := readFiles(input_data_file, input_benign_file)

	output_data := check_data(dns_data, benign_list)

	output_file, err := os.Create("output-data.txt")
	check(err)

	datawriter := bufio.NewWriter(output_file)

	// Write output data to file
	for _, data := range output_data {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	output_file.Close()

}

// Returns variables containing the data from both the data file and benign list
func readFiles(input_data_list string, input_benign_list string) ([]string, [][]string) {
	// Open specified data file
	dns_data_file, err := os.Open(input_data_list)
	check(err)

	// Read dns data to slice
	var dns_data []string
	dns_data_scanner := bufio.NewScanner(dns_data_file)
	for dns_data_scanner.Scan() {
		dns_data = append(dns_data, dns_data_scanner.Text())
	}

	// Open specified benign list
	benign_list_file, err := os.Open(input_benign_list)
	check(err)

	benign_list, err := csv.NewReader(benign_list_file).ReadAll()
	check(err)

	return dns_data, benign_list

}

// Compares each entry in the data_list with the benign list.
// If the entry has a suffix in the benign list, remove from list.
// Returns list of strings that do not have suffix in the benign list.
func check_data(data_list []string, benign_list [][]string) []string {

	var output_data []string

	// Check dns data for suffix in the benign list
	for _, data_line := range data_list {
		// Get next line of data and format
		data_line := strings.ToLower(strings.TrimSpace(string(data_line)))
		output_data = append(output_data, data_line)
		// Check if the data line has a suffix in the benign list. Remove from the output list if true
		for _, val := range benign_list {
			if strings.HasSuffix(data_line, "."+strings.TrimSpace(val[1])) || data_line == strings.TrimSpace(val[1]) {
				output_data = output_data[:len(output_data)-1]
				break
			}
		}
	}
	return output_data
}

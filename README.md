# Mandiant_Programming_Exercise
Programming challenge problem for the interview with Lipyeow/Mandiant. This program accepts a list of DNS names (TXT format) and a benign list (CSV format) then compares the files and outputs a file containing a list of DNS entries that do NOT have a suffix in the benign list. 


### PREREQUISITES

1. Go 1.17 https://go.dev/dl/


### Running

To run the program, use the following command, with the command line arguments to specify files. Default files are included and will be used if other files are not specified.

`go run .\main.go -d data/pcap_data_file.txt -b data/top-1m.csv`


### Testing

To run tests use the following command:

`go test -v`

or 

`go test -v -run TestCheckData`

or 

`go test -v -run TestReadFiles`

### Output

The program will generate/update the file "output-data.txt" which contains a list of DNS entries that do not have a suffix in the benign list
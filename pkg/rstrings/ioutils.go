// package rstrings contains string and file-based operations, such as reading a file into a string
package rstrings

import (
	"encoding/csv"
	"io/ioutil"
	"os"
)

// FileToString reads from the given filename and returns the contents as a string
func FileToString(fname string) (result string, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}
	err = file.Close()

	result = string(b)
	return
}

// CsvToMap takes the given filename, reads the file in CSV format, and returns a 2d map
// first dimension is the column name, second dimension is the row values for that column.
// input CSV file is expected to have a header row.
func CsvToMap(fname string) (result map[string] []string, err error) {
	result = make(map[string] []string)

	file, err := os.Open(fname)
	if err != nil {
		return
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	tmp := make(map[int] []string)
	for i, record := range records { // for each row
		if i == 0 {  // skip the header row
			continue
		} else {  // build the details indexed by column number
			for colNum, entry := range record {
				tmp[colNum] = append(tmp[colNum], entry)
			}
		}
	}

	// build the result map, marrying the column header and number
	for colNum, header := range records[0] {
		result[header] = tmp[colNum]
	}

	return
}
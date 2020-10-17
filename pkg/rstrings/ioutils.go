// Package rstrings contains string and file-based operations, such as reading a file into a string
package rstrings

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
func CsvToMap(fname string) (result map[string][]string, err error) {
	result = make(map[string][]string)

	file, err := os.Open(fname)
	if err != nil {
		return
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	tmp := make(map[int][]string)
	for i, record := range records { // for each row
		if i == 0 { // skip the header row
			continue
		} else { // build the details indexed by column number
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

type OnlyWhitespace struct{
	contents string
}

func (o OnlyWhitespace) Error() string {
	return o.contents
}

// ReadStdInOrArgs is a forgiving function for receiving user input for CLI applications.  It attempts to read from
// std input (e.g. `ls | myCoolProgram`).  If this fails, it takes the `os.arguments`, and joins them with a whitespace.
// If the resulting string is only whitespace (e.g. `echo " " | myCoolProgram`), then the original string is returned
// alongside an error of type OnlyWhitespace.
func ReadStdInOrArgs() (string, error) {
	var text string

	info, err := os.Stdin.Stat()
	// error is tolerable, this is a best-effort function, and we have alternatives
	if err == nil {
		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			text = strings.Join(os.Args[1:], " ")
		}
	}

	if len(text) == 0 {
		reader := bufio.NewReader(os.Stdin)
		var output []rune
		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}
		text = string(output)
	}

	if len(strings.TrimSpace(text)) == 0 {
		return text, OnlyWhitespace{}
	}
	return text, nil
}
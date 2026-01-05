package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EvaluateFilename(filename string, fileType string) (string, error) {
	// check if file with same filename exists already
	fileInfoFilename, _ := os.Stat(fmt.Sprintf("%s.%s", filename, fileType))

	if fileInfoFilename != nil {
		number := 1
		newFilename := fmt.Sprintf("%s (1)", filename)

		for {
			// check if fallback name already exists
			fileInfoFilenameWithSuffix, _ := os.Stat(fmt.Sprintf("%s.%s", newFilename, fileType))
			if fileInfoFilenameWithSuffix == nil {
				filename = newFilename
				break
			}

			// file seems to exist
			parts := strings.Split(fileInfoFilenameWithSuffix.Name(), ".")

			// take filename itself and check if it already contains a number with brackets as suffix
			numberWithBrackets, found := strings.CutPrefix(parts[0], filename+" ")
			if !found {
				return "", fmt.Errorf("Something went wrong with filename stuff")
			}

			// remove brackets from number suffix
			replacer := strings.NewReplacer("(", "", ")", "")
			numberWithBrackets = replacer.Replace(numberWithBrackets)

			var err error
			number, err = strconv.Atoi(numberWithBrackets)
			if err != nil {
				return "", err
			}

			number++
			newFilename = fmt.Sprintf("%s (%d)", filename, number)
		}
	}

	return fmt.Sprintf("%s.%s", filename, fileType), nil
}

package files

import (
  "fmt"
  "os"
  "runtime"
)

func removeSection(filename, dbName string) error {
	// Modify config file
  inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	inSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if line == fmt.Sprintf("[%s]", dbName) {
				inSection = true
				continue
			} else {
				inSection = false
			}
		}

		if !inSection {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
  
  // Detect OS?
  opSys = runtime.GOOS
  // Delete folder for DB depending on OS used
  if opSys == "windows" {
    err := os.RemoveAll("C:\\Users\\$USER\\AppData\\goPass\\stores\\local\\" + dbname)
    if err != nil {
      fmt.Println("Error encountered: " + err)
      return err
    }
  }
  if opSys == "linux" || opSys == "darwin" { // Not sure if this will work right on mac?
    err := os.RemoveAll("$HOME/.goPass/stores/local/" + dbname)
    if err != nil {
      fmt.Println("Error Encountered: " + err)
      return err
    }
  } 

	return nil
}


func AddSection(dbName, recipient) error {

}

func UpdateConfig(dbName, recipient, operation) error {
  // if operation is "Create"
    // err := createDB(dbname, recipient)
    // if err != nil {
//       return error
//      }
    // return _
  // 
  // if operation is "Delete"
    // err := removeDB(dbname)
    // if err != nil
    //  return err
    // return _
}


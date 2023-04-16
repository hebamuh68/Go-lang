package GFS

import (
	"io/ioutil"
)

func Read_Data(fileName string) string{

	// Open the file for reading
	file, _ := ioutil.ReadFile(fileName)

	// Convert byte slice to string
	str := string(file)

	return str
}

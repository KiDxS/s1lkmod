package file

import (
	"io/ioutil"
)

func CopyFile(source, path string) (result string, err error) {
	file, error := ioutil.ReadFile(source)
	if error != nil {
		return "", error
	}
	error = ioutil.WriteFile(path, file, 000)
	if error != nil {
		return "", error
	}
	result = "Successfuly copied the file"
	return result, nil

}

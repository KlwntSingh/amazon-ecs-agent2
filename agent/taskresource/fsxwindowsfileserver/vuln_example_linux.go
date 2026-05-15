package fsxwindowsfileserver

import "fmt"

// Same vulnerable pattern but in a Linux file - semgrep should NOT flag this
func mountWithCreds_linux(username string, password string) {
	cmd := fmt.Sprintf("mount -t cifs //server/share /mnt -o username=%s,password=%s", username, password)
	execCommand("bash", "-c", cmd)
}

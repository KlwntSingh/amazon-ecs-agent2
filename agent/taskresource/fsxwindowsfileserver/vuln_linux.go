package fsxwindowsfileserver

import "fmt"

// Same vulnerable pattern but in a Linux file - semgrep should NOT flag this
func testExecInjectionLinux(username string, password string) {
	cmdFormat := "echo '%s' | su -c 'mount -t cifs //server/share /mnt -o username=%s'"
	cmd := fmt.Sprintf(cmdFormat, password, username)
	execCommand("bash", "-c", cmd)
}

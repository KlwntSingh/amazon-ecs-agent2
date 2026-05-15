package example

import "fmt"

// Vulnerable: unsanitized input into execCommand on Linux
func mountShare_linux(server string, path string) {
	cmd := fmt.Sprintf("mount -t nfs %s:%s /mnt", server, path)
	execCommand("bash", "-c", cmd)
}

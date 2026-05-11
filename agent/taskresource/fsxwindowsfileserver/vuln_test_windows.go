package fsxwindowsfileserver

import (
	"fmt"
	"net/http"
	"os/exec"
)

// Vulnerable: unsanitized user input flows directly into exec.Command
func handleRequest_windows(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	cmd := fmt.Sprintf("$(New-Object System.Management.Automation.PSCredential('%s', $(ConvertTo-SecureString 'pass' -AsPlainText -Force)))", username)
	out, _ := exec.Command("powershell.exe", cmd).CombinedOutput()
	w.Write(out)
}

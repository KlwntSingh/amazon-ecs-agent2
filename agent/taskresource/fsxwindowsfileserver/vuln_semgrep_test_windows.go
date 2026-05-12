package fsxwindowsfileserver

import "fmt"

// Simulating the vulnerable pattern for testing semgrep detection
func testExecInjection(username string, password string) {
	psCredentialCommandFormat := "$(New-Object System.Management.Automation.PSCredential('%s', $(ConvertTo-SecureString '%s' -AsPlainText -Force)))"
	credsCommand := fmt.Sprintf(psCredentialCommandFormat, username, password)
	credsArg := fmt.Sprintf("-Credential %s", credsCommand)
	remotePathArg := fmt.Sprintf("-RemotePath '%s'", username)

	args := []string{
		"New-SmbGlobalMapping",
		credsArg,
		remotePathArg,
	}
	execCommand("powershell.exe", args...)
}

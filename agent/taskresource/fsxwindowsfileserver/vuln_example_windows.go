package fsxwindowsfileserver

import "fmt"

// Vulnerable: unsanitized input interpolated into shell command
func mountWithCreds_windows(username string, password string, remotePath string) {
	psCredentialCommandFormat := "$(New-Object System.Management.Automation.PSCredential('%s', $(ConvertTo-SecureString '%s' -AsPlainText -Force)))"
	credsCommand := fmt.Sprintf(psCredentialCommandFormat, username, password)
	credsArg := fmt.Sprintf("-Credential %s", credsCommand)
	remotePathArg := fmt.Sprintf("-RemotePath '%s'", remotePath)

	args := []string{
		"New-SmbGlobalMapping",
		credsArg,
		remotePathArg,
	}
	execCommand("powershell.exe", args...)
}

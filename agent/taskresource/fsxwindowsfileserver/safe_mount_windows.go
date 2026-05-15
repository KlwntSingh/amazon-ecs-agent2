package fsxwindowsfileserver

// SAFE: uses environment variables to pass credentials instead of string interpolation
// This is the pattern from the fix: https://github.com/aws/amazon-ecs-agent/commit/09f274e0157dcbc777c4bbb2e40b6c55ae6646ad
func mountWithCredsSafe_windows(username string, password string, remotePath string) {
	// Command uses env var references, not interpolated values
	// Credentials are passed via process environment, not command string
	execCommand("powershell.exe", "-Command",
		"New-SmbGlobalMapping -RemotePath $env:ECS_FSX_REMOTE -Credential (New-Object PSCredential($env:ECS_FSX_USER, (ConvertTo-SecureString $env:ECS_FSX_PASS -AsPlainText -Force)))")
}

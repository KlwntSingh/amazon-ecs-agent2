package fsxwindowsfileserver

import "fmt"

// String interpolation for logging/display only - no shell execution
// Neither semgrep rule should trigger here
func buildLogMessage_windows(username string, path string) string {
	msg := fmt.Sprintf("User %s accessed path %s", username, path)
	return msg
}

func formatError_windows(code int, detail string) string {
	return fmt.Sprintf("Error %d: %s", code, detail)
}

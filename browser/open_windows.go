// +build windows
package browser

import (
	"os/exec"
	"context"
	"path/filepath"
	"os"
)

var (
	cmd = "url.dll,FileProtocolHandler"
	runDLL32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func open(url string) *exec.Cmd {
	return exec.Command(runDLL32, cmd, url)
}

// +build darwin
package browser

import (
	"os/exec"
)

func open(url string) *exec.Cmd {
	return exec.Command("open", url)
}

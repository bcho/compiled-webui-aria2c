// Display notification with osascript.
package mac

import (
	"fmt"
	"os/exec"
)

const tmpl = "display notification \"%s\" with title \"%s\""

func OSAScriptDisplay(title, content, icon string) error {
	c := exec.Command(
		"/usr/bin/osascript",
		"-e",
		fmt.Sprintf(tmpl, title, content),
	)
	return c.Run()
}

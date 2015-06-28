// Check download status.
package main

import (
	"fmt"
	"log"

	"github.com/bcho/notification/mac"
)

func checkError() {
	tasks, err := listTasks()
	if err != nil {
		log.Println(err)
		mac.OSAScriptDisplay("aria2", "Error occurred!", "")
		return
	}

	if len(tasks.Stopped) <= 0 {
		return
	}

	for _, task := range tasks.Stopped {
		if task.IsError() {
			mac.OSAScriptDisplay(
				"aria2",
				fmt.Sprintf("Task %s failed!", task.Gid),
				"",
			)
			return
		}
	}
}

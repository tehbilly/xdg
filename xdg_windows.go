package xdg

import "os"

func homeFromEnv() {
	if ev := os.Getenv("USERPROFILE"); ev == "" {
		userHome = "~"
	} else {
		userHome = ev
	}
}

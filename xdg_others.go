// +build !windows

package xdg

import "os"

func homeFromEnv() {
	if ev := os.Getenv("HOME"); ev == "" {
		userHome = "~"
	} else {
		userHome = ev
	}
}

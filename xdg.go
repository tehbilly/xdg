// Package xdg helps you respect the XDG Base Directory spec
package xdg

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	// Will be replaced in init() with os-specific path from environment variables
	userHome = ""
	// Computed locations
	dataHome   = ""
	configHome = ""
	cacheHome  = ""
	dataDirs   = []string{""}
	configDirs = []string{""}
	runtimeDir = ""
)

func init() {
	if u, err := user.Current(); err != nil {
		homeFromEnv()
	} else {
		userHome = u.HomeDir
	}

	// Data home
	if ev := os.Getenv("XDG_DATA_HOME"); ev == "" {
		dataHome = filepath.Join(userHome, ".local", "share")
	} else {
		dataHome = ev
	}

	// Config home
	if ev := os.Getenv("XDG_CONFIG_HOME"); ev == "" {
		configHome = filepath.Join(userHome, ".config")
	} else {
		configHome = ev
	}

	// Cache home
	if ev := os.Getenv("XDG_CACHE_HOME"); ev == "" {
		cacheHome = filepath.Join(userHome, ".cache")
	} else {
		cacheHome = ev
	}

	// Data dirs
	if ev := os.Getenv("XDG_DATA_DIRS"); ev == "" {
		dataDirs = []string{"/usr/local/share", "/usr/share"}
	} else {
		dataDirs = strings.Split(ev, ":")
	}

	// Config dirs
	if ev := os.Getenv("XDG_CONFIG_DIRS"); ev == "" {
		configDirs = []string{"/etc/xdg"}
	} else {
		configDirs = strings.Split(ev, ":")
	}

	if ev := os.Getenv("XDG_RUNTIME_DIR"); ev == "" {
		runtimeDir = os.TempDir()
	} else {
		runtimeDir = ev
	}
}

// ConfigHome returns either the directory specified by the XDG_CONFIG_HOME
// environment variable, or ~/.config
func ConfigHome() string {
	return configHome
}

// DataHome returns either the directory specified by the XDG_DATA_HOME
// environment variable, or ~/.local/share
func DataHome() string {
	return dataHome
}

// CacheHome returns either the directory specified by the XDG_CACHE_HOME
// environment variable, or ~/.cache
func CacheHome() string {
	return cacheHome
}

// DataDirs returns either the directories specified by the XDG_DATA_DIRS
// environment variable, or OS-specific defaults.
func DataDirs() []string {
	return dataDirs
}

// ConfigDirs returns either the directories specified by the XDG_CONFIG_DIRS
// environment variable, or OS-specific defaults.
func ConfigDirs() []string {
	return configDirs
}

// RuntimeDir returns either the directory specified by the XDG_RUNTIME_DIR
// environment variable, or the OS-specific version of /tmp. There are no other
// sane defaults.
func RuntimeDir() string {
	return runtimeDir
}

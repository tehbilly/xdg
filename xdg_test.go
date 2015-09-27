package xdg_test

import (
	"fmt"
	"path/filepath"

	"github.com/tehbilly/xdg"
)

func Example() {
	// Each of these will be absolute paths to the equivalient of ~/.config/my-cool-app
	// unless the appropriate environment variables are set.

	// Windows: C:\Users\JohnDoe\.config\my-cool-app
	// Others:  /home/JohnDoe/.config/my-cool-app

	confDir := filepath.Join(xdg.ConfigHome(), "my-cool-app")
	dataDir := filepath.Join(xdg.ConfigHome(), "my-cool-app")

	fmt.Println("My config dir:", confDir)
	fmt.Println("My data dir:", dataDir)
}

// +build solaris

package logrus

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal() bool {
	return true
}

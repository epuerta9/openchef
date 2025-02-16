package version

var (
	// Version is the current version of OpenChef
	Version = "0.1.0"

	// Commit is the git commit hash
	Commit string

	// BuildTime is when the binary was built
	BuildTime string
)

// Info returns version information
func Info() string {
	return Version + " (commit: " + Commit + ", built: " + BuildTime + ")"
}

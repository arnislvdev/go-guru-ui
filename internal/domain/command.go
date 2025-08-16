package domain

// Command represents a CLI command with context and explanation
type Command struct {
	Command     string            `json:"command"`
	Arguments   []string          `json:"arguments,omitempty"`
	Flags       map[string]string `json:"flags,omitempty"`
	Platform    string            `json:"platform"`
	Context     string            `json:"context,omitempty"`
	Explanation string            `json:"explanation"`
}

// Platform constants
const (
	PlatformLinux   = "linux"
	PlatformMacOS   = "macos"
	PlatformWindows = "windows"
	PlatformUnknown = "unknown"
)

// CommandCategory represents the type of command
type CommandCategory struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Examples    []string `json:"examples"`
}

// Common command categories
var CommandCategories = []CommandCategory{
	{
		Name:        "file_management",
		Description: "File and directory operations",
		Examples:    []string{"ls", "cp", "mv", "rm", "mkdir"},
	},
	{
		Name:        "system_info",
		Description: "System information and monitoring",
		Examples:    []string{"ps", "top", "df", "free", "uname"},
	},
	{
		Name:        "network",
		Description: "Network operations and diagnostics",
		Examples:    []string{"ping", "curl", "wget", "netstat", "ifconfig"},
	},
	{
		Name:        "package_management",
		Description: "Software package operations",
		Examples:    []string{"apt", "yum", "brew", "pip", "npm"},
	},
}

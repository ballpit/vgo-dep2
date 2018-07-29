package dep2

// RootFunc returns some data.
func RootFunc() string {
	return "dep val 2"
}

// BreakingChange returns some data.
func BreakingChange() string {
	return "here is a breaking change"
}

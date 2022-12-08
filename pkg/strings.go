package pkg

// StringToPointer take in a string a return the memory address
// we use this method because Golang does not support it directly in return functions
func StringToPointer(s string) *string {
	return &s
}

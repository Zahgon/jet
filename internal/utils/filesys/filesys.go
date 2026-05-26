package filesys

// FormatAndSaveGoFile saves go file at folder dir, with name fileName and contents text.
func FormatAndSaveGoFile(dirPath, fileName string, text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec 304

// if there is a format error we will write unformulated text for debug purposes

// EnsureDirPathExist ensures dir path exists. If path does not exist, creates new path.
func EnsureDirPathExist(dirPath string) error { _ = "STUB: not implemented"; return nil }

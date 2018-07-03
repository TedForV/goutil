package filesystem

import "os"

// IsPathExisted is a func the tell the path is existed or not
func IsPathExisted(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

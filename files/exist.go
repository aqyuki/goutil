package files

import (
	"os"
)

// ExistFile check if the file exists
func ExistFile(path string) (bool, error) {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}

// ExistDir check if the directory exists
func ExistDir(path string) (bool, error) {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false, nil
	} else {
		return true, nil
	}

}

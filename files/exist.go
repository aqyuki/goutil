package files

import (
	"os"
)

func ExistFile(path string) (bool, error) {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}

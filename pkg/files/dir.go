package files

import "os"

func EnsureExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		return err
	} else {
		return err
	}
}

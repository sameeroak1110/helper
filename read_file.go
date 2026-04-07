package helper

import (
	"ioutil"
)

func Read_file(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b)
}

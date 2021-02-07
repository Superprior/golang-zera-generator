package pkg

import (
	"fmt"
	"io/ioutil"
)

func GitIgnoreData(p string) error {
	fmt.Println("Configuring the gitignore...")
	content, err := ioutil.ReadFile("./pkg/file_information/.gitignore.txt")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(p+"/.gitignore", content, 0777)
	if err != nil {
		return err
	}
	return nil
}

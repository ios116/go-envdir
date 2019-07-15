package envdir

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

func RunWithEnv(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		content := string(b)
		fmt.Println(content)

	}
	return nil
}

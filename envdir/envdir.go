package envdir

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func RunWithEnv(dir string, prog string) (b []byte, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	if prog == "" {
		return b, errors.New("programma is empty")
	}
	cmd := exec.Command(prog)
	envs := make([]string, len(files))

	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		b, err = ioutil.ReadFile(filePath)
		if err != nil {
			return
		}
		env := strings.Join([]string{file.Name(), string(b)}, "=")
		envs = append(envs, env)
	}
	cmd.Env = append(os.Environ(), envs...)
	b, err = cmd.Output()
	if err != nil {
		return
	}
	return
}

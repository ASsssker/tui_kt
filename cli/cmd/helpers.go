package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func clearDir(name string) RunResMsg {
	var err error
	dir, err := os.ReadDir(name)
	if err != nil {
		return createErrMsg(err)
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{name, d.Name()}...))
		if err != nil {
			return createErrMsg(err)
		}
	}
	return Successfully
}

func getHomeDir() (string, error) {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if strings.Contains(str, "1000") {
			username := strings.Split(str, ":")[0]

			return fmt.Sprintf("/home/%s", username), nil
		}
	}

	return "", fmt.Errorf("client dir not found")
}

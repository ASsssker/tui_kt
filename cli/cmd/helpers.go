package cmd

import (
	"os"
	"path"
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
package utils

import (
	"io/ioutil"
)

func WalkDir(path string, test_fun func(string2 string), excludeDirs ...string) {
	for _, ed := range excludeDirs {
		if path == ed {
			return
		}
	}

	dir, err2 := ioutil.ReadDir(path)
	if err2 != nil {
		println(err2.Error())
	}
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			WalkDir(path+fileInfo.Name()+"/", test_fun, excludeDirs...)
		} else {
			test_fun(path + fileInfo.Name())
		}
	}
}

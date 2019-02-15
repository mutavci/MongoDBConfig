package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func MongoServiceInstall() {
	pwd, _ := os.Getwd()
	startmongo := path.Join(pwd, "mongodb-win32-x86_64-2008plus-ssl-4.0.5", "bin") + "/mongod.exe --config " + path.Join(pwd, "data/conf.yaml --install")
	fmt.Println(path.Clean(startmongo))
	startmongodexe := path.Join(pwd, "mongodb-win32-x86_64-2008plus-ssl-4.0.5", "bin") + "/mongod.exe"
	cmd := exec.Command(path.Clean(startmongodexe), "--config", path.Join(pwd, "data/conf.yaml"), "--install")
	out, errcmd := cmd.CombinedOutput()
	if errcmd != nil {
		fmt.Println(errcmd)
	}
	fmt.Println(string(out))

}

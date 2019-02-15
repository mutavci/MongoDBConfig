package main

import (
	"fmt"
)

func main() {
	Config := MongodbConf{}
	Login := MongodbLogin{}
	Download("mongodb-win32-x86_64-2008plus-ssl-4.0.5.zip")
	ConfigFile(&Config)
	Service("MongoDB", "stop")
	MongoServiceInstall()
	Service("MongoDB", "start")
	LoginConfig(&Config, &Login)
	Service("MongoDB", "stop")
	ConfigSecurty(&Config)
	Service("MongoDB", "start")
	fmt.Println("finished")
}

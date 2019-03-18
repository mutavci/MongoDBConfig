package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

//ConfigFile Dosya yollarını ve configrasyon dosyasını yarlıyor.
func ConfigFile(Config *MongodbConf) {

	pwd, err := os.Getwd()
	if _, err := os.Stat("./data/"); os.IsNotExist(err) {
		os.Mkdir("data", os.ModePerm)
	}
	if _, err := os.Stat("./data/db"); os.IsNotExist(err) {
		os.Mkdir("./data/db", os.ModePerm)
	}
	log := path.Clean(path.Join(pwd, "data", "db", "mongodb.log"))
	db := path.Clean(path.Join(pwd, "data", "db"))
	ConfigValues, err := ioutil.ReadFile("./data/conf.yaml") // just pass the file name
	if err != nil {
		ConfigValues = []byte(Exampledata)
	}

	yaml.Unmarshal([]byte(ConfigValues), &Config)
	Config.SystemLog.Path = log
	Config.Storage.DbPath = db
	Config.Security.Authorization = "disabled"
	d, _ := yaml.Marshal(&Config)
	fileHandle, _ := os.Create("./data/conf.yaml")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	fmt.Fprintln(writer, string(d))
	writer.Flush()
}

//ConfigSecurty Guvenlik Aktif Etmek için
func ConfigSecurty(Config *MongodbConf) {
	Config.Security.Authorization = "enabled"
	d, _ := yaml.Marshal(&Config)
	fileHandle, _ := os.Create("./data/conf.yaml")
	writer := bufio.NewWriter(fileHandle)
	fmt.Fprintln(writer, string(d))
	writer.Flush()
}

//LoginConfig kullanıcı adı ve şifreyi yaml dosyasına kaydediyor
func LoginConfig(Config *MongodbConf, Login *MongodbLogin) {
	pwd, err := os.Getwd()
	startmongoexe := path.Join(pwd, "mongodb-win32-x86_64-2008plus-ssl-4.0.5", "bin") + "/mongo.exe"
	username := RandomString(10)
	password := RandomString(10)
	x := `db.createUser({user: "` + username + `",pwd: "` + password + `",roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]})`
	err1 := ioutil.WriteFile(path.Join(pwd, "data")+"/script.js", []byte(x), 0644)

	if err1 != nil {
		fmt.Println(err.Error())
	}
	cmd := exec.Command(startmongoexe, "localhost:"+strconv.Itoa(Config.Net.Port)+"/admin", path.Join(pwd, "data/script.js"))
	out, errcmd := cmd.CombinedOutput()
	if errcmd != nil {
		fmt.Println(errcmd)
	}
	fmt.Println(string(out))
	ConfigLogin, err := ioutil.ReadFile("./data/login.yaml") // just pass the file name
	if err != nil {
		ConfigLogin = []byte(ExampleLogin)
	}

	yaml.Unmarshal([]byte(ConfigLogin), &Login)
	Login.Login.Username = username
	Login.Login.Password = password
	d, _ := yaml.Marshal(&Login)
	fileHandle, _ := os.Create("./data/login.yaml")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	fmt.Fprintln(writer, string(d))
	writer.Flush()
}

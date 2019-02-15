package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

var Exampledata = `
net:
  bindIp: 0.0.0.0
  port: 27019
security:
  authorization: disabled
storage:
  dbPath: ./data/db
  journal:
    enabled: true
systemLog:
  destination: file
  logAppend: true
  path: ./data/db/mongodb.log
  timeStampFormat: iso8601-utc
`
var ExampleLogin = `
login:
  username: test1
  password: pass1
`

//RandomString Otomatik mongodb user ve password belirlemek için
func RandomString(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	time.Sleep(1 * time.Second)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[seededRand.Intn(len(chars))])
	}
	return b.String()
}

//Service Servis Durdurmak ve Başlatmak
func Service(serviceName, action string) {
	cmd := exec.Command("net", action, serviceName)
	out, errcmd := cmd.CombinedOutput()
	if errcmd != nil {
		fmt.Println(errcmd)
	}
	fmt.Println(string(out))
}

//MongodbConf mongodbnin configrasyon ayarları
type MongodbConf struct {
	Net struct {
		BindIp string `yaml:"bindIp"`
		Port   int    `yaml:"port"`
	} `yaml:"net"`
	Security struct {
		Authorization string `yaml:"authorization"`
	} `yaml:"security"`
	Storage struct {
		DbPath  string `yaml:"dbPath"`
		Journal struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"journal"`
	} `yaml:"storage"`
	SystemLog struct {
		Destination     string `yaml:"destination"`
		LogAppend       bool   `yaml:"logAppend"`
		Path            string `yaml:"path"`
		TimeStampFormat string `yaml:"timeStampFormat"`
	} `yaml:"systemLog"`
}

//MongodbLogin Login Bilgilerini Yaml Dosyasına yazmak için Kullandığımız Model
type MongodbLogin struct {
	Login struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"login"`
}

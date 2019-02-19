package configs

import (
	"encoding/json"
	"io/ioutil"
)

var appConfigs Configs

// Configs App
type Configs struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

// Server Options
type Server struct {
	Port string `json:"port"`
}

// Database Options
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// Init App_Configs
func Init() (err error) {
	// Read File

	data, err := ioutil.ReadFile("app.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &appConfigs); err != nil {
		panic(err)
	}
	return
}

// Get App Configs
func Get() Configs {
	return appConfigs
}

package cmd

import (
	"bytes"
	_ "crypto/sha512"
	"crypto/tls"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"strings"
)

func PING() {
	if viper.GetBool("ping") == true {
		ticker := time.NewTicker(20 * time.Second)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					runPing()
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}

}


func runPing(){
	switch viper.GetString("env") {
	case "dev":
		go req(viper.GetString("dev-api"))
	case "prod":
		go req(viper.GetString("api"))
	default:
		return
	}
}
//StatusResponse is a type
type Device struct {
	Hostname string            `json:"hostname"`
	Alias    string            `json:"alias"`
	Channels int               `json:"channels"`
	Owner    string            `json:"owner"`
	Location map[string]string `json:"location"`
}

func (device Device) addLocation() {
	location := make(map[string]string, viper.GetInt("channels"))

	for k := range make([]int, viper.GetInt("channels")) {
		location[strconv.Itoa(k + 1)] = "undefined"
	}
	device.Location = location

}

func req(url string) {

	device := Device{
		Hostname: viper.GetString("device"),
		Channels: viper.GetInt("channels"),
		Owner:    viper.GetString("owner"),
		Alias:    "BRANDNEW",
	}
	device.addLocation()

	ERRR, _ := json.Marshal(device)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(ERRR))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Custom-Header", "myvalue")

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(resp.Status, "409") {
		log.Println("Device", viper.GetString("device"), "has just sent an alive request to", url)

	} else {
		log.Println("Device", viper.GetString("device"), "has just registered at", url)
		log.Println("Response body:> ", string(body))

	}

	time.Sleep(60 * time.Second)
}

package cmd

import (
	"net/http"
	"github.com/spf13/viper"
	"time"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"encoding/json"
	"log"
	_ "crypto/sha512"
	"crypto/tls"
)

func PING() {
	if viper.GetBool("ping") == true {
		ticker := time.NewTicker(60 * time.Second)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:

					switch viper.GetString("ENV") {
					case "Development":
						go req(viper.GetString("api"))
						time.Sleep(30 *time.Second)
						go req(viper.GetString("dev-api"))
					default:
						return

					}
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}

}

//StatusResponse is a type
type Device struct {
	Hostname string                `json:"hostname"`
	Alias    string                `json:"alias"`
	Channels int                `json:"channels"`
	Owner    string                `json:"owner"`
	Location map[string]string     `json:"location"`
}

func req(url string) {


	location := make(map[string]string, viper.GetInt("channels"))

	for k := range make([]int, viper.GetInt("channels")) {
		location[strconv.Itoa(k + 1)] = "undefined"
	}


	test := Device{
		Hostname: viper.GetString("device"),
		Channels: viper.GetInt("channels"),
		Owner:viper.GetString("owner"),
		Alias:"BRANDNEW",
		Location:location,
	}

	ERRR, _ := json.Marshal(test)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(ERRR))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Custom-Header", "myvalue")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
		log.Println("Device", viper.GetString("device"),"has just sent an alive request to", url)

	}else{
		log.Println("Device", viper.GetString("device"), "has just registered at", url)
		log.Println("Response body:> ", body)

	}

	time.Sleep(60 * time.Second)
}

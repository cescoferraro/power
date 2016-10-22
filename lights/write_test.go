package lights_test

import (
	"github.com/cescoferraro/power/util"
	"strconv"
	"net/http"
	"log"
	"github.com/fatih/color"
)

func (t *LightsTests) ChangeChannel(arg string, number int) {


	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/"+strconv.Itoa(number)+"/" + arg,
		Status:      http.StatusOK,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}


func (t *LightsTests) ChangeChannelOutOfRange() {


	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/"+strconv.Itoa(99)+"/false" ,
		Status:      http.StatusBadRequest,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}


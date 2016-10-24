package lights_test

import (
	"github.com/cescoferraro/power/util"
	"github.com/fatih/color"
	"log"
	"net/http"
	"strconv"
)

func (t *LightsTests) ChangeChannel(arg string, number int) {

	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/" + strconv.Itoa(number) + "/" + arg,
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
		Path:        "/lights/" + strconv.Itoa(99) + "/false",
		Status:      http.StatusBadRequest,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}

func (t *LightsTests) ActionOutOfRange() {

	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/1/sdfsdf",
		Status:      http.StatusBadRequest,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}

package lights_test

import (
	"github.com/cescoferraro/power/util"
	"github.com/fatih/color"
	"log"
	"net/http"
	"strconv"
)

func (t *LightsTests) ReadChannel(number int) {

	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/" + strconv.Itoa(number),
		Status:      http.StatusOK,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}

func (t *LightsTests) ReadChannelOutOfRange() {

	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/" + strconv.Itoa(88),
		Status:      http.StatusBadRequest,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}

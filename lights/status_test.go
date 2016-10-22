package lights_test

import (
	"github.com/cescoferraro/power/util"
	"net/http"
	"log"
	"github.com/fatih/color"
)

func (t *LightsTests) CurrentStatus() {
	loginUserRegularUser := util.TableTest{
		Method:      "GET",
		Path:        "/lights/status",
		Status:      http.StatusOK,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, loginUserRegularUser)

	log.Println(response)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}


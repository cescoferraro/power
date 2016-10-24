package lights_test

import (
	"github.com/cescoferraro/power/util"
	"github.com/fatih/color"
	"net/http"
)

func (t *LightsTests) HelthEndpoint() {

	healthEndpoint := util.TableTest{
		Method:      "GET",
		Path:        "/lights/health",
		Status:      http.StatusOK,
		Name:        "LoginRegularUser",
		Description: "Should return a token",
	}

	response := util.SpinSingleTableTests(t.Test, healthEndpoint)
	util.NEWLogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}

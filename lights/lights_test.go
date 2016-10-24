package lights_test

import (
	"flag"
	"github.com/cescoferraro/power/cmd"
	"github.com/cescoferraro/power/lights"
	"github.com/cescoferraro/power/util"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"math/rand"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type LightsTests struct{ Test *testing.T }

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func setup() {

	var channels int
	flag.IntVar(&channels, "channels", 8, "number of channels of the testing device")
	flag.Parse()
	cmd.ViperInit()
	viper.Set("channels", channels)
	util.LogIfVerbose("verbose enabled")
	util.RunIfVerbose(util.PrintViperConfig)
	ROUTER := mux.NewRouter()
	lights_router := lights.Routes(ROUTER)
	util.Server = httptest.NewServer(lights_router)

}

func TestRunner(t *testing.T) {

	t.Run("A=health", func(t *testing.T) {
		test := LightsTests{Test: t}
		test.HelthEndpoint()
	})

	t.Run("A=status", func(t *testing.T) {
		test := LightsTests{Test: t}
		test.CurrentStatus()
	})

	t.Run("A=read", func(t *testing.T) {
		test := LightsTests{Test: t}
		test.ReadChannel(2)
		test.ReadChannelOutOfRange()
	})

	t.Run("A=action", func(t *testing.T) {
		test := LightsTests{Test: t}
		all := []string{"false", "true"}
		for i := 0; i < 100; i++ {
			test.ChangeChannel(all[rand.Intn(len(all))], randInt(1, 8))
			time.Sleep(200 * time.Microsecond)
		}
		test.ChangeChannelOutOfRange()
		test.ActionOutOfRange()
	})

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

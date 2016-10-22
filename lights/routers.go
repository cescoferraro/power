package lights

import (
	"github.com/gorilla/mux"
	"github.com/cescoferraro/power/util"
)

func  Routes(router *mux.Router) (*mux.Router){

	lights_router := router.PathPrefix("/lights").Subrouter()

	lights_router.
		Path("/health").
		Methods("OPTIONS","GET").
		Handler(util.Adapt(HealthHandler(router), util.EnableCORS()))

	lights_router.
		Path("/status").
		Handler(util.Adapt(StatusHandler(router), util.EnableCORS()))

	lights_router.
		Path("/{channel}").
		Methods("OPTIONS","GET").
		Handler(util.Adapt(ReadSerialHandler(router),util.EnableCORS()))


	lights_router.Path("/{channel}/{action}").
		Methods("OPTIONS","GET").
		Handler( util.Adapt(SerialHandler(router), util.EnableCORS()))

	return lights_router

}

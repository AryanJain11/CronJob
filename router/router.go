package router

import (
	"CronJob/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/cronjob", controller.CronJobStart).Methods("POST")
	router.HandleFunc("/stop", controller.Stop).Methods("POST")
	// router.HandleFunc("/getAllTags", controller.GetAllTags).Methods("GET")
	// router.HandleFunc("/getAllUsers", controller.GetAllUsers).Methods("GET")
	// router.HandleFunc("/createUser", controller.CreateUser).Methods("POST")
	// router.HandleFunc("/addTags", controller.AddTags).Methods("POST")
	// router.HandleFunc("/tester", controller.Tester).Methods("GET")
	// router.HandleFunc("/updateTags/{id}", controller.UpdateTags).Methods("PUT")
	// router.HandleFunc("/profile/{id}", controller.ProfileUpdate).Methods("PUT")
	// router.HandleFunc("/profile/{id}", controller.Profile).Methods("GET")
	// router.HandleFunc("/updateFriendsList/{id}", controller.UpdateFriendsList).Methods("PUT")
	return router
}

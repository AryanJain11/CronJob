package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

// var cancelFunc context.CancelFunc
var cancelMap = make(map[string]context.CancelFunc)

func CronJobStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var request = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&request)

	// fmt.Println(request["repositaryLink"])
	component := path.Base(request["repositaryLink"].(string))

	// Remove the ".git" extension
	component = strings.TrimSuffix(component, ".git")
	// Create a new ZIP file
	ctx, cancel := context.WithCancel(context.Background())
	cancelMap[component] = cancel
	go runCronJobs(request, ctx)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "cronjob started",
	})
}

func Stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var request = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&request)

	// fmt.Println(request["repositaryLink"])
	component := path.Base(request["repositaryLink"].(string))

	// Remove the ".git" extension
	component = strings.TrimSuffix(component, ".git")

	stopHelper(cancelMap, component)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "cronjob stopped",
	})
}

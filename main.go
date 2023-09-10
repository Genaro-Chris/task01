package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Slack_name      string `json:"slack_name"`
	Current_day     string `json:"current_day"`
	Utc_time        string `json:"utc_time"`
	Track           string `json:"track"`
	Github_file_url string `json:"github_file_url"`
	Github_repo_url string `json:"github_repo_url"`
	Status_code     string `json:"status_code"`
}

func new_result(slack_name string, track string) Result {
	return Result{
		Slack_name:      slack_name,
		Current_day:     time.Now().Weekday().String(),
		Utc_time:        time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		Track:           track,
		Github_file_url: "https://github.com/Genaro-Chris/task01/blob/main/main.go",
		Github_repo_url: "https://github.com/Genaro-Chris/task01",
		Status_code:     "200",
	}
}

func handle(w http.ResponseWriter, r *http.Request) {

	slack_name, track := r.URL.Query().Get("slack_name"), r.URL.Query().Get("track")
	result := new_result(slack_name, track)
	res, _ := json.Marshal(result)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	http.HandleFunc("/api", handle)
	log.Fatal(http.ListenAndServe(":80", nil))
}

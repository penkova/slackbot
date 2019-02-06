package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/apenkova/slackbot/service/slackbot"
	"github.com/apenkova/slackbot/service/structs"
	"io/ioutil"
	"net/http"
)

type versionResponse struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
}

// TasksResponse describes response entity of Tasks handler
type TasksResponse struct {
	State   string `json:"state"`
	Details string `json:"details,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(res)
	if err != nil {

		w.Write([]byte(fmt.Sprintf("{\"state\": \"error\", \"err\": \"%s\"}", err.Error())))
	}
	w.Write(js)
}

func writeError(w http.ResponseWriter, err error, code int) {
	type resp struct {
		State   string `json:"state"`
		Details string `json:"details,omitempty"`
	}
	w.WriteHeader(code)
	writeJSONResponse(w, resp{
		State:   "error",
		Details: err.Error(),
	})
}

// Version returns version number and build time setup in build time
func Version(w http.ResponseWriter, _ *http.Request) {
	js, err := json.Marshal(versionResponse{
		Version:   "1",
		BuildTime: "today",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// Health tests health of the service and return health result
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func PostEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var res = TasksResponse{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var msg structs.MsgText
	if err = json.Unmarshal(body, &msg); err != nil {
		writeError(w, err, 400)
		return
	}

	slackbot.Work.Queue <- msg
	res.State = "done"
	writeJSONResponse(w, &res)
}

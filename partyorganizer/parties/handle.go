package parties

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.dcx.bz/flypath/flypath/src/gomessages/common/glog"
	"io/ioutil"
	"net/http"
	"partyApp/partyorganizer/models/protobuf"
)

//GetParties returns actual, available parties
func GetParties(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, _ = w.Write([]byte("elo"))
}

func CreateParty(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	status := http.StatusInternalServerError
	defer func() {
		w.WriteHeader(status)
	}()

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Error("Could not read request body: ", err)
		return
	}
	req := &models.Party{}
	err = json.Unmarshal(payload, req)
	if err != nil {
		glog.Error("Could not unmarshla body: ", err)
		return
	}
	status = validateInput(req)
}

func validateInput(party *models.Party) int {
	if party.Id == 0 || party.Location == nil || party.User == 0 {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

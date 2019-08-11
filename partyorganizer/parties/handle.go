package parties

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.dcx.bz/flypath/flypath/src/gomessages/common/bqhelper"
	"gitlab.dcx.bz/flypath/flypath/src/gomessages/common/glog"
	"io/ioutil"
	"log"
	"net/http"
	models "partyApp/partyorganizer/models/protobuf"
	"time"
)

type PartiesHandler struct {
	ctx                context.Context
	partiesInserter    *bigquery.Inserter
	partiesTableSchema bigquery.Schema
}

func NewPartiesHandler(ctx context.Context, inserter *bigquery.Inserter, schema bigquery.Schema) *PartiesHandler {
	return &PartiesHandler{
		ctx:                ctx,
		partiesInserter:    inserter,
		partiesTableSchema: schema,
	}
}

//GetParties returns actual, available parties
func (h *PartiesHandler) GetParties(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, _ = w.Write([]byte("elo"))
}

func (h *PartiesHandler) CreateParty(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	status := http.StatusInternalServerError
	req := &models.Party{}
	defer func() {
		if status == http.StatusOK {
			h.writeToBigquery(req)
		}
		w.WriteHeader(status)
	}()

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Error("Could not read request body: ", err)
		return
	}
	err = json.Unmarshal(payload, req)
	if err != nil {
		glog.Error("Could not unmarshla body: ", err)
		return
	}
	status = validateInput(req)
}

func (h *PartiesHandler) writeToBigquery(party *models.Party) {
	now := time.Now()
	saver := &bqhelper.StructSaverWithAdditionalData{
		Schema:     h.partiesTableSchema,
		Struct:     party,
		Additional: map[string]bigquery.Value{"Timestamp": now},
	}

	if err := h.partiesInserter.Put(h.ctx, saver); err != nil {
		log.Fatal("error while writing party to BQ: ", err)
	}
}

func validateInput(party *models.Party) int {
	if party.Id == 0 || party.Location == nil || party.User == 0 {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

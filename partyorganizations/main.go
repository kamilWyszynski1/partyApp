package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"gitlab.dcx.bz/flypath/flypath/src/gomessages/common/glog"
	"log"
	"net/http"
	"os"
	"partyApp/partyorganizations/parties"
	"runtime"
)

var ctx context.Context
var partiesInserter *bigquery.Inserter
var partiesTableSchema bigquery.Schema

type server struct {
	router *httprouter.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *server) handleRequests() {
	s.router.GET("/getParties", parties.GetParties)
	s.router.POST("/createParty", parties.CreateParty)
}

func initializeBigQuery(proj string) {
	bigqueryClient, err := bigquery.NewClient(ctx, proj)
	if err != nil {
		glog.Critical("Could not create bigquery client: ", err)
	}
	partiesInserter = bigqueryClient.Dataset(viper.GetString("BigQuery.PartyDataset")).Table("BigQuery.Tables.Parties").Inserter()
	md, err := bigqueryClient.Dataset(viper.GetString("BigQuery.PartyDataset")).Table("BigQuery.Tables.Parties").Metadata(ctx)
	if err != nil {
		log.Fatal("Could not load bigquery table metadata ", err)
	}
	partiesTableSchema = md.Schema
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Party organization initializaiton")
	viper.SetConfigName("config-dev")
	viper.AddConfigPath("/home/kamil/go/src/partyApp/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Could not load config: ", err)
	}
	ctx = context.Background()

	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		glog.Critical("GOOGLE_CLOUD_PROJECT environment variable must be set.")
		os.Exit(1)
	}
	//initializeBigQuery(proj)

	log.Println("Start receive, on http")
	s := &server{httprouter.New()}
	s.handleRequests()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), s))
}

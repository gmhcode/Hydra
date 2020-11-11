package hydrarestapi

import (
	"log"
	"net/http"

	"github.com/Hydra/HydraConfigurator"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	err := HydraConfigurator.GetConfiguration(HydraConfigurator.JSON, conf, "./hydraweb/apiconfig.json")
	if err != nil {
		log.Println("Error decoding JSON", err)
		return err
	}
	h := newhydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Println("Error connecting to db ", err)
		return err
	}
	http.HandleFunc("/hydracrew/", h.handleHydraCrewRequests)
	return nil
}

func RunAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}
	http.ListenAndServe(":8061", nil)
	return nil
}

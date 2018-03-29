package hydrarestapi

import(
"MasteringGoTutorial/HYDRA/hydraconfigurator"
"log"
"net/http"
)

type DBlayerconfig struct{
	DB string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() error{
	conf := new(DBlayerconfig)
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "../../apiconfig.json")
	if err != nil{
		log.Println("Error decoding JSON", err)
		return err
	}
	h:= NewHydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil{
		log.Println("Error connecting to db", err)
		return err
	}
	http.HandleFunc("/hydracrew", h.handleHydraCrewRequests)
	return nil
}


func RunAPI() error{
	if err := InitializeAPIHandlers(); err != nil{
		return err
	}
	return http.ListenAndServe(":8061", nil)
}
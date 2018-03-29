package hydraportal

import (
	"MasteringGoTutorial/HYDRA/hydraconfigurator"
	"MasteringGoTutorial/HYDRA/hydradblayer"
	"MasteringGoTutorial/HYDRA/hydraweb/hydrarestapi"
	"html/template"
	"log"
	"net/http"
	"crypto/md5"
	"MasteringGoTutorial/HYDRA/hydradblayer/passwordvault"
	"bytes"
)

var hydraWebTemplate *template.Template

func Run() error {
	var err error

	conf := struct {
		Filespath string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}
	err = hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, &conf, "./hydraweb/portalconfig.json")
	if err != nil {
		return err
	}
	//below is called unpacking
	hydraWebTemplate, err = template.ParseFiles(conf.Templates...)
	if err != nil {
		return err
	}


/*
if !verifyPassword(user, pass){
	hydraWebTemplate.ExecuteTemplate(w, "login.html", nil)
	return
}


*/


	hydrarestapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/Crew/", crewHandler)
	http.HandleFunc("/about/", abouthandler)
	return http.ListenAndServe(":8061", nil)
}

func crewHandler(w http.ResponseWriter, r *http.Request) {
	dblayer, err := hydradblayer.ConnectDatabase("mysql", "gouser:gouser@/Hydra")
	if err != nil {
		return
	}
	all, err := dblayer.AllMembers()
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "crew.html", all)
	if err != nil {
		log.Println(err)
	}
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, &about, "./hydraweb/about.json")
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}
}


func verifyPassword(username, pass string) bool{
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil{
		return false
	}
	defer db.Close()
	data, err := passwordvault.GetPasswordBytes(db, username)
	if err != nil{
		return false
	}
	hashedPass := md5.Sum([]byte(pass))
	return bytes.Equal(hashedPass[:], data)
}
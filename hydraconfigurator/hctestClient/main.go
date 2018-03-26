package main

import(
	"fmt"
	"MasteringGoTutorial/HYDRA/hydraconfigurator"
)

type ConfS struct{
	TS string `name:"testString" xml:"testString" json:"testString"`
	TB bool `name:"testBool" xml:"testBool" json:"testBool"`
	TF float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
	TestInt int
}

func main(){
	configstruct := new(ConfS)

	// hydraconfigurator.GetConfiguration(hydraconfigurator.CUSTOM, configstruct, "configfile.conf")
	hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, configstruct, "configfile.json")
	// hydraconfigurator.GetConfiguration(hydraconfigurator.XML, configstruct, "configfile.xml")
	fmt.Println(*configstruct)


	//below are some basic tests to check if the data is correct
	if configstruct.TB{
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)

}
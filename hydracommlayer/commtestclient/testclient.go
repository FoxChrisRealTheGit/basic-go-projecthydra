package main

import(
	"MasteringGoTutorial/HYDRA/hydracommlayer"
	// "MasteringGoTutorial/HYDRA/hydracommlayer/hydraproto"
	"flag"
	"log"
	"strings"
)

func main(){
	op:= flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? hsot:port ")
	flag.Parse()

	switch strings.ToUpper(*op){
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(dest string){
	c:= hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil{
		log.Fatal(err)
	}
	for msg := range recvChan{
		log.Println("Recieved: ", msg)
	}
}
func runClient(dest string){
	// c:= hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	// ship:= &hydraproto.Ship{
	// 	Shipname: "Hydra",
	// 	CaptainName: "hello",
	// 	Crew: []*hydraproto.Ship_CrewMember{

	// 	},
	// }
}
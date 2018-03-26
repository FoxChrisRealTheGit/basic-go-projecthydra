package main


import(
	"fmt"
	"log"
	"os"
)

// func main(){
// 	GenerateFileStatusReport("testfile.txt")
// }

func GenerateFileStatusReport(fname string){

	filestats, err := os.Stat("test3.txt")
	PrintFatalError(err)

	fmt.Println("What's the file name?", filestats.Name())
	fmt.Println("Am I a directry?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("What's the file size?", filestats.Size())
	fmt.Println("When was the last time the file was modified?", filestats.ModTime())
}

func PrintFatalError(err error){
	if err != nil{
		log.Fatal("Error happened while processing file", err)
	}
}
package main


import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close()

	//create a new file
	f2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer f2.Close()
	
	//open file for read write
	f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and Write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create if none exist
	// os.O_TRUNC // Truncate file when opening
	// os.O_CREATE|os.O_RDWR|os.O_WRONLY

	//0666 => Owner: (read & wrte), Group: (read & write), and other (read & write)
	PrintFatalError(err)
	defer f3.Close()

	//rename a file
	err = os.Rename("test1.txt", "test1New.txt")
	PrintFatalError(err)

	//move a file
	err = os.Rename("./test.txt","./testfolder/test1.txt")
	PrintFatalError(err)

	//copy a file
	CopyFile("test3.txt", "./testfolder/test3.txt")

	//delete a file
	err = os.Remove("test2.txt")
	PrintFatalError(err)

	bytes, err := ioutil.ReadFile("test3.txt")
	fmt.Println(string(bytes))

	scanner := bufio.NewScanner(f3)
	count :=0
	for scanner.Scan(){
		count ++
		fmt.Println("found line:", count, scanner.Text())
	}

	//buffered write, efficient store in memory, saves disk I/O
	writebuffer := bufio.NewWriter(f3)
	for i := 1; i<=5; i++{
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writebuffer.Flush()

	GenerateFileStatusReport("test3.txt")

	// filestat1, err := os.Stat("test2.txt")
	// PrintFatalError(err)
	// for{
	// 	time.Sleep(1 * time.Second)
	// 	filestat2, err := os.Stat("test3.txt")
	// }


}

// func PrintFatalError(err error){
// 	if err != nil{
// 		log.Fatal("Error happened while processing file", err)
// 	}
// }


func CopyFile(fname1, fname2 string){
	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	//copy bytes from source to destination
	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	//flush file contents to desc
	err = fNew.Sync()
	PrintFatalError(err)
}

// func GenerateFileStatusReport(fname string){
// 	//Stat returns file info. It wil return an error if there is no file

// 	filestats, err := os.Stat("test3.txt")
// 	PrintFatalError(err)

// 	fmt.Println("What's the file name?", filestats.Name())
// 	fmt.Println("Am I a directry?", filestats.IsDir())
// 	fmt.Println("What are the permissions?", filestats.Mode())
// 	fmt.Println("What's the file size?", filestats.Size())
// 	fmt.Println("When was the last time the file was modified?", filestats.ModTime())
// }
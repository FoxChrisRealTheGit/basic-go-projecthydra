package main

// import(
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func main(){
// 	file, err := os.Open("cfile.csv")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	r := csv.NewReader(file)
// 	r.Comment = '#'
// 	//r.Comma = ";"

// 	/*
// 		records, err := r.ReadAll()
// 		if err != nil {
// 			lof.Fatal(err)
// 		}
// 		fmt.Println(records)
// 		*/

// 		// below is line by line code
// 		for{
// 			record, err := r.Read()
// 			if err == io.EOF{
// 				break
// 			}
// 			if err != nil{
// 				if pe, ok := err.(*csv.ParseError); ok{
// 					fmt.Println("bad column:", pe.Column)
// 					fmt.Println("bad line:", pe.Line)
// 					fmt.Println("Error reported", pe.Err)
// 					if pe.Err == csv.ErrFieldCount{
// 						continue
// 					}
// 				}
// 				log.Fatal(err)
// 			}
// 			fmt.Println("CSV Row:", record)
			
// 			// grabs the int in the csv and multiplies it
// 			i, err := strconv.Atoi(record[1])
// 			if err != nil{
// 				log.Fatal(err)
// 			}
// 			fmt.Println(i*4)
			
// 		}
// }
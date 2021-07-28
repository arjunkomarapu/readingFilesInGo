package main

import (
	"fmt"
	"os"
	//"io"
	//"io/ioutil"
	"bufio"
)
/*
//func for FileStat
func fileStat(){
	f, err := os.Open("reading-FilesInGo.html")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(f.Stat())
}
*/

/*
//Func for FileConcept
func fileConcept(){
	concept, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(concept))
}
*/

/*
//Reading File By Chunks
func readFileByChunks(){
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	//declaring chunkSize
	const chunkSize = 30
	
	//create a buffer
	buffer := make([]byte, chunkSize)
	fmt.Printf("Type of buffer is = %T\n", buffer)
	
	for {
		readTotal, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
			   fmt.Println(err)
			}
			break
		}
		fmt.Println(string(buffer[:readTotal]))
	}
}
*/

/*
//Func readFile line by line
func readFileLineByLine(){
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	
	if err := scanner.Err();err != nil {
	       fmt.Println(err)
	   }
}
*/

//Func readFile Word by Word
func readFileWordbyWord(){
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	
	if err := scanner.Err();err != nil {
	       fmt.Println(err)
	   }
}
func main(){
	//fileStat()
	//fileConcept()
	//readFileByChunks()
	//readFileLineByLine()
	readFileWordbyWord()
}

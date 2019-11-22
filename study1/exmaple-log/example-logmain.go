package main

import (
	"fmt"
	"os"
)

type filelog struct{

	file *os.File
}

func (f *filelog) setfile(filename string){

	var err error
	f.file , err = os.Create(filename)
	if err != nil{

		fmt.Println(err)
	}

}

func (f *filelog) Write(line []byte) error {


	n,err := f.file.Write(line)

	fmt.Println(n)
	if err != nil{

		return err
	}

	return nil


}



type consolelog struct{

}


func (f *consolelog) Write(line []byte) error {

	_, _ = os.Stdout.Write(line)

	return nil


}



func main(){

	var fl *filelog = &filelog{}

	fl.setfile("/tmp/fl.log")

	var cl *consolelog = &consolelog{}

	var mylog *log = &log{}

	mylog.register(fl)
	mylog.register(cl)

	mylog.Writelog([]byte("this is a test!"))









}





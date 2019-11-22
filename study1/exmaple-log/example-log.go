package main

import "fmt"

type LogWriter interface {
	Write([]byte) error
}

type log struct{

	logwriter_list []LogWriter

}


func (l *log) register(lw LogWriter){

	l.logwriter_list = append(l.logwriter_list,lw)


}

func (l *log) Writelog(line []byte)  {

	for _, i := range l.logwriter_list{

		err := i.Write(line)

		if err!=nil{

			fmt.Println(err)
		}

	}


}
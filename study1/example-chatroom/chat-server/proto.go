package main


type Message struct{

	Cmd string `json:"cmd"`
	Data string `json:"data"`

}


type Logincmd struct{
	Id string `json:"userid"`
	Passwd string `json:"passwd"`


}


type Registercmd struct{

	Id string `json:"userid"`
	Passwd string `json:"passwd"`

}


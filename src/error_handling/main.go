package main

import (
	"fmt"
)

type Status int 

const (  
	InvalidLogin Status = iota + 1 
	NotFound  
) 

type StatusErr struct { 
	Status  Status 
	Message string 
}

func (se StatusErr) Error() string {
	return se.Message
} //listend2

func LoginAndGetData(uid, pwd, file string) ([]byte, error) { 
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:    InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:    NotFound,
			Message: fmt.Sprintf("file %s not found", file), 
		}
	}
	return data, nil
}


func GenerateError(flag bool) error { 
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
} 

func login(uid, pwd string) error {
	return nil
}

func getData(file string) ([]byte, error) {
	var dummy []byte;
	return dummy, nil
}
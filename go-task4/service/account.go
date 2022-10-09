package service

import (
	"github.com/WellyShadow/Golang_Homework/go-task4/repository"
)

func CreateUser(id, name, surname, phone string) {
	rep := repository.ConnectBD()
	rep.InputBD(id, name, surname, phone)
	//rep.OutputBD()
}

func AddPhone(id, phone string) {
	rep := repository.ConnectBD()
	rep.InputPhoneBD(id, phone)
	//rep.OutputBD()
}

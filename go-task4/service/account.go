package service

import (
	"github.com/WellyShadow/Golang_Homework/go-task4/repository"
)

type User struct {
	id      string   `bson:"_id" json:"id"`
	name    string   `json:"name"`
	surName string   `json:"surName"`
	phones  []string `json:"phones"`
}

func CreateUser(id, name, surname, phone string) {
	rep := repository.ConnectDBpostgres()
	rep.InputDBpostgres(id, name, surname, phone)
	//coll := repository.ConnectDBmongo()
	Migration(id)
	//rep.OutputDBpostgres()
}

func AddPhone(id, phone string) {
	rep := repository.ConnectDBpostgres()
	rep.InputPhoneDBpostgres(id, phone)
	//repository.ConnectDBmongo()
	//coll := repository.ConnectDBmongo()
	Migration(id)
	//rep.OutputDBpostgres()
}

func GetUser(id string) User {
	//rep := repository.ConnectDBpostgres()
	coll := repository.ConnectDBmongo()
	user := User{}
	user.id, user.name, user.surName, user.phones = coll.OutputMongoDB(id)
	//user.phones = rep.OutputPhonesDBpostgres(id)
	//fmt.Println(user.phones)
	//fmt.Println(user.id, user.name, user.surName)
	//fmt.Println(user)
	return user
}

func Migration(id string) {
	rep := repository.ConnectDBpostgres()
	coll := repository.ConnectDBmongo()
	user := User{}
	user.id, user.name, user.surName = rep.OutputDBpostgres(id)
	user.phones = rep.OutputPhonesDBpostgres(id)
	coll.InputDBmongo(user.id, user.name, user.surName, user.phones)
	//fmt.Println(user.id, user.name, user.surName, user.phones)

}

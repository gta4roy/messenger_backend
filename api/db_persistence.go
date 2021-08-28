package api

import (
	"database/sql"
	"gta4roy/messenger/log"
	"gta4roy/messenger/model"

	_ "github.com/go-sql-driver/mysql"
)

type MessageDB struct {
}

func (store *MessageDB) dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "gta4roy"
	dbPass := "71201"
	dbName := "messagedb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Trace.Println("Service got connected to SQL Server")
	return db
}

func (store *MessageDB) Add(messageDetails *model.MessageModel) model.ResponseModel {

	log.Trace.Println("Add Message Request ")
	db := store.dbConn()
	defer db.Close()

	var serviceResponse model.ResponseModel
	serviceResponse.Status = "Success"
	serviceResponse.Message = "Message Saved Successfully"

	insForm, err := db.Prepare("INSERT INTO message(id,user,date,message) VALUES(?,?,?,?)")
	insForm.Exec(messageDetails.Id, messageDetails.Date, messageDetails.User, messageDetails.Message)
	if err != nil {
		serviceResponse.Status = model.CODE_ERROR_IN_SAVING
		serviceResponse.Message = model.CODE_ERROR_IN_SAVING
		log.Trace.Println("Failed to execute insert request", err.Error())
	}

	db.Close()
	return serviceResponse
}

func (store *MessageDB) ListAll() model.MessageModelArray {

	var messageList model.MessageModelArray
	db := store.dbConn()
	defer db.Close()
	selDB, err := db.Query("SELECT * FROM message ORDER BY id DESC")
	if err != nil {
		log.Trace.Println("Failed to execute select all request", err.Error())
	}
	var messageArray []model.MessageModel
	for selDB.Next() {
		var messageObj model.MessageModel
		err = selDB.Scan(&messageObj.Id, &messageObj.Date, &messageObj.Message, &messageObj.User)
		if err != nil {
			log.Trace.Println("Failed to execute Scan  all request", err.Error())
		}

		messageArray = append(messageArray, messageObj)
	}

	messageList.MessageRecords = messageArray
	db.Close()
	return messageList
}

// func (store *AddressDB) Modify(modifyPersonDetails *protocol.ModifyPerson) protocol.ServiceResponse {

// 	db := dbConn()
// 	log.Trace.Println("Update Request")
// 	insForm, err := db.Prepare("UPDATE addressbook SET name=?,phone=?,address=?,city=?,pin=? WHERE id=?")
// 	if err != nil {
// 		log.Trace.Println("Failed to update request", err.Error())
// 	}
// 	insForm.Exec(modifyPersonDetails.ModifiedPerson.Name, modifyPersonDetails.ModifiedPerson.Phone, modifyPersonDetails.ModifiedPerson.Address, modifyPersonDetails.ModifiedPerson.City, modifyPersonDetails.ModifiedPerson.Pin, modifyPersonDetails.ModifiedPerson.Id)
// 	defer db.Close()
// 	var serviceResponse protocol.ServiceResponse
// 	serviceResponse.IsSuccess = true
// 	serviceResponse.Error = ""

// 	return serviceResponse
// }

// func (store *AddressDB) Search(personId *protocol.PersonID) protocol.PersonList {

// 	log.Trace.Println("Search Request")
// 	var personList protocol.PersonList
// 	db := dbConn()
// 	defer db.Close()
// 	selDB, err := db.Query("SELECT * FROM addressbook WHERE id =?", personId.Id)
// 	if err != nil {
// 		log.Trace.Println("Failed to search request", err.Error())
// 	}

// 	var personArray []*protocol.Person

// 	for selDB.Next() {

// 		var personObj protocol.Person
// 		err = selDB.Scan(&personObj.Id, &personObj.Name, &personObj.Address, &personObj.Phone, &personObj.City, &personObj.Pin)
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		personArray = append(personArray, &personObj)
// 	}

// 	personList.PersonsList = personArray
// 	return personList
// }

// func (store *AddressDB) Delete(personId *protocol.PersonID) protocol.ServiceResponse {

// 	log.Trace.Println("Delete Request")
// 	db := dbConn()
// 	delForm, err := db.Prepare("DELETE FROM addressbook WHERE id=?")
// 	if err != nil {
// 		log.Trace.Println("fail to Delete Request", err.Error())
// 	}
// 	delForm.Exec(personId.Id)
// 	defer db.Close()

// 	var serviceResponse protocol.ServiceResponse
// 	serviceResponse.IsSuccess = true
// 	serviceResponse.Error = ""

// 	return serviceResponse
// }

package api

import (
	"encoding/json"
	"gta4roy/messenger/log"
	"gta4roy/messenger/model"
	"io/ioutil"
	"net/http"
)

var databaseProxy MessageDB

func handleGetHealth(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("Health Request")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"UP"}`))
	return
}

func handleAddMessage(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handleAddMessage Request")

	messageDetails, err := parseRequestParams(w, r)
	if err != model.SUCCESS {
		log.Trace.Println("Error in parsing the request")
	}

	log.Trace.Printf("%s %s %s %s", messageDetails.Id, messageDetails.Date, messageDetails.Message, messageDetails.User)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var responseData model.ResponseModel
	responseData = databaseProxy.Add(&messageDetails)
	json.NewEncoder(w).Encode(responseData)

}

// func handleModifyMessage(w http.ResponseWriter, r *http.Request) {
// 	log.Trace.Println("handleModifyAddress Request")

// 	personDetails, err := parseRequestParams(w, r)
// 	if err != model.SUCCESS {
// 		log.Trace.Println("Error in parsing the request")
// 	}

// 	log.Trace.Printf("%s %s %s %s", personDetails.Name, personDetails.Id, personDetails.Phone, personDetails.Address)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	var responseData model.ResponseModel
// 	responseData = serviceProxy.ModifyAddress(personDetails.Id, personDetails)
// 	json.NewEncoder(w).Encode(responseData)

// }
// func handleSearchMessage(w http.ResponseWriter, r *http.Request) {
// 	log.Trace.Println("handleModifyAddress Request")
// 	personId := r.FormValue("id")
// 	log.Trace.Printf("%s", personId)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	var personList model.PersonModelArray
// 	personList.PersonRecords = serviceProxy.SearchAddress(personId)

// 	json.NewEncoder(w).Encode(personList)

// }
func handlePrintAllMessage(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlePrintAllAddress Request")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var messageList model.MessageModelArray
	messageList = databaseProxy.ListAll()

	json.NewEncoder(w).Encode(messageList)
}

// func handleDeleteMessage(w http.ResponseWriter, r *http.Request) {
// 	log.Trace.Println("handleDeleteAddress Request")
// 	personId := r.FormValue("id")
// 	log.Trace.Printf("%s", personId)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	responseData := serviceProxy.DeleteAddress(personId)
// 	json.NewEncoder(w).Encode(responseData)
// }

func parseRequestParams(w http.ResponseWriter, r *http.Request) (model.MessageModel, model.ErrorType) {
	log.Trace.Println("parsing the input parameters ")
	var messageDataRequestSet model.MessageModel

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || reqBody == nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		return messageDataRequestSet, model.WRONG_INPUTS
	}

	err = json.Unmarshal(reqBody, &messageDataRequestSet)
	if err != nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		return messageDataRequestSet, model.WRONG_INPUTS
	}

	return messageDataRequestSet, model.SUCCESS
}

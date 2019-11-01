package server

import (
	"processor"
	"net/http"
	"model"
	"io/ioutil"
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	//"github.com/xeipuuv/gojsonschema"
)

type Handler struct {
	processor *processor.Processor
}

/*func post(response http.ResponseWriter, request *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(newEvent)
}*/


func NewHandler (_processor *processor.Processor) *Handler {
	return &Handler{
		processor: _processor,
	}
}




func (handler *Handler) CreateOrder(response http.ResponseWriter, request *http.Request) {

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Error reading request body", http.StatusNotAcceptable)
	}

	var newTask model.Task
	
    if err := json.Unmarshal(reqBody, &newTask); err != nil {
		fmt.Println(err.Error())
		http.Error(response, "Error decoding json body", http.StatusBadRequest)
	}


	status := handler.processor.CheckRequest(&newTask) 


	if !status{
		http.Error(response, "Error checking request parameters", http.StatusBadRequest)
	}
	

	responseJson, err := json.Marshal(newTask)
	if err != nil {
		http.Error(response, "Error converting results to json", http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(responseJson)

	//response.WriteHeader(http.StatusCreated)

}

func (handler *Handler) GetOrder(response http.ResponseWriter, request *http.Request) {
	/*vars := mux.Vars(request)

	taskId := vars["taskId"]

	task := getTaskOr404(db, taskId, response, r)
	if task == nil {
		return
	}
	respondJSON(w, http.StatusOK, task)*/
}

func (handler *Handler) UpdateOrder(response http.ResponseWriter, request *http.Request) {
	/*vars := mux.Vars(r)

	taskId := vars["taskId"]
	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := getTaskOr404(db, id, w, r)
	if task == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(response, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, task)*/
}

func (handler *Handler) DeleteOrder(response http.ResponseWriter, request *http.Request) {
	/*vars := mux.Vars(request)

	id := vars["orderId"]
	if id == nil {
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)*/
}
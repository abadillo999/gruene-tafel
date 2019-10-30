package server

import (  
	"processor"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

type Handler struct {  
	_processor *processor.Processor
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


func NewHandler (processor *processor.Processor) *Handler {
	return &Handler{
		_processor: processor,
	}
}




func CreateOrder(response http.ResponseWriter, request *http.Request) {
	var newOrder order
	reqBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newOrder)
	orders = append(orders, newOrder)
	w.WriteHeader(http.StatusCreated)



	vars := mux.Vars(r)

	project := getProjectOr404(db, projectTitle, w, r)
	if project == nil {
		return
	}

	task := model.Task{ProjectID: project.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, task)
}

func GetOrder(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	taskId := vars["taskId"]

	task := getTaskOr404(db, taskId, response, r)
	if task == nil {
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func UpdateOder(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(r)

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
	respondJSON(w, http.StatusOK, task)
}

func DeleteOrder(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id := vars["orderId"]
	if id == nil {
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
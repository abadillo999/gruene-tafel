package server

import (
	"processor"
	"net/http"
	"model"
	"io/ioutil"
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

type Handler struct {
	processor *processor.Processor
	schema *gojsonschema.Schema
}

func NewHandler (processor *processor.Processor) *Handler {
	loader := gojsonschema.NewStringLoader(model.TaskSchema)
	// call the target function
	_schema, err := gojsonschema.NewSchema(loader)
	if err != nil {
		fmt.Println("Error processing schema file %s", err.Error())
        panic(err.Error())
    }
	return &Handler{
		processor: processor,
		schema: _schema,
	}
}

func (handler *Handler) CreateTask(response http.ResponseWriter, request *http.Request) {

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Error reading request body", http.StatusNotAcceptable)
		return
	}

	//schema validation
    documentLoader := gojsonschema.NewBytesLoader([]byte(reqBody))

    result, err := handler.schema.Validate(documentLoader)
    if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
    }

	if result.Valid(){
        fmt.Printf("Valid request received.\n")
	}else {
		fmt.Printf("The document is not valid. see errors :\n")
		var errors string
		for _, desc := range result.Errors() {
			errors = ( errors + desc.String())
			fmt.Printf("- %s\n", desc)
        }
		http.Error(response, errors, http.StatusBadRequest)
		return
	}

	//Task creation
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
}

func (handler *Handler) GetTask(response http.ResponseWriter, request *http.Request) {
	/*vars := mux.Vars(request)

	taskId := vars["taskId"]

	task := getTaskOr404(db, taskId, response, r)
	if task == nil {
		return
	}
	respondJSON(w, http.StatusOK, task)*/
}

func (handler *Handler) UpdateTask(response http.ResponseWriter, request *http.Request) {
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

func (handler *Handler) DeleteTask(response http.ResponseWriter, request *http.Request) {
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
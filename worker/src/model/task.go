package model

import (
	//"time"
	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	Id           string               `json:"id"`
	ClientId     string               `json:"client-id"`
	Lang         string               `json:"lang"`
	LangVersion  string               `json:"lang-version"`
	Lib          string               `json:"lib"`
	LibVersion   string               `json:"lib-version"`
	ImageUrls    [] ImageRef	      `json:"image-urls"`
	Script       string               `json:"script"`
	Done         bool                 `json:"done"`
}
type ImageRef struct {
	Url		string               `json:"url"`
	Id      string               `json:"id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

const TaskSchema = `{
    "type": "object",
    "required": [
        "client-id",
        "lang",
        "lang-version",
        "lib",
        "lib-version",
        "image-urls",
        "script"
    ],
    "additionalProperties": false,
    "properties": {
        "client-id": {"type": "string"},
        "lang": {"type": "string"},
        "lang-version": {"type": "string"},
        "lib": {"type": "string"},
        "lib-version": {"type": "string"},
        "image-urls": {
            "type": "array",
            "additionalProperties": false,
            "items": {
                "type": "object",
                "properties": {
                    "url": {"type": "string"},
                    "id": {"type": "string"}
                }
             },
            "required": [
                "url",
                "id"
            ]
        },
        "script": {"type": "string"},
        "done": {"type": "boolean"}
    }
}`

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
/*func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}*/
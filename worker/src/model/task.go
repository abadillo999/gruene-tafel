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
	imageUrls    [] ImageRef	      `json:"image-urls"`
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

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
/*func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}*/
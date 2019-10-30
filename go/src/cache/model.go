package cache

import (
	"time"

	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	id           string               `json:"id"`
	lang         string               `json:"lang"`
	langVersion  string               `json:"lang-version"`
	lib          string               `json:"lib"`
	libVersion   string               `json:"lib-version"`
	//TODO: imageUrls    map[string]string	  `json:"imageUrls"`
	script       string               `json:"script"`
	done         bool                 `json:"done"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
// Package repository provides the database operations for the todo endpoint.
package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

// Todo is the repository for the todo endpoint.
type Todo interface {
	Create(t *model.Todo) error
	Delete(id int) error
	Update(t *model.Todo) error
	Find(id int) (*model.Todo, error)
	FindAll(task string, status string, sort bool) ([]*model.Todo, error)
}

type todo struct {
	db *gorm.DB
}

// NewTodo returns a new instance of the todo repository.
func NewTodo(db *gorm.DB) Todo {
	return &todo{
		db: db,
	}
}

func (td *todo) Create(t *model.Todo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Update(t *model.Todo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Delete(id int) error {
	result := td.db.Where("id = ?", id).Delete(&model.Todo{})
	if result.RowsAffected == 0 {
		return model.ErrNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	log.Info("Deleted todo with id: ", id)
	return nil
}

func (td *todo) Find(id int) (*model.Todo, error) {
	var todo *model.Todo
	err := td.db.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (td *todo) FindAll(task string, status string, sort bool) ([]*model.Todo, error) {
	var todos []*model.Todo
	query := td.db.Model(&todos)

	if task != "" {
		query.Where("task LIKE ?", "%"+task+"%")
	}

	if status != "" {
		if !model.IsValidStatus(status) {
			return nil, fmt.Errorf("invalid status: %s", status)
		}
		query.Where("status LIKE ?", "%"+status+"%")
	}

	if sort {
		query.Order("priority DESC")
	}

	err := query.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

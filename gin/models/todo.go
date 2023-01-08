package models

import (
	"errors"
)

type Todo struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"type:varchar(100);not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

func GetTodos() ([]Todo, error) {
	var todos []Todo

	if err := DB.Find(&todos).Error; err != nil {
		return todos, errors.New("todos not found")
	}

	return todos, nil
}

func (todo *Todo) AddTodo() error {

	if err := DB.Create(&todo).Error; err != nil {
		return errors.New("todo not created")
	}

	return nil
}

func GetTodoByID(id string) (Todo, error) {
	var todo Todo

	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, errors.New("todo not found")
	}

	return todo, nil
}

func (todo *Todo) UpdateTodo(id string) error {

	if err := DB.Where("id = ?", id).Updates(&todo).Error; err != nil {
		return errors.New("todo not updated")
	}

	return nil
}

func (todo *Todo) DeleteTodo(id string) error {

	if err := DB.Where(id).Delete(&todo).Error; err != nil {
		return errors.New("todo not deleted")
	}

	return nil
}

package repository

import (
	"errors"
	"gorm.io/gorm"
	"todo-api/dto"
	"todo-api/model"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {

	return &Repository{
		db: db,
	}

}

func (r *Repository) Get(id int64) (*model.TodoModel, error) {
	var todo model.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", id).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return &todo, errors.New("invalid id")
	}

	return &todo, nil
}

func (r *Repository) All() (*[]model.TodoModel, error) {
	var todos []model.TodoModel
	db := r.db.Model(&todos)

	checkTodo := db.Debug().Select("*").Find(&todos)

	if checkTodo.Error != nil {
		return &todos, errors.New("not found")
	}

	return &todos, nil
}

func (r *Repository) Create(dto *dto.CreateDTO) (*model.TodoModel, error) {
	var todo model.TodoModel
	db := r.db.Model(&todo)

	todo.Title = dto.Title
	todo.Notes = dto.Notes

	createTodo := db.Debug().Create(&todo)
	db.Commit()

	if createTodo.Error != nil {
		return &todo, errors.New("failed to create a todo")
	}

	return &todo, nil
}

func (r *Repository) Delete(id int64) error {
	var todo model.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", id).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return errors.New("invalid id")
	}

	deleteTodo := db.Debug().Delete(&todo)
	db.Commit()

	if deleteTodo.Error != nil {
		return errors.New("failed to delete a todo")
	}

	return nil
}

func (r *Repository) Update(dto *dto.UpdateDTO) (*model.TodoModel, error) {
	var todo model.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", dto.ID).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return &todo, errors.New("invalid id")
	}

	if dto.Title != "" {
		todo.Title = dto.Title
	}

	if dto.Done != nil {
		todo.Done = *dto.Done
	}

	updateTodo := db.Debug().Select("title", "done").Where("id = ?", dto.ID).Updates(todo)
	db.Commit()

	if updateTodo.Error != nil {
		return &todo, errors.New("failed to create a todo")
	}

	return &todo, nil
}

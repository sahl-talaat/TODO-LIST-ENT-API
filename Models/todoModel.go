package models

import (
	"context"
	config "todoList/Config"
	"todoList/ent"
)

func GeTodos(ctx context.Context) (todo []*ent.Todo, err error) {
	todo, err = config.DB.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func GeTodo(ctx context.Context, id int) (todo *ent.Todo, err error) {
	todo, err = config.DB.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func CreateTodo(ctx context.Context, todo *ent.Todo) (err error) {
	_, err = config.DB.Todo.Create().SetTitle(todo.Title).SetContent(todo.Content).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(ctx context.Context, id int, todo *ent.Todo) (err error) {
	_, err = config.DB.Todo.UpdateOneID(id).SetTitle(todo.Title).SetContent(todo.Content).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(ctx context.Context, id int) (err error) {
	err = config.DB.Todo.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

package models

import (
	"blog_echo/app/configs"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Blog) TableName() string {
	return "blogs"
}

func (m *Blog) Validate() bool {
	if m.Title != "" && m.Body != "" {
		return true
	}
	return false
}

func (m *Blog) All() ([]Blog, error) {
	blog := Blog{}
	blogs := []Blog{}

	rows, err := configs.DB.Query("SELECT * FROM " + m.TableName())
	for rows.Next() {
		rows.Scan(&blog.Id, &blog.Title, &blog.Body, &blog.CreatedAt, &blog.UpdatedAt)
		blogs = append(blogs, blog)
	}

	return blogs, err
}

func (m *Blog) FindById(id int) Blog {
	blog := Blog{}

	row := configs.DB.QueryRow("SELECT * FROM "+m.TableName()+" WHERE id = ?", id)
	row.Scan(&blog.Id, &blog.Title, &blog.Body, &blog.CreatedAt, &blog.UpdatedAt)

	return blog
}

func (m *Blog) Save() (sql.Result, error) {
	result, err := configs.DB.Exec("INSERT INTO "+m.TableName()+" (title, body, created_at, updated_at) VALUES (?, ?, NOW(), NOW())",
		m.Title, m.Body)
	return result, err
}

func (m *Blog) Update() (sql.Result, error) {
	result, err := configs.DB.Exec("UPDATE "+m.TableName()+" SET title = ?, body = ? WHERE id = ?", m.Title, m.Body, m.Id)
	return result, err
}

func (m *Blog) Delete() (sql.Result, error) {
	result, err := configs.DB.Exec("DELETE FROM "+m.TableName()+" WHERE id = ?", m.Id)
	return result, err
}

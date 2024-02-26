package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name, description, categoryId string) (*Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryId)

	if err != nil {
		return nil, err
	}

	return &Course{ID: id, Name: name, Description: description}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var id, name, description, categoryId string

		err := rows.Scan(&id, &name, &description, &categoryId)

		if err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: categoryId})
	}

	return courses, err
}

func (c *Course) FindAllByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses WHERE category_id = $1", categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var id, name, description string

		err := rows.Scan(&id, &name, &description)

		if err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: categoryId})
	}

	return courses, err
}

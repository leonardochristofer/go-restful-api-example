package model

import (
	"context"
	"errors"
	"time"
)

type Cake struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      *float64   `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (c Cake) Validate() error {
	if c.Title == "" {
		return errors.New("Title Required")
	}
	if c.Description == "" {
		return errors.New("Description Required")
	}
	if c.Rating == nil {
		return errors.New("Rating Required")
	}
	return nil
}

//go:generate mockery --name CakeInterface --output ../mocks --filename cakes.go

type CakeInterface interface {
	GetAllCakes(ctx context.Context) ([]Cake, error)
	GetCake(ctx context.Context, id uint64) (Cake, error)
	CreateCake(ctx context.Context, cake Cake) error
	UpdateCake(ctx context.Context, id uint64, cake Cake) error
	DeleteCake(ctx context.Context, id uint64) error
}

type CakeRepository struct{}

func (_ CakeRepository) GetAllCakes(ctx context.Context) ([]Cake, error) {
	var cakes []Cake
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM cakes ORDER BY rating DESC, title ASC`
	rows, err := db.Query(query)

	if err != nil {
		return cakes, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title, description, image string
		var rating *float64
		var createdAt time.Time
		var updatedAt, deletedAt *time.Time
		err := rows.Scan(&id, &title, &description, &rating, &image, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return cakes, err
		}

		cake := Cake{
			ID:          id,
			Title:       title,
			Description: description,
			Rating:      rating,
			Image:       image,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   deletedAt,
		}

		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func (_ CakeRepository) GetCake(ctx context.Context, id uint64) (Cake, error) {
	var cake Cake
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM cakes WHERE id=?`
	row, err := db.Query(query, id)

	if err != nil {
		return cake, err
	}
	defer row.Close()

	if row.Next() {
		var id uint64
		var title, description, image string
		var rating *float64
		var createdAt time.Time
		var updatedAt, deletedAt *time.Time
		err := row.Scan(&id, &title, &description, &rating, &image, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return cake, err
		}

		cake = Cake{
			ID:          id,
			Title:       title,
			Description: description,
			Rating:      rating,
			Image:       image,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   deletedAt,
		}
	}

	return cake, nil
}

func (_ CakeRepository) CreateCake(ctx context.Context, cake Cake) error {
	query := `INSERT INTO cakes(title, description, rating, image, created_at) values(?, ?, ?, ?, ?);`
	_, err := db.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (_ CakeRepository) UpdateCake(ctx context.Context, id uint64, cake Cake) error {
	query := `UPDATE cakes SET title=?, description=?, rating=?, image=?, updated_at=? WHERE id=?;`
	_, err := db.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (_ CakeRepository) DeleteCake(ctx context.Context, id uint64) error {
	query := `UPDATE cakes SET deleted_at=? WHERE id=?;`
	_, err := db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

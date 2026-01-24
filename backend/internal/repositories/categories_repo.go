package repositories

import (
	"context"
	"log"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"
)

func GetAllCategories() ([]dtos.CategoryDto, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT id, name, slug, description 
		FROM categories 
		ORDER BY name`)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.CategoryDto

	for rows.Next() {
		var cat dtos.CategoryDto
		if err := rows.Scan(&cat.Id, &cat.Name, &cat.Slug, &cat.Description); err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, cat)
	}

	return results, nil
}

func GetCategoryBySlug(slug string) (*models.Category, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	var cat models.Category

	err = pool.QueryRow(ctx,
		`SELECT id, name, slug, description, created_at 
		FROM categories 
		WHERE slug = $1`, slug).
		Scan(&cat.Id, &cat.Name, &cat.Slug, &cat.Description, &cat.CreatedAt)

	if err != nil {
		log.Printf("Category not found for slug '%s': %v", slug, err)
		return nil, err
	}

	return &cat, nil
}

func GetCategoryById(id int32) (*models.Category, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	var cat models.Category

	err = pool.QueryRow(ctx,
		`SELECT id, name, slug, description, created_at 
		FROM categories 
		WHERE id = $1`, id).
		Scan(&cat.Id, &cat.Name, &cat.Slug, &cat.Description, &cat.CreatedAt)

	if err != nil {
		log.Printf("Category not found for id %d: %v", id, err)
		return nil, err
	}

	return &cat, nil
}

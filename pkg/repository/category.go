package repository

import (
	"database/sql"
	"errors"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	DB *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{DB: db}
}

func (cp *CategoryPostgres) GetAll(language string) ([]onlinedilerv3.Category, error) {
	query := `
		SELECT c.id, ct.language_code, ct.name, ct.description
		FROM categories c
		INNER JOIN categories_translations ct ON c.id = ct.category_id
		WHERE ct.language_code = $1
	`

	rows, err := cp.DB.Query(query, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make(map[int]*onlinedilerv3.Category)

	for rows.Next() {
		var (
			categoryID  int
			langCode    string
			name        string
			description string
		)

		if err := rows.Scan(&categoryID, &langCode, &name, &description); err != nil {
			return nil, err
		}

		if cat, ok := categories[categoryID]; ok {
			cat.Translations = append(cat.Translations, onlinedilerv3.CategoryTranslations{
				ID:           0, // You can set this appropriately if needed
				CategoryID:   categoryID,
				LanguageCode: langCode,
				Name:         name,
				Description:  description,
			})
		} else {
			categories[categoryID] = &onlinedilerv3.Category{
				ID: categoryID,
				Translations: []onlinedilerv3.CategoryTranslations{{
					ID:           0, // You can set this appropriately if needed
					CategoryID:   categoryID,
					LanguageCode: langCode,
					Name:         name,
					Description:  description,
				}},
			}
		}
	}

	result := make([]onlinedilerv3.Category, 0, len(categories))
	for _, cat := range categories {
		result = append(result, *cat)
	}

	return result, nil
}

func (cp *CategoryPostgres) Get(lang string, id int) (onlinedilerv3.Category, error) {
	query := `
		SELECT ct.language_code, ct.name, ct.description
		FROM categories_translations ct
		WHERE ct.category_id = $1 AND ct.language_code = $2
	`

	row := cp.DB.QueryRow(query, id, lang)

	var (
		langCode    string
		name        string
		description string
	)

	if err := row.Scan(&langCode, &name, &description); err != nil {
		if err == sql.ErrNoRows {
			return onlinedilerv3.Category{}, errors.New("category not found")
		}
		return onlinedilerv3.Category{}, err
	}

	category := onlinedilerv3.Category{
		ID: id,
		Translations: []onlinedilerv3.CategoryTranslations{{
			ID:           0, // You can set this appropriately if needed
			CategoryID:   id,
			LanguageCode: langCode,
			Name:         name,
			Description:  description,
		}},
	}

	return category, nil
}

func (cp *CategoryPostgres) Create(input onlinedilerv3.Category) (int, error) {
	// Begin a transaction
	tx, err := cp.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Insert into categories table
	var categoryID int
	err = tx.QueryRow("INSERT INTO categories DEFAULT VALUES RETURNING id").Scan(&categoryID)
	if err != nil {
		return 0, err
	}

	// Insert into categories_translations table
	for _, translation := range input.Translations {
		_, err = tx.Exec(`
			INSERT INTO categories_translations (category_id, language_code, name, description)
			VALUES ($1, $2, $3, $4)`,
			categoryID, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			return 0, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func (cp *CategoryPostgres) Delete(id int) error {
	// Begin a transaction
	tx, err := cp.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Delete from categories_translations table
	_, err = tx.Exec("DELETE FROM categories_translations WHERE category_id = $1", id)
	if err != nil {
		return err
	}

	// Delete from categories table
	_, err = tx.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (cp *CategoryPostgres) Update(id int, input onlinedilerv3.Category) error {
	// Begin a transaction
	tx, err := cp.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Update category translations
	for _, translation := range input.Translations {
		_, err = tx.Exec(`
			UPDATE categories_translations
			SET name = $3, description = $4
			WHERE category_id = $1 AND language_code = $2`,
			id, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (cp *CategoryPostgres) Search(lang, categoryName string) ([]onlinedilerv3.Category, error) {
	query := `
		SELECT c.id, ct.language_code, ct.name, ct.description
		FROM categories c
		INNER JOIN categories_translations ct ON c.id = ct.category_id
		WHERE ct.language_code = $1 AND ct.name ILIKE '%' || $2 || '%'
	`

	rows, err := cp.DB.Query(query, lang, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make(map[int]*onlinedilerv3.Category)

	for rows.Next() {
		var (
			categoryID  int
			langCode    string
			name        string
			description string
		)

		if err := rows.Scan(&categoryID, &langCode, &name, &description); err != nil {
			return nil, err
		}

		if cat, ok := categories[categoryID]; ok {
			cat.Translations = append(cat.Translations, onlinedilerv3.CategoryTranslations{
				ID:           0, // You can set this appropriately if needed
				CategoryID:   categoryID,
				LanguageCode: langCode,
				Name:         name,
				Description:  description,
			})
		} else {
			categories[categoryID] = &onlinedilerv3.Category{
				ID: categoryID,
				Translations: []onlinedilerv3.CategoryTranslations{{
					ID:           0, // You can set this appropriately if needed
					CategoryID:   categoryID,
					LanguageCode: langCode,
					Name:         name,
					Description:  description,
				}},
			}
		}
	}

	result := make([]onlinedilerv3.Category, 0, len(categories))
	for _, cat := range categories {
		result = append(result, *cat)
	}

	return result, nil
}

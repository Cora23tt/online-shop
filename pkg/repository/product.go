package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductPostgres struct {
	DB *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{DB: db}
}

func (s *ProductPostgres) Search(language, productName string) ([]onlinedilerv3.ProductComplex, error) {
	query := fmt.Sprintf(`
		SELECT p.id, p.rating, p.category_id, pt.language_code, pt.name, pt.description
		FROM products p
		LEFT JOIN products_translations pt ON p.id = pt.product_id
		WHERE pt.name ILIKE '%s' AND pt.language_code = '%s'
	`, "%"+productName+"%", language)

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []onlinedilerv3.ProductComplex
	for rows.Next() {
		var product onlinedilerv3.ProductComplex
		err := rows.Scan(&product.ID, &product.Rating, &product.CategoryID, &product.LanguageCode, &product.Name, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductPostgres) SearchWithLimit(limit, offset int, language, productName string) ([]onlinedilerv3.ProductComplex, error) {
	query := `
        SELECT p.id, p.rating, p.category_id, pt.language_code, pt.name, pt.description
        FROM products p
        JOIN products_translations pt ON p.id = pt.product_id
        WHERE pt.language_code = $1 AND pt.name ILIKE $2
        LIMIT $3 OFFSET $4
    `

	rows, err := s.DB.Query(query, language, "%"+productName+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []onlinedilerv3.ProductComplex
	for rows.Next() {
		var product onlinedilerv3.ProductComplex
		err := rows.Scan(
			&product.ID,
			&product.Rating,
			&product.CategoryID,
			&product.LanguageCode,
			&product.Name,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductPostgres) TopRatedWithLimit(limit, offset int, language string) ([]onlinedilerv3.ProductComplex, error) {
	query := `
        SELECT p.id, p.rating, p.category_id, pt.language_code, pt.name, pt.description
        FROM products p
        JOIN products_translations pt ON p.id = pt.product_id
        WHERE pt.language_code = $1
        ORDER BY p.rating DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := s.DB.Query(query, language, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []onlinedilerv3.ProductComplex
	for rows.Next() {
		var product onlinedilerv3.ProductComplex
		err := rows.Scan(
			&product.ID,
			&product.Rating,
			&product.CategoryID,
			&product.LanguageCode,
			&product.Name,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductPostgres) ByCategory(language string, categoryID int) ([]onlinedilerv3.ProductTranslationConsignment, error) {
	query := `
        SELECT p.id, p.rating, p.category_id,
            ptc.language_code, ptc.name, ptc.description,
            c.expiration_date, c.quantity, c.price,
            c.description AS consignment_description, p.image_urls
        FROM
            products p
        JOIN
            products_translations ptc ON p.id = ptc.product_id
        JOIN
            consignments c ON p.id = c.product_id
        WHERE
            p.category_id = $1
            AND ptc.language_code = $2
    `
	rows, err := s.DB.Query(query, categoryID, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []onlinedilerv3.ProductTranslationConsignment
	for rows.Next() {
		var product onlinedilerv3.ProductTranslationConsignment
		var imageUrls pq.StringArray
		err := rows.Scan(
			&product.ID,
			&product.Rating,
			&product.CategoryID,
			&product.LanguageCode,
			&product.Name,
			&product.Description,
			&product.ExpirationDate,
			&product.Quantity,
			&product.Price,
			&product.ConsignmentDescription,
			&imageUrls,
		)
		if err != nil {
			return nil, err
		}
		product.ImageURLs = []string(imageUrls)
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductPostgres) TopRated(language string) ([]onlinedilerv3.ProductComplex, error) {
	query := `
        SELECT p.id, p.rating, p.category_id, pt.language_code, pt.name, pt.description
        FROM products p
        JOIN products_translations pt ON p.id = pt.product_id
        WHERE pt.language_code = $1
        ORDER BY p.rating DESC
    `

	rows, err := s.DB.Query(query, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []onlinedilerv3.ProductComplex
	for rows.Next() {
		var product onlinedilerv3.ProductComplex
		err := rows.Scan(
			&product.ID,
			&product.Rating,
			&product.CategoryID,
			&product.LanguageCode,
			&product.Name,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductPostgres) GetAll(language string) ([]onlinedilerv3.ProductTranslationConsignment, error) {
	query := `
    SELECT
        p.id AS id,
        p.rating AS rating,
        p.category_id AS category_id,
        pt.language_code AS language_code,
        pt.name AS name,
        pt.description AS description,
        c.expiration_date AS expiration_date,
        c.quantity AS quantity,
        c.price AS price,
        c.description AS consignment_description,
        p.image_urls AS image_urls
    FROM products p
    JOIN products_translations pt ON p.id = pt.product_id AND pt.language_code = $1
    JOIN consignments c ON p.id = c.product_id
    `
	var products []onlinedilerv3.ProductTranslationConsignment
	err := s.DB.Select(&products, query, language)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (s *ProductPostgres) Create(product onlinedilerv3.Product, translations []onlinedilerv3.ProductTranslation) (int, error) {
	query := `
		INSERT INTO products (rating, category_id) 
		VALUES ($1, $2) RETURNING id
	`
	var productID int
	err := s.DB.QueryRow(query, product.Rating, product.CategoryID).Scan(&productID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, errors.New("product already exists")
		}
		return 0, err
	}

	queryTranslations := `
		INSERT INTO products_translations (product_id, language_code, name, description) 
		VALUES ($1, $2, $3, $4)
	`
	for _, translation := range translations {
		_, err := s.DB.Exec(queryTranslations, productID, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				return 0, errors.New("product translation already exists")
			}
			return 0, err
		}
	}

	return productID, nil
}

func (s *ProductPostgres) GetByID(lang string, id int) (onlinedilerv3.ProductTranslationConsignment, error) {
	query := `
        SELECT
            p.id AS id,
            p.rating AS rating,
            p.category_id AS category_id,
            pt.language_code AS language_code,
            pt.name AS name,
            pt.description AS description,
            c.expiration_date AS expiration_date,
            c.quantity AS quantity,
            c.price AS price,
            c.description AS consignment_description,
            p.image_urls AS image_urls
        FROM products p
        JOIN products_translations pt ON p.id = pt.product_id
        JOIN consignments c ON p.id = c.product_id
        WHERE p.id = $1 AND pt.language_code = $2
    `

	var result onlinedilerv3.ProductTranslationConsignment

	err := s.DB.Get(&result, query, id, lang)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, fmt.Errorf("product not found")
		}
		return result, err
	}

	return result, nil
}

func (s *ProductPostgres) Update(id int, input onlinedilerv3.ProductComplect) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	updateProductQuery := `
        UPDATE products
        SET rating = $1, category_id = $2
        WHERE id = $3
    `
	_, err = tx.Exec(updateProductQuery, input.Rating, input.CategoryID, id)
	if err != nil {
		return err
	}

	deleteTranslationsQuery := `
        DELETE FROM products_translations
        WHERE product_id = $1
    `
	_, err = tx.Exec(deleteTranslationsQuery, id)
	if err != nil {
		return err
	}

	insertTranslationQuery := `
        INSERT INTO products_translations (product_id, language_code, name, description)
        VALUES ($1, $2, $3, $4)
    `
	for _, translation := range input.Translations {
		_, err = tx.Exec(insertTranslationQuery, id, translation.LanguageCode, translation.Name, translation.Description)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ProductPostgres) Delete(product_id int) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	deleteTranslationsQuery := `
        DELETE FROM products_translations
        WHERE product_id = $1
    `
	_, err = tx.Exec(deleteTranslationsQuery, product_id)
	if err != nil {
		return err
	}

	deleteProductQuery := `
        DELETE FROM products
        WHERE id = $1
    `
	_, err = tx.Exec(deleteProductQuery, product_id)
	if err != nil {
		return err
	}

	return nil
}

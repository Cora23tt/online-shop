package repository

import (
	"github.com/cora23tt/onlinedilerv3"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user onlinedilerv3.User) (int, error)
	GetUser(username, password string) (onlinedilerv3.User, error)
}

type Products interface {
	Search(language, productName string) ([]onlinedilerv3.ProductComplex, error)
	SearchWithLimit(limit, offset int, language, productName string) ([]onlinedilerv3.ProductComplex, error)
	TopRatedWithLimit(limit, offset int, language string) ([]onlinedilerv3.ProductComplex, error)
	ByCategory(language string, categoryID int) ([]onlinedilerv3.ProductTranslationConsignment, error)
	TopRated(language string) ([]onlinedilerv3.ProductComplex, error)
	GetAll(language string) ([]onlinedilerv3.ProductTranslationConsignment, error)
	Create(product onlinedilerv3.Product, translations []onlinedilerv3.ProductTranslation) (int, error)
	GetByID(lang string, id int) (onlinedilerv3.ProductTranslationConsignment, error)
	Update(int, onlinedilerv3.ProductComplect) error
	Delete(int) error
}

type Categories interface {
	GetAll(lang string) ([]onlinedilerv3.Category, error)
	Get(lang string, id int) (onlinedilerv3.Category, error)
	Create(onlinedilerv3.Category) (int, error)
	Delete(int) error
	Update(int, onlinedilerv3.Category) error
	Search(string, string) ([]onlinedilerv3.Category, error)
}

type Users interface {
	Search(string) ([]onlinedilerv3.User, error)
	GetAll() ([]onlinedilerv3.User, error)
	GetByID(int) (onlinedilerv3.User, error)
	Delete(int) error
	Update(int, onlinedilerv3.User) error
}

type Consignments interface {
	GetAll() ([]onlinedilerv3.Consignment, error)
	GetByID(int) (onlinedilerv3.Consignment, error)
	Create(onlinedilerv3.Consignment) (int, error)
	Update(int, onlinedilerv3.Consignment) error
	Delete(int) error
}

type Discounts interface {
	GetAll() ([]onlinedilerv3.Discount, error)
	GetByID(int) (onlinedilerv3.Discount, error)
	Create(onlinedilerv3.DiscountInput) (int, error)
	Update(int, onlinedilerv3.DiscountInput) error
	Delete(int) error
}

type ClientDiscounts interface {
	GetAll() ([]onlinedilerv3.ClientDiscount, error)
	GetByID(int) (onlinedilerv3.ClientDiscount, error)
	Create(onlinedilerv3.ClientDiscount) (int, error)
	Update(int, onlinedilerv3.ClientDiscount) error
	Delete(int) error
}

type Orders interface {
	GetAll() ([]onlinedilerv3.Order, error)
	GetByID(int) (onlinedilerv3.Order, error)
	Create(onlinedilerv3.Order) (int, error)
	Update(int, onlinedilerv3.Order) error
	Delete(int) error
}

type OrderItems interface {
	GetItems(int) ([]onlinedilerv3.OrderItem, error)
	Add(int, onlinedilerv3.OrderItem) ([]onlinedilerv3.OrderItem, error)
	Update(int, onlinedilerv3.OrderItem) error
	Delete(int, int) error
}

type Repository struct {
	Authorisation
	Products
	Categories
	Users
	Consignments
	Discounts
	ClientDiscounts
	Orders
	OrderItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation:   NewAuthPostgres(db),
		Products:        NewProductPostgres(db),
		Categories:      NewCategoryPostgres(db),
		Users:           NewUserPostgres(db),
		Consignments:    NewConsignmentPostgres(db),
		Discounts:       NewDiscountPostgres(db),
		ClientDiscounts: NewClientDiscountsPostgres(db),
		Orders:          NewOrderPostgres(db),
		OrderItems:      NewOrderItemsPostgres(db),
	}
}

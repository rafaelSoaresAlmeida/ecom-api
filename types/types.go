package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProductById(id int) (*Product, error)
	GetProductsById(ids []int) ([]Product, error)
	GetProductByName(productName string)  (*Product, error)
	GetProduct() ([]*Product, error)
	CreateProduct(RegisterProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterProductPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity  int       `json:"quantity"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required,min=4,max=32"`
	LastName  string `json:"lastName" validate:"required,min=4,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
}

type LoginUserPayload struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderId"`
	ProductID int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
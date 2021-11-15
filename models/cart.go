package models

import (
	"errors"
	"fmt"
)

type Cart struct {
	ID         int
	UserID     int
	ProductIDs []int
}

type CartUpdate struct {
	Path      string
	ProductID int
}

var (
	carts      []*Cart
	nextCartID = 1
)

func GetCartById(id int) (*Cart, error) {
	for _, c := range carts {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, fmt.Errorf("No cart found with id '%v'", id)
}

func GetCartByUserId(userId int) (Cart, error) {
	for _, c := range carts {
		if c.UserID == userId {
			return *c, nil
		}
	}
	return Cart{}, fmt.Errorf("No cart found with id '%v'", userId)
}

func CreateCart(c Cart) (Cart, error) {
	if c.ID != 0 {
		return Cart{}, errors.New("New Cart must not include an ID or have and ID = 0")
	}
	if c.UserID <= 0 {
		return Cart{}, errors.New("Carts need a positive number for the UserID")
	}
	c.ID = nextCartID

	nextCartID++
	carts = append(carts, &c)
	return c, nil
}

func AddItemToCart(userID int, productID int) (Cart, error) {
	cart, err := GetCartByUserId(userID)
	if err != nil {
		var newCart Cart
		newCart.UserID = userID
		cart, err = CreateCart(newCart)
		if err != nil {
			return Cart{}, fmt.Errorf("Something when wrong while trying to create a new cart for user with id '%v'", userID)
		}
	}

	cartAddress, err := getCartAddress(cart.ID)
	if err != nil {
		return Cart{}, fmt.Errorf("Something went wrong while trying to get Cart address")
	}

	cartAddress.ProductIDs = append(cartAddress.ProductIDs, productID)
	return *cartAddress, nil
}

func RemoveItemFromCart(userID int, productID int) (Cart, error) {
	cart, err := GetCartByUserId(userID)
	if err != nil {
		return Cart{}, fmt.Errorf("No cart found with userID '%v'", userID)
	}

	cartAddress, err := getCartAddress(cart.ID)
	if err != nil {
		return Cart{}, fmt.Errorf("Something went wrong while trying to get Cart address")
	}

	for i, p := range cartAddress.ProductIDs {
		if p == productID {
			cartAddress.ProductIDs = append(cartAddress.ProductIDs[:i], cartAddress.ProductIDs[i+1:]...)
			return *cartAddress, nil
		}
	}
	return Cart{}, fmt.Errorf("Product with id '%v' not found", productID)
}

func getCartAddress(cartID int) (*Cart, error) {
	for _, c := range carts {
		if c.ID == cartID {
			return c, nil
		}
	}
	return nil, fmt.Errorf("No cart found with id '%v'", cartID)
}

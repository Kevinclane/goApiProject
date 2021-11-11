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

var (
	carts      []*Cart
	nextCartID = 1
)

func GetCartById(id int) (Cart, error) {
	for _, c := range carts {
		if c.ID == id {
			return *c, nil
		}
	}
	return Cart{}, fmt.Errorf("No cart found with id '%v'", id)
}

func CreateCart(c Cart) (Cart, error){
	if c.ID != 0 {
		return Cart{}, errors.New("New Cart must not include an ID or have and ID = 0")
	}
	if c.UserID <= 0  {
		return Cart{}, errors.New("Carts need a positive number for the UserID")
	}
	c.ID = nextCartID

	nextCartID++
	carts = append(carts, &c)
	return c, nil
}

func AddItemToCart(cartID int, productID int) (Cart, error){
	for _, c := range carts {
		if c.ID == cartID {
			cart := *c;
			cart[] = append(cart[ProductIDs], productID)
		}
	}
	return Cart{}, fmt.Errorf("Cart with id '%v' not found", cartID)
}
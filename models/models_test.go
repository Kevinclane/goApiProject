package models

import (
	"fmt"
	"strconv"
	"testing"
)

var testCart Cart

func TestCarts(t *testing.T) {

	colorRed := "\033[31m"
	colorReset := "\033[0m"
	testCart.ID = 0
	testCart.UserID = 2

	//Helper function to print values to terminal in event of a failed test
	printError := func(message, value string) {
		fmt.Println(string(colorRed))
		fmt.Printf("%v", message)

		if value != "" {
			fmt.Println()
			fmt.Printf("%v", value)
		}

		fmt.Println(string(colorReset))
	}

	//Helper function to test the inner values of both carts to be the same
	checkCartValues := func(t testing.TB, cart, expectedCart Cart, stage string) {
		if cart.ID != expectedCart.ID {
			printError("Stage: ", stage)
			printError("Mismatch IDs", "")
			printError("TestCart: ", strconv.Itoa(cart.ID))
			printError("ExpectedCart: ", strconv.Itoa(expectedCart.ID))
			t.Fail()
		}
		if cart.UserID != expectedCart.UserID {
			printError("Stage: ", stage)
			printError("Mismatch IDs", "")
			printError("TestCart: ", strconv.Itoa(cart.UserID))
			printError("ExpectedCart: ", strconv.Itoa(expectedCart.UserID))
			t.Fail()
		}

		cartProductIDLength := len(cart.ProductIDs)
		expectedCartProductIdLength := len(expectedCart.ProductIDs)

		if cartProductIDLength == expectedCartProductIdLength {
			for i := 0; i < cartProductIDLength; i++ {
				if cart.ProductIDs[i] != expectedCart.ProductIDs[i] {
					printError("Stage: ", stage)
					printError("Mismatch IDs", "")
					printError("TestCart: ", strconv.Itoa(cart.UserID))
					printError("ExpectedCart: ", strconv.Itoa(expectedCart.UserID))
					t.Fail()
				}
			}
		} else {
			printError("Stage: ", stage)
			printError("Mismatched ProductID length", "")
			t.Fail()
		}

	}

	t.Run("Testing creating a cart", func(t *testing.T) {
		createdCart, err := CreateCart(testCart)
		if err != nil {
			fmt.Errorf("Error creating cart: %v", err)
		}

		var expectedCart Cart
		expectedCart.ID = 1
		expectedCart.UserID = 2

		checkCartValues(t, createdCart, expectedCart, "Creating Cart")

	})

	t.Run("Testing getting a cart by UserID", func(t *testing.T) {
		cart, err := GetCartByUserId(2)
		if err != nil {
			fmt.Printf("Error getting cart. Error: %v", err)
			t.Fail()
		}

		var expectedCart Cart
		expectedCart.ID = 1
		expectedCart.UserID = 2

		checkCartValues(t, cart, expectedCart, "Getting Cart by UserID")

	})
}

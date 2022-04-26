package main

import (
	"customerapp/domain"
	"customerapp/mapstore"
	"fmt"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (d CustomerController) Add(c domain.Customer) {
	err := d.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (d CustomerController) Update(k string, c domain.Customer) {
	err := d.store.Update(k, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been updated")
}

func (d CustomerController) Delete(c string) {
	err := d.store.Delete(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been Deleted")
}

func (d CustomerController) GetById(ID string) {
	k, err := d.store.GetbyId(ID)
	fmt.Printf("New Customer has been Deleted", k)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}

	customer := domain.Customer{
		ID:    "cust101",
		Name:  "JPM",
		Email: "JPM@gmail.com",
	}
	controller.Add(customer)
	controller.Delete("1")
	controller.Update("2")

}

/*
2: Working with Collections
Objective: To get familiarise Go basics with Map data structure
▪ Define a package level variable of type map[string]string
▪ Create functions for making insert, update, delete and get items to
and from the map (package level variable of type map) with the following
signature:
• addItem (k,v string)
• updateItem (k,v string)
• getById (k string)
• getAll()
• deleteItem (k string)
// Declare package level variable for storing map
var data map[string]string
// init function will be automatically invoked before main function
// init function is used to initialise package level variables
func init() {
 data = make(map[string]string) // Initialise map with make
}
func addItem(k,v string) {
 // ToDo: Check if key exists
 data [k] = v
}
*/

package main

import (
	"fmt"
)

// ▪ Define a package level variable of type map[string]string
var data map[string]string

/*
signature:
• addItem (k,v string)
• updateItem (k,v string)
• getById (k string)
• getAll()
• deleteItem (k string)
*/
func addItem(k, v string) {
	data = make(map[string]string)
	data["1"] = "Surya"
	data["2"] = "Shiju"
	data["3"] = "Go-Lang"
	if _, ok := data[k]; !ok {
		data[k] = v
		for k, v := range data {

			fmt.Printf("key: %s Value: %s\n", k, v)
		}
	}
}

func updateItem(k, v string) {
	if _, ok := data[k]; ok {
		data[k] = v
		for k, v := range data {
			fmt.Printf("key: %s Value: %s\n", k, v)
		}
	}
}

func getById(k string) {
	if dataKey, ok := data[k]; ok {

		fmt.Printf("key: %s\n", dataKey)
	}
}

func getAll() {
	for k, v := range data {
		fmt.Printf("key: %s Value: %s\n", k, v)
	}
}

func deleteItem(k string) {
	delete(data, k)
	if _, ok := data[k]; !ok {
		fmt.Println("Data is deleted")
	}
}

func main() {
	k := "4"
	v := "Test"
	addItem(k, v)
	k = "4"
	v = "UpdatedTest"
	updateItem(k, v)
	getById(k)
	getAll()
	deleteItem(k)
}

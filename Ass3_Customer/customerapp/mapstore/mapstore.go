package mapstore

import (
	"customerapp/domain"
	"errors"
	"fmt"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

/*
type Customer struct {
	ID    string
	Name  string
	Email string
}
*/

//Implement interface methods of domain.CustomerStore
func (p *MapStore) Update(k string, c domain.Customer) error {
	for k := range p.store {
		if p.store[k] == c.ID {
			p.store = domain.Customer{c.Email, c.ID, c.Name}
			fmt.Println("Customer data is updated")
			return nil
		}
		return errors.New("Customer data is not updated")
	}
}
func (p *MapStore) GetById(string) (domain.Customer, error) {
	for k := range p.store {
		if p.store[k] == c.ID {
			fmt.Printf("Customer ID: %v,Name: %v, Email: %v", c.ID, c.Name, c.Email)
			return p[k], nil
		}
		return _, errors.New("Customer data is not found")
	}
}

/*
func(p *MapStore) GetAll() ([]domain.Customer, error){
			p = make([]domain.Customer{c.ID,c.Name,c.Email},len(p))
			for index, element := range p{
				return p, nil
			}
			}
*/
func (p *MapStore) Delete(k string) error {
	for k := range p.store {
		if p.store[k] == c.ID {
			delete(p, p.ID)
			fmt.Println("Customer data is deleted")
			return nil
		}
		return errors.New("Customer data is not deleted")
	}
}

func (p *MapStore) Create(c domain.Customer) error {
	for d := range p.store {
		if p.store[d] != c {
			p.store = domain.Customer{c.Email, c.ID, c.Name}
			fmt.Println("Customer data is created")
			return nil
		}
		return errors.New("Customer data is not created")
	}
}

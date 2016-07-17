package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	customerCollectionName = "customers"
)

// Customer represents the structure of our customer resource
type Customer struct {
	ID       bson.ObjectId `form:"id,omitempty" json:"id" bson:"_id"`
	Name     string        `form:"name" json:"name" bson:"name"`
	Email    string        `form:"email" json:"email" bson:"email"`
	Password string        `form:"password,omitempty" json:"-" bson:"password"`
}

func (ds *DataStore) customerCollection() *mgo.Collection {
	return ds.session.DB(Database).C(customerCollectionName)
}

// CustomerExists check if a customer exists by id
func (ds *DataStore) CustomerExists(id string) bool {
	c := ds.customerCollection()
	exists := false

	count, _ := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Count()

	if count > 0 {
		exists = true
	}

	return exists
}

// GetCustomerByID find a customer by id
func (ds *DataStore) GetCustomerByID(id string) (*Customer, error) {
	c := ds.customerCollection()
	customer := &Customer{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(customer)

	return customer, err
}

// GetCustomerByEmail find a customer by email
func (ds *DataStore) GetCustomerByEmail(email string) (*Customer, error) {
	c := ds.customerCollection()
	customer := &Customer{}
	err := c.Find(bson.M{"email": email}).One(customer)

	return customer, err
}

// NewCustomer create a new customer
func (ds *DataStore) NewCustomer(customer *Customer) error {
	c := ds.customerCollection()
	err := c.Insert(customer)

	return err
}

// UpdateCustomer updates an existing customer
func (ds *DataStore) UpdateCustomer(customer *Customer) error {
	c := ds.customerCollection()
	err := c.Update(bson.M{"_id": customer.ID}, customer)

	return err
}

// RemoveCustomer removes an existing customer
func (ds *DataStore) RemoveCustomer(id string) error {
	c := ds.customerCollection()
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}
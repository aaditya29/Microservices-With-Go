/*
Here we are creating a College structure type that holds the map of Student objects.
This type implements Get and Add method to fetch or add Student objects to the College database.
*/

package common

import (
	"fmt"
)

// Student struct representing a student.
type Student struct {
	ID                  int    //User id
	FirstName, LastName string //Username
}

// FullName returns the fullname of the student.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

// College struct representing college.
type College struct {
	database map[int]Student // private
}

// Adding two methods i.e. here we add a student to the college (procedure).
func (c *College) Add(payload Student, reply *Student) error { // request or response payload is also called message. Undestand codec if confusion

	// check if student already exists in the database
	if _, ok := c.database[payload.ID]; ok { // ok will recieve a bool that is set to be true if student's value is already present in the map
		return fmt.Errorf("student with id '%d' already exists", payload.ID)
	}

	// adding student `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Get methods returns a student with specific id (procedure).
func (c *College) Get(payload int, reply *Student) error {

	// get student with id `p` from the database
	result, ok := c.database[payload]

	// check if student exists in the database
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", payload)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

// NewCollege function returns a new instance of College (pointer).
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}

package model

type Candidate struct {
	Date       string
	Name       string
	Address    string
	Address2   string
	City       string
	State      string
	Zipcode    string
	Telephone  string
	Mobile     string
	Amount     string
	Processor  string
	ImportDate string
}

/*Equals method checks if the given candidate and the other one are equal.
Assumption: the same candidate but on different date should not be treated as duplicate.
Import date is not considered in this method. The same candidates but with different import date
will be treated as duplicates */
func (c *Candidate) Equals(o Candidate) bool {
	if c.Date != o.Date {
		return false
	}
	if c.Name != o.Name {
		return false
	}
	if c.Address != o.Address {
		return false
	}
	if c.Address2 != o.Address2 {
		return false
	}
	if c.City != o.City {
		return false
	}
	if c.State != o.State {
		return false
	}
	if c.Zipcode != o.Zipcode {
		return false
	}
	if c.Telephone != o.Telephone {
		return false
	}
	if c.Mobile != o.Mobile {
		return false
	}
	if c.Amount != o.Amount {
		return false
	}
	if c.Processor != o.Processor {
		return false
	}
	return true
}

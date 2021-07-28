package model

type Candidate struct {
	Date       string `json:"date"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zipCode"`
	Telephone  string `json:"telephone"`
	Mobile     string `json:"mobile"`
	Amount     string `json:"amount"`
	Processor  string `json:"processor"`
	ImportDate string `json:"importDate"`
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
	if c.ZipCode != o.ZipCode {
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

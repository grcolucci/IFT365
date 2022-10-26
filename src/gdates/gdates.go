package gdates

import (
	"errors"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}

	d.year = year
	return nil
}

func (d *Date) Year() int {

	return d.year
}

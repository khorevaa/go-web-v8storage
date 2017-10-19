package models

type Hotel struct {
	HotelId          int
	Name, Address    string
	City, State, Zip string
	Country          string
	Price            int
}

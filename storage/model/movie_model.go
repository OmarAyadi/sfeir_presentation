package model

type Director struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type Movie struct {
	ID       string    `json:"id" bson:"_id"`
	Title    string    `json:"title" bson:"title"`
	Year     int       `json:"year" bson:"year"`
	Director *Director `json:"director" bson:"director"`
}

type MoviesPaginated struct {
	Movies     *[]Movie `json:"movies"`
	TotalCount int64    `json:"total_count"`
}

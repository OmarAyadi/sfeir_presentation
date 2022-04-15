package dto

type DirectorDto struct {
	FirstName string `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type MovieDto struct {
	Title       string       `json:"title" bson:"title" binding:"required"`
	Year        int          `json:"year" bson:"year" binding:"required,gte=1940,lte=2100"`
	DirectorDto *DirectorDto `json:"director" bson:"director"`
}

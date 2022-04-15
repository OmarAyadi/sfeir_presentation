package global

const (
	ApiV1        = "/api/v1"
	EmptyString  = ""
	MongoID      = "_id"
	MovieIDParam = "movie_id"

	MongoTestURI = "mongodb://localhost:27017"
	MongoTestDB  = "unit_tests"

	Page         = "page"
	DefaultPage  = int64(0)
	Limit        = "limit"
	DefaultLimit = int64(20)
)

var (
	InvalidObjectIDsFormat = []string{
		"",
		"1",
		"12",
		"123",
		"1234",
		"12345",
		"123456",
		"1234567",
		"12345678",
		"123456789",
		"1234567890",
		"12345678901",
		"123456789012",
		"1234567890123",
		"12345678901234",
		"123456789012345",
		"1234567890123456",
		"12345678901234567",
		"123456789012345678",
		"1234567890123456789",
		"12345678901234567890",
		"123456789012345678901",
		"1234567890123456789012",
		"12345678901234567890123",
		"1234567890123456789012345",
	}
)

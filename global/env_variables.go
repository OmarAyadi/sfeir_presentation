package global

import (
	"os"
)

var MongoDB string
var MongoHost string
var GinMode string

func LazyEnvVariableInit() {
	MongoDB = os.Getenv("MONGO_DB")
	MongoHost = os.Getenv("MONGO_HOST")
	GinMode = os.Getenv("GIN_MODE")
}

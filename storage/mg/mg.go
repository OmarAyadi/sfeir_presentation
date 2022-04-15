package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sfeir/global"
	"testing"
	"time"
)

func OpenMongoConnectionAndGetDB(uri string, databaseName string) (db *mongo.Database, err error) {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		return
	}

	db = mongoClient.Database(databaseName)

	return
}

func OpenTestDataBaseHelper(t *testing.T) *mongo.Database {
	t.Helper()

	t.Log("Open test database....")
	db, err := OpenMongoConnectionAndGetDB(global.MongoTestURI, global.MongoTestDB)
	if err != nil {
		t.Fatalf("Cannot open test database: %v", err)
	}

	t.Cleanup(func() {
		t.Log("Closing test database...")
		err = db.Drop(nil)
		if err != nil {
			t.Fatalf("error while closing db: %v", err)
		}
	})

	return db
}

func GetCollection(db *mongo.Database, collectionName global.CollectionName) *mongo.Collection {
	return db.Collection(global.ToStr(collectionName))
}

type StorageImpl struct {
	MoviesStorageImpl
}

func NewStorageImpl(db *mongo.Database) *StorageImpl {
	return &StorageImpl{
		MoviesStorageImpl: NewMoviesStorageImpl(db),
	}
}

type MoviesStorageImpl struct {
	MoviesCol *mongo.Collection
}

func NewMoviesStorageImpl(db *mongo.Database) MoviesStorageImpl {
	return MoviesStorageImpl{
		MoviesCol: GetCollection(db, global.CNMovies),
	}
}

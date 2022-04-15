package repository

import (
	"Desktop/shopi/assignment/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Add(item model.OrderV2) error
	Filter(filterv2 model.OrderFilterModel2, query interface{}) ([]model.OrderV2, error)
}

type repository struct {
	db          *mongo.Database
	mongoClient *mongo.Client
	Collection  *mongo.Collection
}

var _ Repository = repository{}

func NewRepository(db *mongo.Database, mongoClient *mongo.Client, Collection *mongo.Collection) Repository {
	return repository{db: db, mongoClient: mongoClient, Collection: Collection}
}

func (r repository) Add(item model.OrderV2) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := r.Collection.InsertOne(ctx, item)

	if err != nil {
		return err
	}

	return nil
}

func (r repository) Filter(filterv2 model.OrderFilterModel2, query interface{}) (result []model.OrderV2, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	opts := options.Find()

	opts.SetSort(bson.D{{filterv2.SortBy, 1}})

	skipNumber := (filterv2.PageNumber - 1) * (filterv2.PageSize)

	opts.SetSkip(int64(skipNumber))
	opts.SetLimit(int64(filterv2.PageSize))

	// regexPattern := ".*" + filterv2.SearchText + ".*"

	// filter3 := bson.M{
	// 	"createdon": bson.M{
	// 		"$gt": filterv2.StartDate,
	// 		"$lt": filterv2.EndDate,
	// 	},
	// 	"$or": bson.A{
	// 		bson.M{"storename": bson.M{"$regex": regexPattern, "$options": "im"}},
	// 		bson.M{"customername": bson.M{"$regex": regexPattern, "$options": "im"}},
	// 	}, "status": bson.M{"$in": filterv2.Statuses}}

	cursor, err := r.Collection.Find(ctx, query, opts)

	if err != nil {
		log.Fatal(err)
	}

	var elements []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	bsonBytes, _ := bson.Marshal(elements)
	bson.Unmarshal(bsonBytes, &result)

	return result, err

}

package service

import (
	"Desktop/shopi/assignment/model"
	"Desktop/shopi/assignment/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Add(item model.Order) error
	Filter(model.OrderFilterModel) ([]model.OrderV2, error)
}

type service struct {
	repo repository.Repository
}

var _ Service = service{}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}

func (s service) Add(item model.Order) error {

	createdv2, err := time.Parse("2006-01-02 15:04", item.CreatedOn)
	if err != nil {
		return err
	}
	itemv2 := model.OrderV2{Id: item.Id, BrandId: item.BrandId, Price: item.Price, StoreName: item.StoreName, CustomerName: item.CustomerName, CreatedOn: createdv2, Status: item.Status}

	return s.repo.Add(itemv2)
}

func (s service) Filter(filter model.OrderFilterModel) ([]model.OrderV2, error) {
	startData, err := time.Parse("2006-01-02 15:04", filter.StartDate)

	if err != nil {
		return nil, err
	}

	endData, err := time.Parse("2006-01-02 15:04", filter.EndDate)
	if err != nil {
		return nil, err

	}

	filterv2 := model.OrderFilterModel2{PageSize: filter.PageSize, PageNumber: filter.PageNumber, SearchText: filter.SearchText, StartDate: startData, EndDate: endData, Statuses: filter.Statuses, SortBy: filter.SortBy}

	regexPattern := ".*" + filterv2.SearchText + ".*"

	query := bson.M{
		"createdon": bson.M{
			"$gt": filterv2.StartDate,
			"$lt": filterv2.EndDate,
		},
		"$or": bson.A{
			bson.M{"storename": bson.M{"$regex": regexPattern, "$options": "im"}},
			bson.M{"customername": bson.M{"$regex": regexPattern, "$options": "im"}},
		}, "status": bson.M{"$in": filterv2.Statuses}}

	return s.repo.Filter(filterv2, query)
}

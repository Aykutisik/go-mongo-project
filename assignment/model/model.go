package model

import "time"

type OrderStatus int

const (
	Created    OrderStatus = 10
	InProgress OrderStatus = 20
	Failed     OrderStatus = 30
	Completed  OrderStatus = 40
	Canceled   OrderStatus = 50
	Returned   OrderStatus = 60
)

type Order struct {
	Id           int         `json:"id"`
	BrandId      int         `json:"brandid"`
	Price        int         `json:"price"`
	StoreName    string      `json:"storename"`
	CustomerName string      `json:"customername"`
	CreatedOn    string      `json:"createdon"`
	Status       OrderStatus `json:"status"`
}

type OrderV2 struct {
	Id           int         `json:"id"`
	BrandId      int         `json:"brandid"`
	Price        int         `json:"price"`
	StoreName    string      `json:"storename"`
	CustomerName string      `json:"customername"`
	CreatedOn    time.Time   `json:"createdon"`
	Status       OrderStatus `json:"status"`
}

type OrderFilterModel struct {
	PageSize   int    `json:"pagesize"`
	PageNumber int    `json:"pagenumber"`
	SearchText string `json:"searchtext"`
	StartDate  string `json:"startdate"`
	EndDate    string `json:"enddate"`
	Statuses   []int  `json:"statuses"`
	SortBy     string `json:"sortby"`
}

type OrderFilterModel2 struct {
	PageSize   int       `json:"pagesize"`
	PageNumber int       `json:"pagenumber"`
	SearchText string    `json:"searchtext"`
	StartDate  time.Time `json:"startdate"`
	EndDate    time.Time `json:"enddate"`
	Statuses   []int     `json:"statuses"`
	SortBy     string    `json:"sortby"`
}

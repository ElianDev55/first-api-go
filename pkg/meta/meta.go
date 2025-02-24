package meta

import (
	"os"
	"strconv"
)


type Meta struct {

	TotalCount int `json:"total_count"`
	Page int `json:"page"`
	PerPage int `json:"perPage"`

}

func New(page, perPage, total int)  (*Meta, error){
	
	if perPage <= 0 {
		var err error 
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page <0 {
		page = 1
	}
	
	return &Meta{
		TotalCount: total,
		Page: page,
		PerPage: perPage,
	}, nil
}

func (m *Meta) Offset() int {
	return (m.Page - 1 ) * m.PerPage	
}

func (m *Meta) Limit() int {
	return m.PerPage	
}


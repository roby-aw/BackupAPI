package utils

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GeneratePagination(query url.Values) Pagination {
	limit := 25
	page := 1
	sort := "created_at desc"
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

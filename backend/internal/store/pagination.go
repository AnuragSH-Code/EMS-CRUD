package store

import (
	"net/http"
	"strconv"
)

type PaginatedQuery struct {
	Limit  int
	Offset int
}

func ParsePagination(r *http.Request) PaginatedQuery {
	qs := r.URL.Query()

	limit, _ := strconv.Atoi(qs.Get("limit"))
	offset, _ := strconv.Atoi(qs.Get("offset"))

	if limit <= 0 {
		limit = 10
	}

	return PaginatedQuery{
		Limit:  limit,
		Offset: offset,
	}
}

package http

import (
	"net/url"
	"strconv"
)

// Paginate parses query params used for pagination.
func Paginate(query url.Values) (int64, int64, error) {
	pageString := query.Get("page")
	limitString := query.Get("limit")

	if pageString == "" {
		pageString = "1"
	}

	if limitString == "" {
		limitString = "10"
	}

	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return page, limit, nil
}

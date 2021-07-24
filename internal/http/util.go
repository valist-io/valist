package http

import (
	"math/big"
	"net/url"
)

// Paginate parses query params used for pagination.
func Paginate(query url.Values) (*big.Int, *big.Int, error) {
	var limit big.Int
	if _, ok := limit.SetString(query.Get("limit"), 10); !ok {
		limit.SetInt64(10)
	}

	var page big.Int
	if _, ok := page.SetString(query.Get("page"), 10); !ok {
		page.SetInt64(1)
	}

	return &page, &limit, nil
}

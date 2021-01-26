package httputils

import (
	"strconv"

	"github.com/eavillacis/velociraptor/pkg/errors"
	"github.com/gin-gonic/gin"
)

// PaginationResponse ...
type PaginationResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// ExtractPaginationParams ...
func ExtractPaginationParams(c *gin.Context) (offset, limit int, err error) {
	offset, err = strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return 0, 0, errors.Wrap(err, "error parsing offset")
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		return 0, 0, errors.Wrap(err, "error parsing limit")
	}

	// Never allow more than 150 results in a single query
	if limit > 150 {
		return 0, 0, errors.New("limit is too big")
	}

	return
}

// ExtractLimitParam ...
func ExtractLimitParam(c *gin.Context) (int, error) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		return 0, errors.Wrap(err, "error parsing limit")
	}

	// Never allow more than 100 results in a single query
	if limit > 100 {
		return 0, errors.New("limit is too big")
	}

	return limit, nil
}

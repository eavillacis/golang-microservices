package httputils

import "strings"

// ParseSortBy ...
func ParseSortBy(sortBy string) (sortKey, sortDirection string) {
	if strings.HasPrefix(sortBy, "-") {
		sortDirection = "DESC"
		sortKey = sortBy[1:]
	} else {
		sortDirection = "ASC"
		sortKey = sortBy
	}

	return
}

package util

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Paging struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
	Search []map[string]string
}

func NewPaging(query url.Values, searcheables []string) *Paging {
	queryPage := query.Get("page")
	queryPerPage := query.Get("perPage")
	sort := query.Get("sort")

	page, _ := strconv.Atoi(queryPage)
	perPage, _ := strconv.Atoi(queryPerPage)

	// set page
	if page < 1 {
		page = 1
	}

	// set limit
	if perPage < 1 {
		perPage, _ = strconv.Atoi(os.Getenv("QUERY_LIMIT_DEFAULT"))
	}

	// set sort
	if sort == "" {
		sort = os.Getenv("QUERY_SORT_DEFAULT")
	} else {
		sorts := strings.Split(sort, " ")
		sortField := strings.Trim(sorts[0], " ")
		sortValue := strings.ToLower(strings.Trim(sorts[1], " "))

		if contains(searcheables, sortField) && (sortValue == "asc" || sortValue == "desc") {
			sortFieldSnakeCase := CamelCaseToSnakeCase(sortField)
			sort = fmt.Sprintf("%s %s", sortFieldSnakeCase, sortValue)
		} else {
			sort = os.Getenv("QUERY_SORT_DEFAULT")
		}
	}

	// set searches
	var searches []map[string]string
	for k, v := range query {
		if k != "page" && k != "perPage" && k != "sort" {
			if contains(searcheables, k) {
				searchFieldSnakeCase := CamelCaseToSnakeCase(k)
				searches = append(searches, map[string]string{searchFieldSnakeCase: v[0]})
			}
		}
	}

	return &Paging{
		Page:   page,
		Limit:  perPage,
		Offset: (page - 1) * perPage,
		Sort:   sort,
		Search: searches,
	}
}

func contains(list []string, item string) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}

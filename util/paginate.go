package util

type Pagination struct {
  CurrentPage  int         `json:"currentPage"`
  PreviousPage interface{} `json:"previousPage"`
  NextPage     interface{} `json:"nextPage"`
  Total        int         `json:"total"`
  PerPage      int         `json:"perPage"`
  Data         interface{} `json:"data"`
}

func Paginate(paging *Paging, data interface{}, total int) *Pagination {

  // calculate previous page
  previousPage := calculatePreviousPage(paging.Page)

  var resPreviousPage interface{}
  if previousPage == -1 {
    resPreviousPage = nil
  } else {
    resPreviousPage = previousPage
  }

  // calculate next page
  nextPage := calculateNextPage(paging.Page, total, paging.Limit)

  var resNextPage interface{}
  if nextPage == -1 {
    resNextPage = nil
  } else {
    resNextPage = nextPage
  }

  return &Pagination{
    CurrentPage:  paging.Page,
    PreviousPage: resPreviousPage,
    NextPage:     resNextPage,
    Total:        total,
    PerPage:      paging.Limit,
    Data:         data,
  }
}

func calculatePreviousPage(page int) int {
  previousPage := page - 1

  if previousPage < 1 {
    return -1
  }

  return previousPage
}

func calculateNextPage(page int, total int, limit int) int {
  leftover := float64(total) / float64(page*limit)

  if leftover <= 1 {
    return -1
  }

  return page + 1
}

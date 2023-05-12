package util

import (
  "fmt"
  "net/url"
  "testing"
)

func TestPaginate(t *testing.T) {
  t.Run("beginning of the page", func(t *testing.T) {
    page := 1
    perPage := 10
    total := 100

    // get query string
    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&perPage=%d`, page, perPage))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    // get searcheables
    searcheables := getSearcheables()

    // get paging
    paging := NewPaging(q, searcheables)

    data := []*testEntity{
      {
        Name: "aaa111",
      },
    }

    paginate := Paginate(paging, data, total)

    if paginate.CurrentPage != page {
      t.Errorf(`got: %d, want: %d`, paginate.CurrentPage, page)
    }
    if paginate.PreviousPage != nil {
      t.Errorf(`got: %v, want: %v`, paginate.PreviousPage, nil)
    }
    if paginate.NextPage != page+1 {
      t.Errorf(`got: %d, want: %d`, paginate.NextPage, page+1)
    }
    if paginate.PerPage != perPage {
      t.Errorf(`got: %d, want: %d`, paginate.PerPage, perPage)
    }
    if paginate.Total != total {
      t.Errorf(`got: %d, want: %d`, paginate.Total, total)
    }
  })

  t.Run("middle page", func(t *testing.T) {
    page := 2
    perPage := 10
    total := 100

    // get query string
    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&perPage=%d`, page, perPage))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    // get searcheables
    searcheables := getSearcheables()

    // get paging
    paging := NewPaging(q, searcheables)

    data := []*testEntity{
      {
        Name: "aaa111",
      },
    }

    paginate := Paginate(paging, data, total)

    if paginate.CurrentPage != page {
      t.Errorf(`got: %d, want: %d`, paginate.CurrentPage, page)
    }
    if paginate.PreviousPage != page-1 {
      t.Errorf(`got: %v, want: %v`, paginate.PreviousPage, page-1)
    }
    if paginate.NextPage != page+1 {
      t.Errorf(`got: %d, want: %d`, paginate.NextPage, page+1)
    }
    if paginate.PerPage != perPage {
      t.Errorf(`got: %d, want: %d`, paginate.PerPage, perPage)
    }
    if paginate.Total != total {
      t.Errorf(`got: %d, want: %d`, paginate.Total, total)
    }
  })

  t.Run("end of page", func(t *testing.T) {
    page := 10
    perPage := 10
    total := 100

    // get query string
    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&perPage=%d`, page, perPage))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    // get searcheables
    searcheables := getSearcheables()

    // get paging
    paging := NewPaging(q, searcheables)

    data := []*testEntity{
      {
        Name: "aaa111",
      },
    }

    paginate := Paginate(paging, data, total)

    if paginate.CurrentPage != page {
      t.Errorf(`got: %d, want: %d`, paginate.CurrentPage, page)
    }
    if paginate.PreviousPage != page-1 {
      t.Errorf(`got: %d, want: %d`, paginate.PreviousPage, page-1)
    }
    if paginate.NextPage != nil {
      t.Errorf(`got: %v, want: %v`, paginate.NextPage, nil)
    }
    if paginate.PerPage != perPage {
      t.Errorf(`got: %d, want: %d`, paginate.PerPage, perPage)
    }
    if paginate.Total != total {
      t.Errorf(`got: %d, want: %d`, paginate.Total, total)
    }
  })
}

func TestCalculatePreviousPage(t *testing.T) {
  tests := []struct {
    page   int
    expect int
  }{
    {page: 1, expect: -1},
    {page: 2, expect: 1},
    {page: 100, expect: 99},
  }

  for i, tt := range tests {
    res := calculatePreviousPage(tt.page)

    if res != tt.expect {
      t.Errorf(`test #%d - page: %d, res: %d, expected: %d`, i+1, tt.page, res, tt.expect)
    }
  }
}

func TestCalculateNextPage(t *testing.T) {
  tests := []struct {
    page   int
    total  int
    limit  int
    expect int
  }{
    {page: 1, total: 100, limit: 10, expect: 2},
    {page: 1, total: 100, limit: 100, expect: -1},
    {page: 10, total: 100, limit: 10, expect: -1},
    {page: 10, total: 101, limit: 10, expect: 11},
  }

  for i, tt := range tests {
    t.Run("", func(t *testing.T) {
      res := calculateNextPage(tt.page, tt.total, tt.limit)

      if res != tt.expect {
        t.Errorf(`test #%d - page: %d, total: %d, limit: %d, res: %d, expected: %d`, i+1, tt.page, tt.total, tt.limit, res, tt.expect)
      }
    })
  }
}

func getSearcheables() []string {
  return []string{"id", "name"}
}

type testEntity struct {
  Name string
}

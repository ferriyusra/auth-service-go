package util

import (
  "fmt"
  "net/url"
  "os"
  "strconv"
  "testing"
)

func TestMain(m *testing.M) {
  os.Setenv("APP_ENV", "test")
  os.Setenv("QUERY_SORT_DEFAULT", "created_at desc")
  os.Setenv("QUERY_LIMIT_DEFAULT", "10")
  os.Exit(m.Run())
}

func TestNewPaging(t *testing.T) {
  t.Run("no page parameter", func(t *testing.T) {
    perPage := 20

    u, err := url.Parse(fmt.Sprintf(`https://example.org?perPage=%d`, perPage))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    searcheables := getSearcheables()

    paging := NewPaging(q, searcheables)

    if paging.Page != 1 {
      t.Errorf(`got: %d, want: %d`, paging.Page, 1)
    }
    if paging.Limit != perPage {
      t.Errorf(`got: %d, want: %d`, paging.Limit, perPage)
    }
    if paging.Offset != 0 {
      t.Errorf(`got: %d, want: %d`, paging.Offset, 0)
    }
  })

  t.Run("no perPage parameter", func(t *testing.T) {
    page := 2

    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d`, page))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    searcheables := getSearcheables()

    paging := NewPaging(q, searcheables)

    limit, _ := strconv.Atoi(os.Getenv("QUERY_LIMIT_DEFAULT"))
    offset := (page - 1) * limit

    if paging.Page != page {
      t.Errorf(`got: %d, want: %d`, paging.Page, page)
    }
    if paging.Limit != limit {
      t.Errorf(`got: %d, want: %d`, paging.Limit, limit)
    }
    if paging.Offset != offset {
      t.Errorf(`got: %d, want: %d`, paging.Offset, offset)
    }
  })

  t.Run("with sort parameter", func(t *testing.T) {
    page := 2
    sort := "name asc"

    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&sort=%s`, page, sort))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    searcheables := getSearcheables()

    paging := NewPaging(q, searcheables)

    limit, _ := strconv.Atoi(os.Getenv("QUERY_LIMIT_DEFAULT"))
    offset := (page - 1) * limit

    if paging.Page != page {
      t.Errorf(`got: %d, want: %d`, paging.Page, page)
    }
    if paging.Limit != limit {
      t.Errorf(`got: %d, want: %d`, paging.Limit, limit)
    }
    if paging.Offset != offset {
      t.Errorf(`got: %d, want: %d`, paging.Offset, offset)
    }
    if paging.Sort != sort {
      t.Errorf(`got: %s, want: %s`, paging.Sort, sort)
    }
  })

  t.Run("sort with unsearcheable parameter", func(t *testing.T) {
    page := 2
    sort := "dontexist asc"

    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&sort=%s`, page, sort))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    searcheables := getSearcheables()

    paging := NewPaging(q, searcheables)

    limit, _ := strconv.Atoi(os.Getenv("QUERY_LIMIT_DEFAULT"))
    offset := (page - 1) * limit

    if paging.Page != page {
      t.Errorf(`got: %d, want: %d`, paging.Page, page)
    }
    if paging.Limit != limit {
      t.Errorf(`got: %d, want: %d`, paging.Limit, limit)
    }
    if paging.Offset != offset {
      t.Errorf(`got: %d, want: %d`, paging.Offset, offset)
    }
    if paging.Sort != os.Getenv("QUERY_SORT_DEFAULT") {
      t.Errorf(`got: %s, want: %s`, paging.Sort, os.Getenv("QUERY_SORT_DEFAULT"))
    }
  })

  t.Run("with search parameter", func(t *testing.T) {
    page := 2
    search := "clark kent"

    u, err := url.Parse(fmt.Sprintf(`https://example.org?page=%d&name=%s`, page, search))
    if err != nil {
      t.Fatal(err)
    }
    q := u.Query()

    searcheables := getSearcheables()

    paging := NewPaging(q, searcheables)

    limit, _ := strconv.Atoi(os.Getenv("QUERY_LIMIT_DEFAULT"))
    offset := (page - 1) * limit

    if paging.Page != page {
      t.Errorf(`got: %d, want: %d`, paging.Page, page)
    }
    if paging.Limit != limit {
      t.Errorf(`got: %d, want: %d`, paging.Limit, limit)
    }
    if paging.Offset != offset {
      t.Errorf(`got: %d, want: %d`, paging.Offset, offset)
    }
    if paging.Search[0]["name"] != search {
      t.Errorf(`got: %d, want: %s`, paging.Offset, search)
    }
  })
}

func TestContains(t *testing.T) {
  list := []string{"aaa", "bbb", "ccc"}

  tests := []struct {
    item string
    want bool
  }{
    {item: "aaa", want: true},
    {item: "zzz", want: false},
  }

  for i, tt := range tests {
    res := contains(list, tt.item)

    if res != tt.want {
      t.Errorf(`test #%d - got: %t, want: %t`, i+1, res, tt.want)
    }
  }
}

package util

import (
  "testing"
)

func TestNullString(t *testing.T) {
  tests := []struct {
    text  string
    want  string
    valid bool
  }{
    {text: "爱情", want: "爱情", valid: true},
    {text: "", want: "", valid: false},
  }

  for i, tt := range tests {
    res := NullString(tt.text)

    if res.String != tt.want || res.Valid != tt.valid {
      t.Errorf(`test #%d: got: %s, want: %s`, i+1, res.String, tt.want)
    }
  }
}

func TestNullInt64(t *testing.T) {
  tests := []struct {
    text  string
    want  int64
    valid bool
  }{
    {text: "123", want: 123, valid: true},
    {text: "-123", want: -123, valid: true},
    {text: "", want: 0, valid: false},
    {text: "abc", want: 0, valid: false},
  }

  for i, tt := range tests {
    res := NullInt64(tt.text)

    if res.Int64 != tt.want || res.Valid != tt.valid {
      t.Errorf(`test #%d: got: %d, want: %d`, i+1, res.Int64, tt.want)
    }
  }
}

func TestNullInt32(t *testing.T) {
  tests := []struct {
    text  string
    want  int32
    valid bool
  }{
    {text: "123", want: 123, valid: true},
    {text: "-123", want: -123, valid: true},
    {text: "", want: 0, valid: false},
    {text: "abc", want: 0, valid: false},
  }

  for i, tt := range tests {
    res := NullInt32(tt.text)

    if res.Int32 != tt.want || res.Valid != tt.valid {
      t.Errorf(`test #%d: got: %d, want: %d`, i+1, res.Int32, tt.want)
    }
  }
}

func TestNullBool(t *testing.T) {
  tests := []struct {
    text  string
    want  bool
    valid bool
  }{
    {text: "true", want: true, valid: true},
    {text: "false", want: false, valid: true},
    {text: "", want: false, valid: false},
    {text: "abc", want: false, valid: false},
  }

  for i, tt := range tests {
    res := NullBool(tt.text)

    if res.Bool != tt.want || res.Valid != tt.valid {
      t.Errorf(`test #%d: got: %t, want: %t`, i+1, res.Bool, tt.want)
    }
  }
}

func TestNullUuid(t *testing.T) {
  t.Run("Valid uuid", func(t *testing.T) {
    v4 := "210b9652-e95b-4872-b246-dee08da8d5ef" // version 4
    uuidV4 := NullUuid(v4)

    if val := uuidV4.UUID.String(); val != v4 {
      t.Errorf("got: %s want: %s\b", val, v4)
    }

    if uuidV4.Valid != true {
      t.Errorf("got: %t want: %t\b", uuidV4.Valid, true)
    }
  })

  t.Run("Empty string", func(t *testing.T) {
    uuidV4 := NullUuid("")

    if val := uuidV4.UUID.String(); val != "00000000-0000-0000-0000-000000000000" {
      t.Errorf("got: %s want: %s\b", val, "00000000-0000-0000-0000-000000000000")
    }

    if uuidV4.Valid != false {
      t.Errorf("got: %t want: %t\b", uuidV4.Valid, false)
    }
  })

  t.Run("Invalid string", func(t *testing.T) {
    uuidV4 := NullUuid("abc")

    if val := uuidV4.UUID.String(); val != "00000000-0000-0000-0000-000000000000" {
      t.Errorf("got: %s want: %s\b", val, "00000000-0000-0000-0000-000000000000")
    }

    if uuidV4.Valid != false {
      t.Errorf("got: %t want: %t\b", uuidV4.Valid, false)
    }
  })
}

func TestGetNullableString(t *testing.T) {
  t.Run("Valid", func(t *testing.T) {
    originalString := "爱情"
    ns := NullString(originalString)
    res := GetNullableString(ns)

    if res != originalString {
      t.Errorf("got: %s want: %s", res, originalString)
    }
  })

  t.Run("Empty string", func(t *testing.T) {
    ns := NullString("")
    res := GetNullableString(ns)

    if res != nil {
      t.Errorf("res should be nil")
    }
  })
}

func TestGetNullableInt64(t *testing.T) {
  t.Run("Valid", func(t *testing.T) {
    originalString := "123"
    want := int64(123)
    ns := NullInt64(originalString)
    res := GetNullableInt64(ns)

    if res != want {
      t.Errorf("got: %d want: %d", res, want)
    }
  })

  t.Run("Empty string", func(t *testing.T) {
    ns := NullInt64("")
    res := GetNullableInt64(ns)

    if res != nil {
      t.Errorf("res should be nil")
    }
  })
}

func TestGetNullableInt32(t *testing.T) {
  t.Run("Valid", func(t *testing.T) {
    originalString := "123"
    want := int32(123)
    ns := NullInt32(originalString)
    res := GetNullableInt32(ns)

    if res != want {
      t.Errorf("got: %d want: %d", res, want)
    }
  })

  t.Run("Empty string", func(t *testing.T) {
    ns := NullInt32("")
    res := GetNullableInt32(ns)

    if res != nil {
      t.Errorf("res should be nil")
    }
  })
}

func TestGetNullableBool(t *testing.T) {
  t.Run("Valid true", func(t *testing.T) {
    originalString := "true"
    want := true
    ns := NullBool(originalString)
    res := GetNullableBool(ns)

    if res != want {
      t.Errorf("got: %t want: %t", res, want)
    }
  })

  t.Run("Valid false", func(t *testing.T) {
    originalString := "false"
    want := false
    ns := NullBool(originalString)
    res := GetNullableBool(ns)

    if res != want {
      t.Errorf("got: %t want: %t", res, want)
    }
  })

  t.Run("Invalid", func(t *testing.T) {
    ns := NullBool("")
    res := GetNullableBool(ns)

    if res != nil {
      t.Errorf("res should be nil")
    }
  })
}

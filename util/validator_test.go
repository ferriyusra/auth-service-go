package util

import (
  "fmt"
  "testing"
)

func TestValidateRequest(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    payload := &testRequest{
      Name:  "浜辺美波",
      Level: 3,
    }

    if errors := ValidateRequest(payload); len(errors) > 0 {
      fmt.Printf("%+v\n", errors)
      t.Errorf("len(errors) is %d, should be 0", len(errors))
    }
  })

  t.Run("error", func(t *testing.T) {
    payload := &testRequest{
      Name:  "浜辺美波",
      Level: 300,
    }

    if errors := ValidateRequest(payload); len(errors) == 0 {
      fmt.Printf("%+v\n", errors)
      t.Errorf("len(errors) should be more than 0")
    }
  })
}

type testRequest struct {
  Name     string `json:"name" validate:"required"`
  Level    int    `json:"level" validate:"required,min=1,max=3"`
  Password int    `json:"-"`
}

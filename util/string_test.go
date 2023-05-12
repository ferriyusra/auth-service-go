package util

import (
	"testing"
)

func TestCamelCaseToSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Name", want: "name"},
		{name: "CreatedAt", want: "created_at"},
	}

	for i, tt := range tests {
		res := CamelCaseToSnakeCase(tt.name)

		if res != tt.want {
			t.Errorf(`test #%d - got: %s, want: %s`, i+1, res, tt.want)
		}
	}
}

func TestRandString(t *testing.T) {
	tests := []struct {
		length int
		want   int
	}{
		{length: 0, want: 0},
		{length: 10, want: 10},
		{length: -1, want: 0},
	}

	for i, tt := range tests {
		res := RandString(tt.length)

		if len(res) != tt.want {
			t.Errorf(`test #%d - got: %d, want: %d`, i+1, len(res), tt.want)
		}
	}
}

func Test_ContainsString(t *testing.T) {
	tests := []struct {
		list []string
		item string
		want bool
	}{
		{list: []string{"a", "b", "c"}, item: "a", want: true},
		{list: []string{"a", "b", "c"}, item: "d", want: false},
	}

	for i, tt := range tests {
		res := contains(tt.list, tt.item)
		if res != tt.want {
			t.Errorf(`test #%d - got: %v, want: %v`, i+1, res, tt.want)
		}
	}
}

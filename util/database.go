package util

import (
	"database/sql"
	"strconv"

	"github.com/google/uuid"
)

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NullInt64(s string) sql.NullInt64 {
	if len(s) == 0 {
		return sql.NullInt64{}
	}

	intValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: intValue,
		Valid: true,
	}
}

func NullInt32(s string) sql.NullInt32 {
	if len(s) == 0 {
		return sql.NullInt32{}
	}

	intValue, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: int32(intValue),
		Valid: true,
	}
}

func NullBool(s string) sql.NullBool {
	if len(s) == 0 {
		return sql.NullBool{}
	}

	boolValue, err := strconv.ParseBool(s)
	if err != nil {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  boolValue,
		Valid: true,
	}
}

func NullUuid(s string) uuid.NullUUID {
	if len(s) == 0 {
		return uuid.NullUUID{}
	}

	uuidValue, err := uuid.Parse(s)
	if err != nil {
		return uuid.NullUUID{}
	}

	return uuid.NullUUID{
		UUID:  uuidValue,
		Valid: true,
	}
}

func GetNullableString(ns sql.NullString) interface{} {
	if ns.Valid == true {
		return ns.String
	}

	return nil
}

func GetNullableInt64(ns sql.NullInt64) interface{} {
	if ns.Valid == true {
		return ns.Int64
	}

	return nil
}

func GetNullableInt32(ns sql.NullInt32) interface{} {
	if ns.Valid == true {
		return ns.Int32
	}

	return nil
}

func GetNullableBool(ns sql.NullBool) interface{} {
	if ns.Valid == true {
		return ns.Bool
	}

	return nil
}

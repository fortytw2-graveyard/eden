package types

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
)

// IntArray allows for postgres to deal with integer arrays
type IntArray []int

// Value lets intArrays be inserted into PG
func (a IntArray) Value() (val driver.Value, err error) {
	str := "{"
	for _, i := range a {
		str += strconv.Itoa(i) + ","
	}
	str += "}"

	val = driver.Value(str)

	return
}

// Scan allows postgres to read out an intarray
func (a *IntArray) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	asString := string(asBytes)
	(*a) = strToIntArray(asString)
	return nil
}

func strToIntArray(s string) []int {
	r := strings.Trim(s, "{}")
	a := make([]int, 0, 10)
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.Atoi(t)
		a = append(a, i)
	}
	return a
}

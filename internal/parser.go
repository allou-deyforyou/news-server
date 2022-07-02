package internal

import (
	"errors"
	"net/url"
	"strconv"
)

type Params url.Values

func (p Params) Int(name string) (int, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.Atoi(value)
}

func (p Params) Ints(name string) ([]int, error) {
	values := url.Values(p)
	var ints []int
	for _, value := range values[name] {
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func (p Params) Float(name string) (float64, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.ParseFloat(value, 64)
}

func (p Params) Floats(name string) ([]float64, error) {
	values := url.Values(p)
	var floats []float64
	for _, value := range values[name] {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}
	return floats, nil
}

func (p Params) Bool(name string) (bool, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.ParseBool(value)
}

func (p Params) Bools(name string) ([]bool, error) {
	values := url.Values(p)
	var bools []bool
	for _, value := range values[name] {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		bools = append(bools, b)
	}
	return bools, nil
}

func (p Params) String(name string) (string, error) {
	values := url.Values(p)
	value := values.Get(name)
	if len(value) == 0 {
		return "", errors.New("empty value")
	}
	return value, nil
}

func (p Params) Strings(name string) ([]string, error) {
	values := url.Values(p)
	value := values[name]
	if len(value) == 0 {
		return nil, errors.New("empty value")
	}
	return value, nil
}

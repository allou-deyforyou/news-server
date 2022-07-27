package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"sync"

	"google.golang.org/protobuf/proto"
)

type Params url.Values

func (p Params) Int(name string) (int, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.Atoi(value)
}

func (p Params) IntX(name string, defaultValue int) int {
	value, err := p.Int(name)
	if err != nil {
		return defaultValue
	}
	return value
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

func (p Params) IntsX(name string, defaultValue []int) []int {
	value, err := p.Ints(name)
	if err != nil {
		return defaultValue
	}
	return value
}

func (p Params) Float(name string) (float64, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.ParseFloat(value, 64)
}

func (p Params) FloatX(name string, defaultValue float64) float64 {
	value, err := p.Float(name)
	if err != nil {
		return defaultValue
	}
	return value
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

func (p Params) FloatsX(name string, defaultValue []float64) []float64 {
	value, err := p.Floats(name)
	if err != nil {
		return defaultValue
	}
	return value
}

func (p Params) Bool(name string) (bool, error) {
	values := url.Values(p)
	value := values.Get(name)
	return strconv.ParseBool(value)
}

func (p Params) BoolX(name string, defaultValue bool) bool {
	value, err := p.Bool(name)
	if err != nil {
		return defaultValue
	}
	return value
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

func (p Params) BoolsX(name string, defaultValue []bool) []bool {
	value, err := p.Bools(name)
	if err != nil {
		return defaultValue
	}
	return value
}

func (p Params) String(name string) (string, error) {
	values := url.Values(p)
	value := values.Get(name)
	if len(value) == 0 {
		return "", errors.New("empty value")
	}
	return value, nil
}

func (p Params) StringX(name string, defaultValue string) string {
	value, err := p.String(name)
	if err != nil {
		return defaultValue
	}
	return value
}

func (p Params) Strings(name string) ([]string, error) {
	values := url.Values(p)
	value := values[name]
	if len(value) == 0 {
		return nil, errors.New("empty value")
	}
	return value, nil
}

func (p Params) StringsX(name string, defaultValue []string) []string {
	value, err := p.Strings(name)
	if err != nil {
		return defaultValue
	}
	return value
}

func Filter[T any](values []T, f func(T) bool) (result []T) {
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func Shuffle[T any](data []T) []T {
	length := len(data)
	result := make([]T, length)
	perm := rand.Perm(length)
	for i, v := range perm {
		result[v] = data[i]
	}
	return result
}

func Remove[T any](data []T, f func(T, T) bool) []T {
	for i := 0; i < len(data); i++ {
		for j := len(data) - 1; j > i; j-- {
			item := data[i]
			last := data[j]
			if f(item, last) {
				data = append(data[:j], data[j+1:]...)
			}
		}
	}
	return data
}

func ProtoDecode(r io.Reader, m proto.Message) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(bytes, m)
}

func ProtoEncode(w io.Writer, m proto.Message) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func JsonDecode(r io.Reader, d interface{}) error {
	return json.NewDecoder(r).Decode(d)
}

func JsonEncode(w io.Writer, d interface{}) error {
	return json.NewEncoder(w).Encode(d)
}

func RecoverFunc(group *sync.WaitGroup) {
	if r := recover(); r != nil {
		group.Done()
		log.Println("Recovered in f", r)
	}
}

package convert

import "strconv"

type StringTo string

func (s StringTo) String() string {
	return string(s)
}

func (s StringTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StringTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StringTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StringTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func (s StringTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	return v, err
}

func (s StringTo) MustInt64() int64 {
	v, _ := s.Int64()
	return v
}

func (s StringTo) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

func (s StringTo) MustFloat64() float64 {
	v, _ := strconv.ParseFloat(s.String(), 64)
	return v
}

package project

type Numeric interface {
	int
	int64
	float64
}

type Value struct {
	val interface{}
}

func newValue(val interface{}) *Value {
	return &Value{val}
}

func (v *Value) asBoolean() bool {
	return v.val != "" && v.val != nil && v.val != 0 && v.val != false
}

func (v *Value) asFloat() float64 {
	value, ok := v.val.(float64)
	if !ok {
		panic("invalid numeric value")
	}
	return value
}

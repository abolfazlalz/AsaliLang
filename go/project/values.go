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
	return toFloat(v.val)
}

func toFloat(val interface{}) float64 {
	if val == 0 {
		return 0
	}
	value, ok := val.(float64)
	if !ok {
		panic("invalid numeric value")
	}
	return value
}

func toBoolean(val interface{}) bool {
	return val != "" && val != nil && val != 0 && val != false
}

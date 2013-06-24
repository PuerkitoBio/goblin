package runtime

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOpAddOnBool = errors.New("cannot apply Add on a Bool value")
	ErrInvalidOpSubOnBool = errors.New("cannot apply Sub on a Bool value")
	ErrInvalidOpMulOnBool = errors.New("cannot apply Mul on a Bool value")
	ErrInvalidOpDivOnBool = errors.New("cannot apply Div on a Bool value")
	ErrInvalidOpPowOnBool = errors.New("cannot apply Pow on a Bool value")
	ErrInvalidOpModOnBool = errors.New("cannot apply Mod on a Bool value")
	ErrInvalidOpUnmOnBool = errors.New("cannot apply Unm on a Bool value")
)

// Bool is the representation of the Boolean type. It is equivalent
// to Go's bool type.
type Bool bool

func (ø Bool) dump() string {
	return fmt.Sprintf("%v (Bool)", bool(ø))
}

// Int returns 1 if true, 0 if false.
func (ø Bool) Int() int {
	if bool(ø) {
		return 1
	}
	return 0
}

// Float returns 1 if true, 0 if false.
func (ø Bool) Float() float64 {
	if bool(ø) {
		return 1.0
	}
	return 0.0
}

// String returns "true" if true, "false" otherwise.
func (ø Bool) String() string {
	if bool(ø) {
		return "true"
	}
	return "false"
}

// Bool returns the boolean value itself.
func (ø Bool) Bool() bool {
	return bool(ø)
}

func (ø Bool) Native() interface{} {
	return bool(ø)
}

// Cmp compares two values as booleans.
func (ø Bool) Cmp(v Val) int {
	if bool(ø) == v.Bool() {
		return 0
	} else if bool(ø) {
		// If it is true, other is necessarily false, so this one is greater
		return 1
	} else {
		// Necessarily smaller
		return -1
	}
}

// Add is an invalid operation.
func (ø Bool) Add(v Val) Val {
	panic(ErrInvalidOpAddOnBool)
}

// Sub is an invalid operation.
func (ø Bool) Sub(v Val) Val {
	panic(ErrInvalidOpSubOnBool)
}

// Mul is an invalid operation.
func (ø Bool) Mul(v Val) Val {
	panic(ErrInvalidOpMulOnBool)
}

// Div is an invalid operation.
func (ø Bool) Div(v Val) Val {
	panic(ErrInvalidOpDivOnBool)
}

// Mod is an invalid operation.
func (ø Bool) Mod(v Val) Val {
	panic(ErrInvalidOpModOnBool)
}

// Pow is an invalid operation.
func (ø Bool) Pow(v Val) Val {
	panic(ErrInvalidOpPowOnBool)
}

// Unm is an invalid operation.
func (ø Bool) Unm() Val {
	panic(ErrInvalidOpUnmOnBool)
}

package simplemath
// so the file name does not have to match the package name
import (
	"errors"
	"fmt"
	"math"
)

// make the first letter of func name capitalized, then the function is public
func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

// return an error along with float64
func Divide(p1, p2 float64) (float64, error) {
	fmt.Println("In divide, p1: ", p1, ", p2: ", p2)
	// in go, check for errors early and short-circuit function
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by 0")
	}
	return p1/p2, nil
}

// naming return values, we can get rid of return keywords except at the end
// may only be good for smaller functions
func NamedDivide(p1, p2 float64) (answer float64, err error) {
	fmt.Println("In divide, p1: ", p1, ", p2: ", p2)
	// in go, check for errors early and short-circuit function
	if p2 == 0 {
		err = errors.New("cannot divide by 0")
	}
	answer = p1/p2
	return
}

// variadic function
// only the final parameter can be variadic
func sum_func(values ...float64) float64 {
	// the arg is a slice of float64 
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

type SemanticVersion struct {
	major, minor, patch int

}

func NewSemanticVersion(majorArg, minorArg, patchArg int) SemanticVersion {
	return SemanticVersion {
		majorArg,
		minorArg,
		patchArg,
	}
}

// this is tied to SemanticVersion
// is this a method because it is tied to a struct? 
// now you can make a SemanticVersion struct (with a name like semVer) and call semVer.ToString()
func (sv SemanticVersion) ToString() string {
	return fmt.Sprintf("%d.%d.%d\n", sv.major, sv.minor, sv.patch)
}
// this will modify a copy, immutable
func (sv SemanticVersion) ImIncrementMajor() SemanticVersion {
	sv.major += 1
	return sv
}
// this will modify a copy, immutable
func (sv SemanticVersion) ImIncrementMinor() SemanticVersion {
	sv.minor += 1
	return sv
}
// this will modify a copy, immutable
func (sv SemanticVersion) ImIncrementPatch() SemanticVersion {
	sv.patch += 1
	return sv
}

// using pointer-based method receiver
// or pointer receiver
// good if you have large data
func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}
// using pointer-based method receiver
func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}
// using pointer-based method receiver
func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}


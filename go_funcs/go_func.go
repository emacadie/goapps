package main

import (
	"errors"
	"fmt"
	"net/http"
	"math"
	"strings"
	"simplemath"
)

func add(p1, p2 float64) float64 {
	return p1 + p2
}

// return an error along with float64
func divide(p1, p2 float64) (float64, error) {
	fmt.Println("In divide, p1: ", p1, ", p2: ", p2)
	// in go, check for errors early and short-circuit function
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by 0")
	}
	return p1/p2, nil
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

func doMathStuff() {
	fmt.Printf("%f\n", add(6, 2))
	answer, err := divide(6, 3)
	if err != nil {
		fmt.Println("Here is err: ", err)
	} else {
		fmt.Println("Divide 3 by 2: ", answer)
	}
	answer2, err2 := divide(6, 0)
	if err2 != nil {
		fmt.Println("Here is err: ", err2)
	} else {
		fmt.Println("Divide 3 by 2: ", answer2)
	}
	// ignore error return (you could ignore both, but do not ignore errors):
	answer3, _ := divide(6, 3)
	fmt.Println("Divide 6 by 3: ", answer3)
	fmt.Println("variadic function, summing 12.4, 14, 16, 18.9: ", sum_func(12.4, 14, 16, 18.9))
	fmt.Println("variadic function, summing 12.4, 14, 16: ", sum_func(12.4, 14, 16))
	// create a slice
	nums := []float64{14.4, 25.5, 36}
	total := sum_func(nums...) // the ... expands the slice into separate args
	fmt.Println("Summing ", nums, ": ", total)
	fmt.Println("Add from simplemath: ", simplemath.Add(6,3))
	// sanswer, serr := simplemath.Divide(6,3)
	// fmt.Println("Divide from simplemath: ", sanswer, ", err: ", serr)
	// fmt.Println("simplemath.sum_func(11.1, 12.2, 13.3, 15.5): ", simplemath.sum_func(11.1, 12.2, 13.3, 15.5))
	// ./go_func.go:61:63: cannot refer to unexported name simplemath.sum_func
	// ./go_func.go:61:63: undefined: simplemath.sum_func
	// it's defined, it's just not accessible
	nAnswer, nErr := simplemath.NamedDivide(6, 5)
	if nErr == nil {
		fmt.Println("simplemath.NamedDivide(6,5): ", nAnswer)
	}
}

func doHttpStuff() {
	// changed return of RoundTrip to a pointer, so must deref here
	var tripper http.RoundTripper = &RoundTripCounter{}
	
	r, _ := http.NewRequest(http.MethodGet, "http://github.com/EMacAdie", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
	// so how to get a count?
}
// our implementation of http.RoundTripper
type RoundTripCounter struct {
	count int
}
// the only method you need to implement
// need to make it a pointer
func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

func doSemanticVersionStuff() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	fmt.Println("Our SemanticVersion, calling sv.ToString(): ", sv.ToString())
	// using our immutable 
	sv.ImIncrementMajor()
	fmt.Println("Our SemanticVersion, calling sv.ToString() after ImIncrementMajor: ", sv.ToString()) // no change
	// to change, got to do this:
	sv = sv.ImIncrementMajor()
	fmt.Println("Our SemanticVersion, calling sv.ToString() after ImIncrementMajor and setting sv to new value: ", sv.ToString()) // no change

	sv2 := simplemath.NewSemanticVersion(1, 2, 3)
	fmt.Println("Our SemanticVersion, calling sv2.ToString(): ", sv2.ToString())
	sv2.IncrementMajor()
	fmt.Println("Our SemanticVersion, calling sv2.ToString() after IncrementMajor: ", sv2.ToString()) 
}

func doAnonFuncs() {
	// anon functions
	func() {
		println("My first anon func in go")
	}()

	// set anon func to a variable
	a := func() {
		println("My second anon func in go")
	} // notice no "()" at the end
	// then call a()
	a()
	// isn't setting it to a variable kind of the same as declaring it?
	b := func(name string) {
		fmt.Printf("Hello to %s from an anon function\n", name)
	} // notice no "()" at the end
	b("joe")
	b("Mary")
	c := func(name string) string {
		fmt.Printf("Hello to %s, we will upper case name", name)
		return strings.ToUpper(name)
	}
	upperCase := c("john")
	fmt.Println("calling c on 'john' gave us:", upperCase)
}

func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

func usingAdd() func(float64, float64) float64 {
	return simplemath.Add
}
// gotta look up "type"
type MathExpr = string

const (
	AddExpr = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func switchMathExpr(expr MathExpr) func(float64, float64) float64 {
	switch expr{
		case AddExpr:
		return simplemath.Add
		case SubtractExpr:
		return simplemath.Subtract
		case MultiplyExpr:
		return simplemath.Multiply
		default:
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}
}

func doFuncsFromFuncs() {
	addExpr := mathExpression()
	println("after setting mathExpression to a var, Math expression gives us: ", addExpr(2, 3))
	// we can't do usingAdd(2,3)
	usingAddExpr := usingAdd()
	println("calling a func that returns simplemath.Add: ", usingAddExpr(2,3))
	
	// using swithMathExpr
	addExpr2 := switchMathExpr(AddExpr)
	subtractExpr := switchMathExpr(SubtractExpr)
	multiplyExpr := switchMathExpr(MultiplyExpr)
	f1 := 2.0
	f2 := 3.0
	println("Calling addExpr2(f1, f2): ", addExpr2(f1, f2))
	println("Calling subtractExpr(f1, f2): ", subtractExpr(f1, f2))
	println("Calling multiplyExpr(f1, f2): ", multiplyExpr(f1, f2))
}

func main() {
	doMathStuff()
	doHttpStuff()
	doSemanticVersionStuff()
	doAnonFuncs()
	doFuncsFromFuncs() 

	
}


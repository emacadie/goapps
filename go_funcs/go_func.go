package main

import (
	"errors"
	"fmt"
	"io"
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

// take a function made from a MathExpr, and double the result
// so our function takes two float64s, and also a func, and we will return a float64
// the func that we are using as an arg takes two float64s and returns a float64
func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
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

	fmt.Printf(" calling double on 3, 2, addExpr2: %f \n", double(3, 2, addExpr2))
	fmt.Printf(" calling double on 3, 2, switchMathExpr(AddExpr): %f \n", double(3, 2, switchMathExpr(AddExpr)))
	fmt.Printf(" calling double on 3, 2, switchMathExpr(MultiplyExpr): %f \n", double(3, 2, switchMathExpr(MultiplyExpr)))
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func workWithFunctionsAndState() {
	p2 := powerOfTwo()
	p2Result := p2() 
	fmt.Println("here is p2Result: ", p2Result)
	p2Result = p2() // changing value of p2Result
	// x got incremented, so go preserves x
	// anti-functional
	fmt.Println("here is p2Result: ", p2Result) 
	
}

func workWithBadStateInAnonFuncs() {
	var funcs []func() int64 // a slice of functions
	for i := 0; i < 10; i++ {
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(i), 2))
		})
	}
	for _, f := range funcs {
		println("calling f: ", f()) // we get 100 each time, so it uses the same i for each
	}
	println("--- let's try again")
	var funcs2 []func() int64 // a slice of functions
	for i2 := 0; i2 < 10; i2++ {
		cleanI := i2 // give each function its own variable
		funcs2 = append(funcs2, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	for _, f2 := range funcs2 {
		println("calling f2 again: ", f2()) // we get 100 each time
	}
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func ReadSomethingBad() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	if _, err := r.Read([]byte("test something")); err != nil {
		fmt.Printf("An error occurred in ReadSomethingBad: %s\n", err)
		return err
	}
	return nil
}

func doErrorHandling() {
	ReadSomethingBad()
}

type SimpleReader struct {
	count int
}

// need to make SimpleReder a pointer here - otherwise we are changing a copy
func (br *SimpleReader) Read(p []byte) (n int, err error) {
	println("br count:", br.count)
	if br.count > 3 {
		return 0, io.EOF
		// return 0, errors.New("bad robot")
	}
	br.count += 1
	return br.count, nil
}

func ReadSomething() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	_, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("An error occurred in ReadSomething: %s\n", err)
		return err
	}
	return nil
}

func ReadFullFile() error {
	// SimpleReader needs to be a reference since our method uses a pointer receiver
	var r io.Reader = &SimpleReader{}
	for {
		value, err := r.Read([]byte("text that does nothing"))
		if err == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if err != nil {
			return err
		}
		println("Here is value:", value)
	}
	return nil
}

func continueOnError() {
	if err := ReadFullFile(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		println("Something bad occurred")
	}
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
}

func ReadFullFile02() error {
	// SimpleReader needs to be a reference since our method uses a pointer receiver
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
	}() // like a finally in Java
	// first on defer stack, last one called
	// so I guess all defer funcs get called if they are put on stack
	// usually you only need one

	defer func() {
		println("before for loop")
	}()
	for {
		value, err := r.Read([]byte("text that does nothing"))
		if err == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if err != nil {
			return err
		}
		println("Here is value:", value)
	}
	defer func() {
		println("after for loop")
	}() // defer funcs go on a stack, it is the last one in, first one called
	return nil
}

func useDeferFunctions() {
	println("In useDeferFunctions")
	// ReadFullFile02()
	if err := ReadFullFile02(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		println("Something bad occurred")
	}
}

type SimplePanicReader struct {
	count int
}

func (br *SimplePanicReader) Close() error {

	println("closing SimplePanicReader")
	return nil
}

// need to make SimplePanicReader a pointer here - otherwise we are changing a copy
func (br *SimplePanicReader) Read(p []byte) (n int, err error) {
	println("br count:", br.count)
	if br.count == 2 {
		panic( "something really really bad")
	}
	if br.count > 3 {
		return 0, io.EOF
		// return 0, errors.New("bad robot")
	}
	br.count += 1
	return br.count, nil
}

func ReadFullFilePanic() error {
	// SimplePanicReader needs to be a reference since our method uses a pointer receiver
	var r io.ReadCloser = &SimplePanicReader{}
	defer func() {
		_ = r.Close()
	}() 

	defer func() {
		println("before for loop")
	}()
	for {
		value, err := r.Read([]byte("text that does nothing"))
		if err == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if err != nil {
			return err
		}
		println("Here is value:", value)
	}
	defer func() {
		println("after for loop")
	}() 
	return nil
}


func usePanics() {
	// you could use panics in switch statements
	// list of planets in the solar system: if someone sends an invalid one, throw a panic

	println("In usePanics")
	if err := ReadFullFilePanic(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		println("Something bad occurred")
	}
}

// here we are naming our return value
func ReadFullFilePanicAndRecover() (err error) {
	// SimplePanicReader needs to be a reference since our method uses a pointer receiver
	var r io.ReadCloser = &SimplePanicReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p != nil {
			println("our panic: ", p)
			err = errors.New("a panic occurred but it is okay")
		}
	}() 

	defer func() {
		println("before for loop")
	}()
	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println("Here is value:", value)
	}
	defer func() {
		println("after for loop")
	}() 
	return nil
}

func panicAndRecover() {
	println("In panicAndRecover")
	if err := ReadFullFilePanicAndRecover(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		println("Something bad occurred: ", err)
		fmt.Printf("our error: %s \n", err)
	}
	// our program finished gracefully
}

type SimplePanicRecoverPanicReader struct {
	count int
}

func (br *SimplePanicRecoverPanicReader ) Close() error {

	println("closing SimplePanicRecoverPanicReader ")
	return nil
}

// need to make SimplePanicRecoverPanicReader  a pointer here - otherwise we are changing a copy
func (br *SimplePanicRecoverPanicReader ) Read(p []byte) (n int, err error) {
	println("br count:", br.count)
	if br.count == 2 {
		// panic(errCatastrophicReader)
		panic(errors.New("Another error"))
	}
	if br.count > 3 {
		return 0, io.EOF
		// return 0, errors.New("bad robot")
	}
	br.count += 1
	return br.count, nil
}

var errCatastrophicReader = errors.New("something catastrophic happened in the reader")

// here we are naming our return value
func ReadFullFilePanicAndRecoverAndPanic() (err error) {
	// SimplePanicRecoverPanicReader  needs to be a reference since our method uses a pointer receiver
	var r io.ReadCloser = &SimplePanicRecoverPanicReader {}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println("our panic: ", p)
			err = errors.New("a panic occurred but it is okay")
		} else if p != nil {
			panic("An unexpected error occurred and we do not want to recover")
		}
	}() 

	defer func() {
		println("before for loop")
	}()
	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println("Here is value:", value)
	}
	defer func() {
		println("after for loop")
	}() 
	return nil
}

func panicAndRecoverAndPanic() {
	println("In panicAndRecoverAndPanic")
	if err := ReadFullFilePanicAndRecoverAndPanic(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		println("Something bad occurred: ", err)
		fmt.Printf("our error: %s \n", err)
	}
	// our program finished gracefully
}

func main() {
	// doMathStuff()
	// doHttpStuff()
	// doSemanticVersionStuff()
	// doAnonFuncs()
	// doFuncsFromFuncs() 
	// workWithFunctionsAndState() 
	// workWithBadStateInAnonFuncs()
	// doErrorHandling()
	// continueOnError()
	// useDeferFunctions()
	// usePanics()
	// panicAndRecover()
	panicAndRecoverAndPanic()
}


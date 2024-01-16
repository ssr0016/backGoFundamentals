package utils

import (
	"errors"
	"fmt"
	"math"
)

// Creating custom errors using the New function
// The simplest way to create a custom error is to use the New function of the errors package.

// func New() {
// 	// New returns an error that formats as the given text.
// 	// Each call to New returns a distinct error value even if the text is identical.
// 	return &errorString{text}
// }

// // errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed. Negative radius.")
	}
	return math.Pi * radius * radius, nil
}

func ErrorCustomImplementation() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

// Adding more inforamation the error using Errorf
func circleaArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed. radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func ErrorCustomImplementation2() {
	radius := -20.0
	area, err := circleaArea2(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func ErrorCustomImplementation3() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			fmt.Printf("Area calculation failed. radius %0.2f is less than zero", areaError.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

type areaError1 struct {
	err    string  //error description
	length float64 // length which caused the error
	width  float64 // width which caused the error
}

func (e *areaError1) Error() string {
	return e.err
}

func (e *areaError1) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError1) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError1{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil

}

func RectAreaImplementation() {
	lengt, width := -5.0, -9.0
	area, err := rectArea(lengt, width)
	if err != nil {
		var areaError1 *areaError1
		if errors.As(err, &areaError1) {
			if areaError1.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero", areaError1.length)
			}
			if areaError1.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero", areaError1.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

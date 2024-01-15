package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// What are errors?
// Errors indicate any abnormal condition occurring in the program. Let’s say we are trying to open a file and the file does not exist in the file system. This is an abnormal condition and it’s represented as an error.

// Errors in Go are plain old values. Just like any other built-in type such as int, float64, … error values can be stored in variables, passed as parameters to functions, returned from functions, and so on.

func ErrorHandlingExample() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

type error interface {
	Error() string
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func ErrorTypeRepresentation() {
	f, err := os.Open("text.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

type DNSError struct {
	Op   string
	Host string
	Err  error
}

func (e *DNSError) Error() string {
	return e.Op + " " + e.Host + ": " + e.Err.Error()
}

func (e *DNSError) Timeout() bool {
	return e.Err == os.ErrDeadlineExceeded
}

func (e *DNSError) Temporary() bool {
	return e.Err == os.ErrClosed
}

func RetrievingMoreInformationUsingMethods() {
	addr, err := net.LookupHost("golang123.com")
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.Timeout() {
				fmt.Println("Operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("Temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

func DirectComparison() {
	files, err := filepath.Glob("[")
	if err != nil {
		if errors.Is(err, filepath.ErrBadPattern) {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic Error", err)
		return
	}
	fmt.Println("matched files", files)
}

func DoNotIgnoreErrors() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}

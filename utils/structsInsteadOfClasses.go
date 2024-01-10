package utils

import "fmt"

type EmployeeStruct struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e EmployeeStruct) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func New(firstName, lastName string, totalLeaves, leavesTaken int) EmployeeStruct {
	e := EmployeeStruct{
		FirstName:   firstName,
		LastName:    lastName,
		TotalLeaves: totalLeaves,
		LeavesTaken: leavesTaken,
	}
	return e
}

func ImplementingEmployeeStructs() {
	e := New("Sams", "Recs", 30, 10)
	e.LeavesRemaining()
}

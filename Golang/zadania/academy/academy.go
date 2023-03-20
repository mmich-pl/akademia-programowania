package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	return int(math.Round(float64(sum) / float64(len(grades))))
}

// Count returns amount of elements in slice satisfying the given predicate
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, val := range slice {
		if predicate(val) {
			count++
		}
	}
	return count
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) == 0 {
		return 0.0
	}
	count := Count(attendance, func(x bool) bool { return x })

	return float64(count) / float64(len(attendance))
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	studentAttendance := AttendancePercentage(s.Attendance)
	studentAverageGrade := AverageGrade(s.Grades)

	if studentAverageGrade == 1 || studentAttendance <= 0.59 || s.Project == 1 {
		return 1
	}

	finalGrade := float64(s.Project+studentAverageGrade) / 2

	if studentAttendance <= 0.79 {
		finalGrade--
	}
	return int(math.Round(finalGrade))
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	summary := make(map[string]uint8, len(students))

	for _, val := range students {
		summary[val.Name] = uint8(FinalGrade(val))
	}
	return summary
}

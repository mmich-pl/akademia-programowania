package academy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeYear(t *testing.T) {
	testScenarios := [3]struct {
		testName       string
		mockRepository func(t *testing.T) *MockRepository
	}{
		{
			testName: "No students",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().List(uint8(2)).Return(nil, nil)
				return mockRepository
			},
		},
		{
			testName: "One student",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().List(uint8(2)).Return([]string{"John Doe"}, nil)
				mockRepository.EXPECT().Get("John Doe").Return(&Sophomore{
					name:       "John Doe",
					grades:     []int{1, 2, 1},
					project:    3,
					attendance: []bool{true, true, false},
				}, nil)
				mockRepository.EXPECT().Save("John Doe", uint8(2)).Return(nil)
				return mockRepository
			},
		},
		{
			testName: "Many students",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().List(uint8(2)).Return([]string{"John Doe", "Jacob Lee", "Merry Ann"}, nil)

				// Fail because of grades
				mockRepository.EXPECT().Get("John Doe").Return(&Sophomore{
					name:       "John Doe",
					grades:     []int{1, 2, 1},
					project:    1,
					attendance: []bool{true, true, false},
				}, nil)

				// Fail because of attendance
				mockRepository.EXPECT().Get("Jacob Lee").Return(&Sophomore{
					name:       "Jacob Lee",
					grades:     []int{3, 4, 4},
					project:    4,
					attendance: []bool{true, true, false, false, false, false},
				}, nil)

				// Pass
				mockRepository.EXPECT().Get("Merry Ann").Return(&Sophomore{
					name:       "Merry Ann",
					grades:     []int{4, 5, 4, 3},
					project:    5,
					attendance: []bool{true, true, false, true, true, true},
				}, nil)

				mockRepository.EXPECT().Save("John Doe", uint8(2)).Return(nil)
				mockRepository.EXPECT().Save("Jacob Lee", uint8(2)).Return(nil)
				mockRepository.EXPECT().Save("Merry Ann", uint8(3)).Return(nil)
				return mockRepository
			},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			mockRepository := scenario.mockRepository(t)
			err := GradeYear(mockRepository, uint8(2))
			assert.NoError(t, err)
		})
	}
}

func TestGradeStudent(t *testing.T) {
	testScenarios := [5]struct {
		testName               string
		mockRepository         func(t *testing.T) *MockRepository
		studentsInTestScenario []string
	}{
		{
			testName: "No existing student",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Merry Ann").Return(nil, ErrStudentNotFound)
				return mockRepository
			},
			studentsInTestScenario: []string{"Merry Ann"},
		},
		{
			testName: "Positive final grade",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Merry Ann").Return(&Sophomore{
					name:       "Merry Ann",
					grades:     []int{4, 5, 4, 3},
					project:    5,
					attendance: []bool{true, true, false, true, true, true},
				}, nil)
				mockRepository.EXPECT().Save("Merry Ann", uint8(3)).Return(nil)
				return mockRepository
			},
			studentsInTestScenario: []string{"Merry Ann"},
		},
		{
			testName: "Negative final grade (grades)",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Merry Ann").Return(&Sophomore{
					name:       "Merry Ann",
					grades:     []int{1, 2, 1, 3},
					project:    1,
					attendance: []bool{true, true, false, true, true, true},
				}, nil)
				mockRepository.EXPECT().Save("Merry Ann", uint8(2)).Return(nil)
				return mockRepository
			},
			studentsInTestScenario: []string{"Merry Ann"},
		},
		{
			testName: "Negative final grade (attendance)",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Merry Ann").Return(&Sophomore{
					name:       "Merry Ann",
					grades:     []int{3, 4, 4},
					project:    4,
					attendance: []bool{true, true, false, false, false, false},
				}, nil)
				mockRepository.EXPECT().Save("Merry Ann", uint8(2)).Return(nil)
				return mockRepository
			},
			studentsInTestScenario: []string{"Merry Ann"},
		},
		{
			testName: "Mixin scenario",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Merry Ann").Return(&Sophomore{
					name:       "Merry Ann",
					grades:     []int{3, 4, 4},
					project:    4,
					attendance: []bool{true, true, true, true, false},
				}, nil)
				mockRepository.EXPECT().Save("Merry Ann", uint8(3)).Return(nil)

				mockRepository.EXPECT().Get("Jacob Lee").Return(&Sophomore{
					name:       "Jacob Lee",
					grades:     []int{3, 3, 5},
					project:    3,
					attendance: []bool{true, true, false, false, false, false},
				}, nil)
				mockRepository.EXPECT().Save("Jacob Lee", uint8(2)).Return(nil)

				return mockRepository
			},
			studentsInTestScenario: []string{"Merry Ann", "Jacob Lee"},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			mockRepository := scenario.mockRepository(t)
			for _, name := range scenario.studentsInTestScenario {
				err := GradeStudent(mockRepository, name)
				assert.NoError(t, err)
			}
		})
	}
}

# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

# Runtime requirements:
Golang 1.24+


# Project Structure:
```
esep-grade-calculator/
│
├── grade-calculator/
│   ├── grade_calculator.go
│   ├── grade_calculator_test.go
│   └── go.mod
│
└── README.MD
```
Clone the repo:
```
git clone https://github.com/yoanexposito/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```


# File Descriptions

### grade_calculator.go:

```
Defines the GradeCalculator struct and related types.
Stores grades in separate slices (assignments, exams, essays).
Provides methods to add grades, compute averages, and calculate the final numerical and letter grade.
```
### grade_calculator_test.go
```
Contains unit tests for the GradeCalculator.
Verifies correct letter grade output (A, B, F, etc.).
Provides examples of using AddGrade and GetFinalGrade.
```
### go.mod
```

Declares the Go module name (esep/grade-calculator) and Go version (1.24.4).
Required for dependency management and module-based builds.
```
# Run unit tests:
```
cd grade_calculator/
go test ./...
```

# Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 0, Assignment)
	gradeCalculator.AddGrade("exam 1", 0, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 0, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assign", 70, Assignment)
    gc.AddGrade("exam", 70, Exam)
    gc.AddGrade("essay", 70, Essay)
    if gc.GetFinalGrade() != "C" {
        t.Errorf("Expected C, got %s", gc.GetFinalGrade())
    }
}

func TestGetGradeD(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assign", 60, Assignment)
    gc.AddGrade("exam", 60, Exam)
    gc.AddGrade("essay", 60, Essay)
    if gc.GetFinalGrade() != "D" {
        t.Errorf("Expected D, got %s", gc.GetFinalGrade())
    }
}

func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Errorf("Expected 'assignment', got %s", Assignment.String())
    }
    if Exam.String() != "exam" {
        t.Errorf("Expected 'exam', got %s", Exam.String())
    }
    if Essay.String() != "essay" {
        t.Errorf("Expected 'essay', got %s", Essay.String())
    }
}

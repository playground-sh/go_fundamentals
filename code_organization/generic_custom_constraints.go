package code_organization

import (
	"fmt"
	"go_fundamentals/internal/types"
)

func GenericCustomConstraintDemo() {
	testScores := []int{95, 75, 82, 63, 85}
	gpa, _ := calculateGPA(testScores)
	fmt.Println("G.P.A.", gpa)

	testScoresF := []float32{97.5, 75.5, 50.5, 63.2, 93.1}
	gpa2, _ := calculateGPA(testScoresF)
	fmt.Println("G.P.A.", gpa2)
}

func getGradePoint(score float64) float32 {
	switch {
	case score >= 80:
		return 4.0
	case score >= 70:
		return 3.5
	case score >= 60:
		return 3.0
	case score >= 50:
		return 2.5
	default:
		return 2.0
	}
}

func calculateGPA[V types.Number](scores []V) (float32, error) {
	count := len(scores)
	if count == 0 {
		return 0, fmt.Errorf("cannot calculate GPA for empty scores")
	}

	var gpa float32
	for _, v := range scores {
		gpa += getGradePoint(float64(v))
	}
	return gpa / float32(count), nil
}

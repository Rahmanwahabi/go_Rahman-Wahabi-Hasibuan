package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func (s Student) avgScore() float64 {
	total := 0
	for _, val := range s.score {
		total += val
	}
	return float64(total) / float64(len(s.score))
}

func (s Student) minMaxScore() (string, int, string, int) {
	minScore := s.score[0]
	minName := s.name
	maxScore := s.score[0]
	maxName := s.name

	for i := 1; i < len(s.score); i++ {
		if s.score[i] < minScore {
			minScore = s.score[i]
			minName = s.name
		}
		if s.score[i] > maxScore {
			maxScore = s.score[i]
			maxName = s.name
		}
	}

	return minName, minScore, maxName, maxScore
}

func main() {
	var students []Student
	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scan(&score)

		students = append(students, Student{name: name, score: []int{score}})
	}

	// Hitung nilai Rata-rata
	var totalScore int
	for _, s := range students {
		totalScore += s.score[0]
	}
	avgScore := float64(totalScore) / float64(len(students))
	fmt.Printf("\nAverage Score: %.0f\n", avgScore)

	// mencari nilai terkecil dan trebesar
	minName, minScore, maxName, maxScore := students[0].minMaxScore()
	for i := 1; i < len(students); i++ {
		nameMin, scoreMin, nameMax, scoreMax := students[i].minMaxScore()
		if scoreMin < minScore {
			minScore = scoreMin
			minName = nameMin
		}
		if scoreMax > maxScore {
			maxScore = scoreMax
			maxName = nameMax
		}
	}
	fmt.Printf("Min Score of Students: %s (%d)\n", minName, minScore)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxName, maxScore)
}

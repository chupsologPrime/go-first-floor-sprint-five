package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps must be positive")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("the weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("the height must be positive")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("the duration must be positive")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes() * walkingCaloriesCoefficient) / minInH
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps must be positive")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("the weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("the height must be positive")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("the duration must be positive")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	avgSpeed := distance / duration.Hours()
	return avgSpeed

}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps)
	return distance / float64(mInKm)
}

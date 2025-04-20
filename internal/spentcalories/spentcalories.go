package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	vals := strings.Split(data, ",")
	if len(vals) < 3 {
		return 0, vals[1], 0, fmt.Errorf("ошибка:недосточно данных")
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return 0, vals[1], 0, fmt.Errorf("ошибка(шаги)")
	}
	dur, err := time.ParseDuration(vals[2])
	if err != nil {
		return 0, vals[1], 0, fmt.Errorf("ошибка (продолжительность)")
	}
	return steps, vals[1], dur, nil
}

func distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	dist := (float64(steps) * stepLenght) / mInKm

	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeActivity, dur, err := parseTraining(data)
	if err != nil {
		return " ", fmt.Errorf("ошибка")
	}
	duration := dur.Hours()
	dist := distance(steps, height)
	meanSp := meanSpeed(steps, height, dur)

	switch typeActivity {
	case "Бег":
		burnCcal, _ := RunningSpentCalories(steps, weight, height, dur)

		return fmt.Sprintf("Тип тренировки:%s\nДлительность: %.2f ч.\nДистанция:%.2f\nСкорость:%.2f км/ч\nСожгли калорий:%.2f\n", typeActivity, duration, dist, meanSp, burnCcal), nil
	case "Ходьба":
		burnCcal, _ := WalkingSpentCalories(steps, weight, height, dur)
		return fmt.Sprintf("Тип тренировки:%s\nДлительность: %.2f ч.\nДистанция:%.2f\nСкорость:%.2f км/ч\nСожгли калорий:%.2f\n", typeActivity, duration, dist, meanSp, burnCcal), nil
	default:
		return " ", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0.0, fmt.Errorf("ошибка:некорректные данные")
	}
	meanSp := meanSpeed(steps, height, duration)
	return (weight * meanSp * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0.0, fmt.Errorf("ошибка:некорректные данные")
	}
	meanSp := meanSpeed(steps, height, duration)
	return ((weight * meanSp * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}

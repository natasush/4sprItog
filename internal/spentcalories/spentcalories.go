package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
	if len(vals) != 3 {
		return 0, "", 0, fmt.Errorf("parseTraining: wrong data")
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return 0, "", 0, errors.New("parseTraining: steps")
	}
	dur, err := time.ParseDuration(vals[2])
	if err != nil {
		return 0, "", 0, errors.New("parseTraining: duration")
	}
	if steps <= 0 || dur <= 0 {
		return 0, "", 0, errors.New("parseTraining: steps <= 0 || duration <= 0")
	}
	return steps, vals[1], dur, nil
}

func distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	dist := (float64(steps) * stepLenght) / mInKm

	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0.0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeActivity, dur, err := parseTraining(data)
	if err != nil {
		log.Println("mistake: ", err)
		return "", err
	}
	if steps <= 0 || dur <= 0 || weight <= 0 || height <= 0 {
		return "", errors.New("TrainingInfo:steps <= 0 || dur <= 0 || weight <= 0 || height <=0 ")
	}
	duration := dur.Hours()
	dist := distance(steps, height)
	meanSp := meanSpeed(steps, height, dur)

	switch typeActivity {
	case "Бег":
		burnCcal, err := RunningSpentCalories(steps, weight, height, dur)
		if err != nil {
			log.Println("mistake: ", err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, duration, dist, meanSp, burnCcal), nil
	case "Ходьба":
		burnCcal, err := WalkingSpentCalories(steps, weight, height, dur)
		if err != nil {
			log.Println("mistake: ", err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, duration, dist, meanSp, burnCcal), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("RunningSpentCalories: wrong data")
	}
	meanSp := meanSpeed(steps, height, duration)
	return (weight * meanSp * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {

		return 0.0, errors.New("RunningSpentCalories: wrong data")
	}

	meanSp := meanSpeed(steps, height, duration)
	return ((weight * meanSp * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}

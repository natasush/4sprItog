package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	vals := strings.Split(data, ",")
	if len(vals) < 2 {
		return 0, 0, fmt.Errorf("недостаточно данных")
	}

	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка (шаги)")
	}
	dur, err := time.ParseDuration(vals[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка (продолжительность)")
	}
	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, dur, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return " "
	}
	if steps <= 0 {
		return " "
	}
	countSteps := steps
	dist := (float64(countSteps) * stepLength) / mInKm

	burnCalories := WalkingSpentCalories(steps, weight, height, dur)

	return fmt.Sprintf("Количество шагов:%d.\nДистанция составила%.2fкм\nВы сожгли%.2f ккал\n ", countSteps, dist, burnCalories)

}

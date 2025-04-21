package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	vals := strings.Split(data, ",")
	if len(vals) != 2 {
		err := errors.New("parsePackage: wrong len data")
		return 0, 0, err
	}

	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return 0, 0, errors.New("parsePackage: steps")
	}

	dur, err := time.ParseDuration(vals[1])
	if err != nil {
		return 0, 0, errors.New("parsePackage: duration")
	}

	if steps <= 0 || dur <= 0 {
		return 0, 0, errors.New("parsePackage: duration<=0 || steps <=0")
	}

	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, dur, err := parsePackage(data)
	if err != nil {
		log.Println("mistake: ", err)
		return ""
	}

	if steps <= 0 || dur <= 0 {
		err = errors.New("данные не могут быть <=0")
		return ""
	}

	dist := (float64(steps) * stepLength) / mInKm

	burnCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, dur)
	if err != nil {
		log.Println("mistake: ", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, burnCalories)

}

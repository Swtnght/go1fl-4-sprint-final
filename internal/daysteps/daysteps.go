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

	var slay = strings.Split(data, ",")
	if len(slay) != 2 {
		return 0, 0, errors.New("invalid data")
	}
	step, err := strconv.Atoi(slay[0])
	if err != nil {
		return 0, 0, err
	}
	if step <= 0 {
		return 0, 0, errors.New("step can't be negative")
	}
	duration, err := time.ParseDuration(slay[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("timestep can't be negative")
	}
	return step, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, dur, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println(err)
		return ""
	}
	longDistance := (float64(steps) * stepLength) / mInKm
	burnedCalories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, dur)
	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", steps, longDistance, burnedCalories)
}

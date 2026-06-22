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

	var slay []string = strings.Split(data, ",")
	if len(slay) != 3 {
		return 0, "", 0, errors.New("incorrect count of slay")
	}
	step, err := strconv.Atoi(slay[0])
	if err != nil {
		return 0, "", 0, err
	}
	if step <= 0 {
		return 0, "", 0, errors.New("step must be greater than zero")
	}
	duration, err := time.ParseDuration(slay[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("duration must be greater than zero")
	}
	return step, slay[1], duration, nil
}

func distance(steps int, height float64) float64 {

	return (float64(steps) * stepLengthCoefficient * height) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration.Seconds() <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	slay := strings.Split(data, ",")
	if len(slay) != 3 {
		return "", errors.New("incorrect count of slay")
	}
	step, types, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	avgSpeed := meanSpeed(step, height, duration)

	switch types {
	case "Бег":
		callories, errInner := RunningSpentCalories(step, weight, height, duration)
		if errInner != nil {
			log.Println(errInner)
		}
		return fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			types, duration.Hours(), float64(distance(step, height)), avgSpeed, callories), nil

	case "Ходьба":
		callories, errInner := WalkingSpentCalories(step, weight, height, duration)
		if errInner != nil {
			log.Println(errInner)
		}
		return fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			slay[1], duration.Hours(), float64(distance(step, height)), avgSpeed, callories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("steps can't be less than 1")
	}
	if duration <= 0 {
		return 0, errors.New("duration can't be less than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight can't be less than 0")
	}
	if height <= 0 {
		return 0, errors.New("height can't be less than 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * avgSpeed * durationInMinutes) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps < 1 {
		return 0, errors.New("steps can't be less than 1")
	}
	if duration <= 0 {
		return 0, errors.New("duration can't be less than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight can't be less than 0")
	}
	if height <= 0 {
		return 0, errors.New("height can't be less than 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * avgSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

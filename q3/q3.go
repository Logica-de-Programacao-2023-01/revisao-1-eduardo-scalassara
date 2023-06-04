package q3

import "errors"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	max := numbers[0]
	min := numbers[0]
	sum := 0

	for _, num := range numbers {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}

		sum += num
	}

	average := float64(sum-max-min) / float64(len(numbers)-2)
	return max, min, average, nil
}

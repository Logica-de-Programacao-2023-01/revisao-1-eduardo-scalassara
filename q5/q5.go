package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	switch fromScale {
	case "C":
		switch toScale {
		case "F":
			return (temp * 9/5) + 32, nil
		case "K":
			return temp + 273.15, nil
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	case "F":
		switch toScale {
		case "C":
			return (temp - 32) * 5/9, nil
		case "K":
			return (temp - 32) * 5/9 + 273.15, nil
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	case "K":
		switch toScale {
		case "C":
			return temp - 273.15, nil
		case "F":
			return (temp - 273.15) * 9/5 + 32, nil
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	default:
		return 0, fmt.Errorf("escala inválida")
	}
}


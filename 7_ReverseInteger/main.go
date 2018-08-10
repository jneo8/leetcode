package reverse_integer

import (
    "math"
)

func reverse(x int) int {

    values := []int{}
    negNumber := false
    // Neg number
    if x < 0 {
        x *= -1
        negNumber = true
    }

    for x > 0 {
        values = append(values, x % 10)
        x -= (x % 10)
        x /= 10
    }
    result := 0
    for i := 0; i < len(values); i++ {
        secondPower := len(values) - (i + 1)
        result += int(float64(values[i]) *  math.Pow10(secondPower))
    }
    if negNumber {
        result *= -1
    }

    // 32 bit
    if float64(result) > math.Pow(2, 31) || float64(result) < - math.Pow(2, 31) {
        return 0
    }
    return result
}

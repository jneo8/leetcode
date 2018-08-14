package string_to_integer_atoi

import (
    "strconv"
    "math"
    "fmt"
)

func myAtoi(str string) int {
    negNum := 1
    resultStr := ""

    for idx, s := range str {
        s := string(s)
        fmt.Printf("%s + ", resultStr)
        fmt.Printf("%s = ", s)
        if s == " " {
            // pass
        } else if s == "+" {
            negNum *= 1
        } else if s == "-" {
            negNum *= -1
        } else if s == "." {
            resultStr += s
        } else {
            _, err := strconv.Atoi(s)
            if err == nil {
                resultStr += s
            }
            if err != nil &&  idx == 0 {
                resultStr = "F"
            }
        }
        fmt.Printf("%s\n", resultStr)

    }
    result, _ := strconv.Atoi(resultStr)
    result *= negNum
    result_float64 := float64(result)

    if result_float64 > math.Pow(2, 31) - 1 {
         result_float64 = math.Pow(2, 31) - 1
    } else if result_float64 < - math.Pow(2, 31) {
         result_float64 = - math.Pow(2, 31)
    }
    result_int := int(result_float64)
    return result_int
}

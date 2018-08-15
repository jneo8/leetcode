package string_to_integer_atoi

import (
    "strconv"
    "math"
)

func myAtoi(str string) int {
    resultStr := ""
    i := 0
    stop := false
    for i < len(str) && stop == false {
        s := string(str[i])
        if s == " " && len(resultStr) > 0 {
            stop = true
        } else if s != " " {
            tmp_resultStr := resultStr + s
            _, err := strconv.Atoi(tmp_resultStr)
            if err == nil {
                resultStr = tmp_resultStr
            } else {
                if tmp_resultStr == "-" || tmp_resultStr == "+" {
                    resultStr = tmp_resultStr
                } else {

                    stop = true
                }
            }
        }
        i++
    }
    result, _ := strconv.ParseFloat(resultStr, 64)

    if result > math.Pow(2, 31) - 1 {
        result = math.Pow(2, 31) - 1
    } else if result < - math.Pow(2, 31) {
        result = - math.Pow(2, 31)
    }
    result_int := int(result)
    return result_int
}

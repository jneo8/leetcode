package reverse_integer

import (
    "log"
)

func reverse(x int) int {

    values := []int{}
    for x > 0 {
        values = append(values, x % 10)
        x -= (x % 10)
        x /= 10
    }
    log.Printf("%v", values)

    return 321
}

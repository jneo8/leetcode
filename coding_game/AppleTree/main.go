// https://www.codingame.com/training/hard/apple-tree/
package main

import (
    "fmt"
    "math"
)
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Apple struct {
    x float64
    y float64
    z float64
    r float64
}

func main() {
    var N, index int
    fmt.Scan(&N, &index)
    apples := []Apple{}
    fApples := []Apple{} // falling apples

    for i := 0; i < N; i++ {
        var x, y, z, r int
        fmt.Scan(&x, &y, &z, &r)
        if i != index {
            apples = append(apples, Apple{float64(x), float64(y), float64(z), float64(r)})
        } else {
            fApples = append(fApples, Apple{float64(x), float64(y), float64(z), float64(r)})
        }
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")

    for len(fApples) > 0 {
        hitApples := []Apple{}
        newApples := []Apple{}
        for i := range(apples) {
            falling := false
            for j :=  range(fApples) {
                if math.Pow(math.Pow(apples[i].x - fApples[j].x, 2) + math.Pow(apples[i].y - fApples[j].y, 2), 0.5) < apples[i].r + fApples[j].r && fApples[j].z > apples[i].z {
                    falling = true
                }
            }
            if falling {
               hitApples = append(hitApples, apples[i])
            } else {
                newApples = append(newApples, apples[i])
            }
        }
        apples = newApples
        fApples = hitApples
    }
    fmt.Println(len(apples))// Write answer to stdout
}

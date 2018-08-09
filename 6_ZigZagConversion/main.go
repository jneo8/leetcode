package zigzagconversion

import "log"

func convert(s string, numRows int) string{
    log.Printf("s: %v", s)
    log.Printf("numRows: %v", numRows)
    matrix := map[int]string{}
    for i := 0; i < numRows; i++ {
        matrix[i] = ""
    }
    log.Printf("%v", matrix)
    oneTrip := (2 * numRows) - 2
    for idx, v := range s {
        log.Printf("%v: %c", idx, v)
        if oneTrip == 0 {
            matrix[0] += string(v)
        } else if idx % oneTrip >= numRows {
            tmpNum := (numRows - 2) - ((idx % oneTrip) % numRows)
            log.Printf("tmpNum A: %v", tmpNum)
            matrix[tmpNum] += string(v)
        } else {
            tmpNum := (idx % oneTrip) % numRows
            log.Printf("tmpNum B: %v", tmpNum)
            matrix[tmpNum] += string(v)
        }
    }

    result := ""
    log.Printf("%v", matrix)
    for j := 0; j < numRows; j++ {
        result += matrix[j]
    }
    log.Printf("%v", result)
    return result
}

package longestPalidrome


func longestPalindrome(s string) string {
    // odd or even
    result := ""
    if s == "" {
        return result
    } else {
        result = string(s[0])
    }

    sArr := []rune(s)
    for i := 0; i < len(sArr); i++ {
        l := i - 1
        r := i + 1

        for r < len(sArr) && sArr[r] == sArr[i] {
            r++
        }
        for l >= 0 && sArr[l] == sArr[i] {
            l--
        }
        for r < len(sArr) && l >= 0 && sArr[l] == sArr[r] {
            r++
            l--
        }
        tmp := string(sArr[l+1:r])
        if len(tmp) > len(result) {
            result = tmp
        }

    }
    return result
}

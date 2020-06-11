package main

var mappings = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := make([]string, 0, len(digits))
	if len(digits) == 0 {
		return res
	}
	return dfs([]rune(digits), 0, []rune{}, res)
}

func dfs(digits []rune, currIdx int, curr []rune, res []string) []string {
	if currIdx == len(digits) {
		res = append(res, string(curr))
		return res
	}

	alphs := mappings[digits[currIdx]-'0']
	for _, ch := range alphs {
		curr = append(curr, ch)
		res = dfs(digits, currIdx+1, curr, res)
		curr = curr[:len(curr)-1]
	}

	return res
}

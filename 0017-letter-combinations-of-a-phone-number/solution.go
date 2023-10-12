func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }

    dial := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    var result []string

    var backtrack func(curr string, number string)
    backtrack = func(curr string, number string) {
        if number == "" {
            result = append(result, curr)
        } else {
            chars := dial[number[0] - '2']
            for _, char := range chars {
                backtrack(curr + string(char), number[1:])
            }
        }
    }

    backtrack("", digits)
    return result
}
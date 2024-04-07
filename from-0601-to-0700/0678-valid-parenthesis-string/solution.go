func checkValidString(s string) bool {
    openParIdxStack := []int{}
    starIdxStack := []int{}

    for i := 0; i < len(s); i ++ {
        if s[i] == '(' {
            openParIdxStack = append(openParIdxStack, i)
        } else if s[i] == '*' {
            starIdxStack = append(starIdxStack, i)
        } else {
            if len(openParIdxStack) == 0 && len(starIdxStack) == 0 {
                return false
            } else if len(openParIdxStack) == 0 {
                starIdxStack = starIdxStack[:len(starIdxStack) - 1]
            } else {
                openParIdxStack = openParIdxStack[:len(openParIdxStack) - 1]
            }
        }
    }

    for len(openParIdxStack) > 0 {
        if len(starIdxStack) == 0 {
            return false
        }
        if openParIdxStack[len(openParIdxStack) - 1] > starIdxStack[len(starIdxStack) - 1] {
            return false
        }
        starIdxStack = starIdxStack[:len(starIdxStack) - 1]
        openParIdxStack = openParIdxStack[:len(openParIdxStack) - 1]
    }
    return true
}
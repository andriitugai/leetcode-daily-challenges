func capitalizeTitle(title string) string {
    capitalize := func(s string) string {
        res := strings.ToLower(s)
        if len(s) > 2 {
            res = strings.Title(res)
        }
        return res
    }
    words := strings.Split(title, " ")
    result := []string{}
    for _, word := range words {
        result = append(result, capitalize(word))
    }
    return strings.Join(result, " ")
}
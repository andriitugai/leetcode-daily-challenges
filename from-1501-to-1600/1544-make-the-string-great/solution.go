func makeGood(s string) string {
    var c string
    result := []string{}
   
    for i := 0; i < len(s); i ++ {
        c = string(s[i])
        if len(result) > 0 && string(result[len(result)-1]) != c && strings.ToLower(c) == strings.ToLower(result[len(result)-1]) {
            result = result[:len(result)-1]
        } else {
            result = append(result, c)
        }
    }
    return strings.Join(result, "")
}
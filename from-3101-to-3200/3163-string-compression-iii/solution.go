func compressedString(word string) string {
    result := []string{}
    curC := word[0]
    cnt := 1
    for i := 1; i < len(word); i ++ {
        if word[i] == curC && cnt < 9 {
            cnt += 1
        } else {
            result = append(result, fmt.Sprintf("%d%c", cnt, curC))
            curC = word[i]
            cnt = 1
        }
    }
    result = append(result, fmt.Sprintf("%d%c", cnt, curC))
    return strings.Join(result, "")
}
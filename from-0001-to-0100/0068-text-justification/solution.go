unc fullJustify(words []string, maxWidth int) []string {
    var result []string
    curr_line := make([]string, 0)
    num_chars := 0

    for _, word := range words {
        if num_chars + len(word) + len(curr_line) > maxWidth {
            for i:=0; i < maxWidth - num_chars; i++ {
                s := len(curr_line) - 1
                if s == 0 {
                    s = 1
                }
                curr_line[i % s] += " "
            }
            result = append(result, strings.Join(curr_line, ""))
            curr_line = make([]string, 0)
            num_chars = 0
        }
        curr_line = append(curr_line, word)
        num_chars += len(word)
    } 
    leftover := strings.Join(curr_line, " ")
    leftover += strings.Repeat(" ", maxWidth - len(leftover))

    return append(result, leftover)
}

func decodeMessage(key string, message string) string {
    abc := "abcdefghijklmnopqrstuvwxyz"
    subst := make(map[byte]string)
    ia := 0
    for ik := 0; ik < len(key); ik ++ {
        if key[ik] != ' ' && subst[key[ik]] == "" && ia < len(abc) {
            subst[key[ik]] = string(abc[ia])
            ia += 1
        }
    }
    result := make([]string, len(message))
    for i := 0; i < len(message); i++ {
        if message[i] == ' ' {
            result[i] = " "
        } else {
            result[i] = subst[message[i]]
        }
    }
    return strings.Join(result, "")
}
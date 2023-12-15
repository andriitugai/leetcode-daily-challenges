func destCity(paths [][]string) string {
    src := make(map[string]int)
    dst := make(map[string]int)

    for _, path := range paths {
        src[path[0]] += 1
        dst[path[1]] += 1
    }
    for dest, _ := range dst {
        if src[dest] == 0 {
            return dest
        }
    }

    return "No way!"
}
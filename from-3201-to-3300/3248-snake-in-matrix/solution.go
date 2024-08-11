func finalPositionOfSnake(n int, commands []string) int {
    position := 0
    for _, c := range commands {
        if c == "RIGHT" {
            position += 1
        } else if c == "LEFT" {
            position -= 1
        } else if c == "UP" {
            position -= n
        } else if c == "DOWN" {
            position += n
        } 
    }
    return position
}
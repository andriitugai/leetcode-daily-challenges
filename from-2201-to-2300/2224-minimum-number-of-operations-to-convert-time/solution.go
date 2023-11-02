func convertTime(current string, correct string) int {
    curHours, _ := strconv.Atoi(current[0:2])
    curMins, _ := strconv.Atoi(current[3:])

    corHours, _ := strconv.Atoi(correct[0:2])
    corMins, _ := strconv.Atoi(correct[3:])

    hours := corHours - curHours
    minutes := corMins - curMins
    if minutes < 0 {
        minutes += 60
        hours -= 1
    }

    minutes += hours * 60

    m60 := minutes / 60
    minutes -= m60 * 60
    m15 := minutes / 15
    minutes -= m15 * 15
    m05 := minutes / 5
    minutes -= m05 * 5 

    return m60 + m15 + m05 + minutes
}
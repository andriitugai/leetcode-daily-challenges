unc poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
    if buckets == 1 {
        return 0
    }

    trials := minutesToTest / minutesToDie + 1
    z := trials

    count := 1
    for trials < buckets {
        trials = trials * z
        count += 1
    }
    return count
}
func categorizeBox(length int, width int, height int, mass int) string {
    isBulky := false
    if length >= 10000 || width >= 10000 || height >= 10000 || length * width * height >= 1000000000 {
        isBulky = true
    }

    isHeavy := false
    if mass >= 100 {
        isHeavy = true
    }

    if isBulky && isHeavy {
        return "Both"
    } else if isBulky {
        return "Bulky"
    } else if isHeavy {
        return "Heavy"
    }
    return "Neither"
}
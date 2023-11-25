func distanceTraveled(mainTank int, additionalTank int) int {
    result := 0
    for mainTank >= 5 {
        result += 50
        mainTank -= 5
        if additionalTank > 0 {
            additionalTank -= 1
            mainTank += 1
        }
    }
    result += mainTank * 10
    return result
}
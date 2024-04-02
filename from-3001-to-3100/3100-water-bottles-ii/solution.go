func maxBottlesDrunk(numBottles int, numExchange int) int {
    drinked := numBottles
    for numBottles >= numExchange {
        numBottles -= numExchange
        numBottles += 1
        numExchange += 1
        drinked += 1
    }
    return drinked
}
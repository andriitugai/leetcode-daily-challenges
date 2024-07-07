func numWaterBottles(numBottles int, numExchange int) int {
    drinkedBottles := numBottles
    emptyBottles := numBottles
    for emptyBottles >= numExchange {
        drinkedBottles += emptyBottles / numExchange
        emptyBottles = emptyBottles / numExchange + emptyBottles % numExchange
    }
    return drinkedBottles
}
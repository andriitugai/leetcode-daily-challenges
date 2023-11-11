func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    energyRequired := 0
    for _, e := range energy {
        energyRequired += e
    }
    energyToGain := energyRequired - initialEnergy + 1
    if energyToGain < 0 {
        energyToGain = 0
    }

    experienceToGain := 0
    curExperience := initialExperience
    for _, e := range experience {
        deltaToWin := e - curExperience + 1
        if deltaToWin > experienceToGain {
            experienceToGain = deltaToWin
        }
        curExperience += e
    }

    return energyToGain + experienceToGain
}
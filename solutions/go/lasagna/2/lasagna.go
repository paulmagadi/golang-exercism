package lasagna

const OvenTime = 40
const TimeToCookEachLayer = 2

func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
    return numberOfLayers * TimeToCookEachLayer
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    return PreparationTime(numberOfLayers) + actualMinutesInOven
}

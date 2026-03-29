package cars
import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
    return (successRate/100) * float64(productionRate)
    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
    workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(workingCarsPerHour / 60) 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	numOfTense := carsCount / 10
	reminder := carsCount % 10
	fmt.Println("Reminder: ", reminder)
	fmt.Println("numOfTense: ", int(numOfTense))
	result := uint((int(numOfTense) * 95000) + (reminder * 10000))
	fmt.Println("result", result)

	return result

}

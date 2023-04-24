/* Rate of losing 3500 calories per pound of body fat
and 100 calories per mile of brisk walking (15min/mi)
is based on Harvard Health Publishing data at
https://www.health.harvard.edu/staying-healthy/simple-math-equals-easy-weight-loss */

package main

import (
	"fmt"
)


func main() {

	// How much body fat in lbs someone would like to lose. Apply language to receive input.
	var bodyFatLose float32 = 15

	// 3500 calories per pound of body fat lost, based in bodyFatLose input.
	var calsToBurn float32 = bodyFatLose * 3500

	// How many miles someone is willing to walk per day. Apply language to receive input.
	var milesPerDay float32 = 5
	
	// Number of daily miles * 100 calories per mile * 7 days per week / 3500 calories yields lbs lost per week.
	var weeklyPounds float32 = milesPerDay * 100 * 7 / 3500

	// Amount of body fat of original goal divided by amount of total pounds per week.
	var weeksToGoal float32 = bodyFatLose / weeklyPounds

	// 4.345 is the average number of weeks per month. So number of weeks / 4.345 yields number of months.
	var monthsToGoal float32 = weeksToGoal / 4.345



	fmt.Print("Disclaimer: This should never be used as a substitute for direct medical advice from your doctor or qualified health professional.\n")

	fmt.Println("If you would like to lose", bodyFatLose, "lbs of body fat,")
	fmt.Println("that means you need to use", calsToBurn, " more calories than you consume. It sounds like a lot, but let's make it simple!")
	fmt.Println("If you are willing to walk", milesPerDay, "miles per day at a pace of about 15 minutes per mile,")
	fmt.Println("that means, if you keep your calorie consumption about the same, you can lose", weeklyPounds, "lbs per week.")
	fmt.Println("You can reach your goal in just", weeksToGoal, "weeks, or", monthsToGoal, "months!")
	fmt.Println("You are very close! Let the countdown begin!")
	
}
package main

import "lecture04/race_condition"

func main() {
	//Run 1st example
	race_condition.Increment_example()

	//Run 2nd example
	race_condition.Map_example()

	//Run 3rd example
	race_condition.Unbuffered_chanel_example()
}

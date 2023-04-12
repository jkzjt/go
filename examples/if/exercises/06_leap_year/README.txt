/* 
 In the Gregorian calendar, which is the most widely used calendar system in the world, 
 a leap year occurs every year that is evenly divisible by 4, except for years that are 
 both divisible by 100 and not divisible by 400. For example, the year 2020 was a leap 
 year, but the year 2100 will not be a leap year, even though it is divisible by 4, 
 because it is also divisible by 100 and not by 400.
*/

// -----------------------------------------------------------------------------
// EXERCISE: Leap Year
//  Check whether a given year from the cmd is leap year
// 
// EXPECTED OUTPUT
//  ./main.exe 
//   Give me a year number
// 
//  ./main.exe ten
//   "ten" is an invalid year number
// 
//  ./main.exe 2019
//   2019 is not a leap year
// 
//  ./main.exe 2020
//   2020 is a leap year
// 
//  ./main.exe 2100
//   2100 is not a leap year
// -----------------------------------------------------------------------------

/*
	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}
*/
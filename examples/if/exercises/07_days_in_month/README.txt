// -------------------------------------------------
// EXERCISE: Days In A Month
//  Output the number of days in a month 
// 
// HINT
//  1. Note that leap year
//   Use package time#Now.Year to get the current local year
//   https://pkg.go.dev/time@latest#Time.Year
//  2. It should be case-insensitive
//   Use package strings#ToLower to reach it
//   https://pkg.go.dev/strings@latest#ToLower
// 
// EXPECTED OUTPUT 
//  ./main.exe 
//   Give me a month
// 
//  ./main.exe day30
//   "day30" is not a month 
// 
//  ./main.exe January
//   "January" has 31 days
// 
//  ./main.exe MAY
//   "MAY" has 31 days
// 
//  ./main.exe JuNe
//   "JuNe" has 30 days
// 
//  ./main.exe FEBRUARY
//   "FEBRUARY" has 28 days
// -------------------------------------------------------------------
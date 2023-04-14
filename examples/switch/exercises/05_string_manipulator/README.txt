// ------------------------------------------------------------
// EXERCISE: String Manipulator
//  Print the treated string according to operator
// 
// HINT
//  See strings#ToLower https://pkg.go.dev/strings@latest#ToLower
// 	See strings#ToUpper https://pkg.go.dev/strings@latest#ToUpper
//  See strings#Title https://pkg.go.dev/strings@latest#Title
// 
// EXPECTED OUTPUT
//  ./mian.exe
//   [command] [argument]
//   
//   Available commands: lower, upper, title
//
//  ./main.exe lower 
//   [command] [argument]
//   
//   Available commands: lower, upper, title
// 
//  ./main.exe snake "Hello World" 
//   Unknown command: "snake"
// 
//  ./main.exe lower "Hello World" 
//   hello world
// 
//  ./main.exe upper "hi, gopher" 
//   HI, GOPHER
// 
//  ./main.exe title "it's so crazy" 
//   It's So Crazy
// 
//  ./main.exe title "looks such a beaty" "where?" 
//   More aguments: expect 1 but find 2
// 
// ------------------------------------------------------
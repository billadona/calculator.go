package main
 
import (
        "bufio"
        "fmt"
        "os"
        "stack"
        )
 
var operatorStack = stack.NewStack()
var operandStack = stack.NewStack()
var fBool bool = false												// tells the other function in the program if the input is a float
var k int = 0 														// holds the numbers from the int input
var mult float64 = 0.1												// holds the multiplier use to shift the number to the right decimal place
var f float64 = 0.0													// holds the numbers from the float input
var w float64 = 0.0													// holds the float value of the given ascii value
var left int = 0													// left int operand
var right int = 0													// right int operand
var fLeft float64 = 0												// left float operand
var fRight float64 = 0												// right float operand

func precedence(op byte) uint8 {									// this function determines the order of precedence
    switch op {
        case '+', '-': return 0
        case '*', '/': return 1
        default: panic("illegal operator")
    }
}
 
func apply() {
    op := operatorStack.Pop().(byte)
 
    if fBool == false {
        right = operandStack.Pop().(int)							// holds the left int operand
        left = operandStack.Pop().(int)								// holds the right int operand

        switch op {													// does the basic operation for the int input
	        case '+': operandStack.Push(left + right)
	        case '-': operandStack.Push(left - right)
	        case '*': operandStack.Push(left * right)
	        case '/': operandStack.Push(left / right)
	        default: panic("illegal operator")
  		}

    } else {
        fRight = operandStack.Pop().(float64)						// holds the left float operand
        fLeft = operandStack.Pop().(float64)						// holds the right float operand

        switch op {													// does the basic operation for the float input
	        case '+': operandStack.Push(fLeft + fRight)
	        case '-': operandStack.Push(fLeft - fRight)
	        case '*': operandStack.Push(fLeft * fRight)
	        case '/': operandStack.Push(fLeft / fRight)				// does the division for float
	        default: panic("illegal operator")
  		}
    }
}
 
func main() {
    // Read a from Stdin.
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    var ind int = 0
    var i int = 0
    for i = 0; i < len(line); {										// goes through the whole input
        switch line[i] {
            case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9','.':
                for {
                    if line[i] == '.' {
                        i++
                        fBool = true 								// toggle the fBool to true								
                        ind = i 									// set ind so we know where we left off
                        break										// get out of the loop if the input is a float
                    }
                    k = k * 10 + int(line[i] - '0')					// change the ascii value to a number, multiply it by 10 to shift the number to the left
                    i++
                    ind = i 										// ind is used to initialize s in the next loop, avoids s from reading in numbers that are already read
                    if i == len(line) || !('0' <= line[i] && line[i] <= '9') {
                        break
                    }
                }
                if fBool == true {									// if the input is a float
                    f = float64(k)									// convert the int to a float and store in f
                    for s := ind; s < len(line); {
                        w = float64(line[s] - '0')					// convert the ascii value from line[s] to an actual number
                        w = w * mult 								// shift the number to the right decimal place
                        f = f + w 									// add the new decimal value to the number
                        mult = mult * 0.1							// multiplier that shifts the number from line[s] to the right
                        s++											// tells the loop to go the next number in line
                        i++
                        if s == len(line) || !('0' <= line[s] && line[s] <= '9') {
                            break
                            fBool = true							// fBool determines if the input contains a float
                        }
                    }
                    operandStack.Push(f)							// if all number in float is read, push it to stack
                } else {
                    operandStack.Push(k)							// if all number in the integer input is read, push it to the stack
                }
            case '+', '-', '*', '/':
                k = 0												// reset the int holder for the loop back to 0
                f = 0.0												// reset the the float holder for the loop back to 0.0
                mult = 0.1											// reset the multiplier back to 0.1
                for !operatorStack.IsEmpty() && precedence(operatorStack.Top().(byte)) >= precedence(line[i]) {
                    apply()
                }
                operatorStack.Push(line[i])
                i++
            case ' ': i++											// this case allows space as a delimeter
            default: panic("illegal character")						// other inputs are not allowed
        }
    }
    for !operatorStack.IsEmpty() {
        apply()
    }
    if fBool == false {
        r := operandStack.Pop().(int)								// prints out the answer in an int var
        fmt.Println(r)
    } else {
        fR := operandStack.Pop().(float64)							// stores the float answer into a float var for printing
        fmt.Println(fR)
    }
}

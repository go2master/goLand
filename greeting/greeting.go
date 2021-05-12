package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*In Go, a function whose name starts with a capital letter can be called by a function not in the same package.
This is known in Go as an exported name.
在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。
*/
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	/*In Go, the := operator is a shortcut for declaring and initializing a variable in one line.
	在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)

	*/
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func Hellos(names []string) (map[string]string, error) {
	/*In Go, you initialize a map with the following syntax: make(map[key-type]value-type). */
	messages := make(map[string]string)
	/*In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value.
	You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.*/
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}

/* Go executes init functions automatically at program startup, after global variables have been initialized. */
func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	/*This tells Go that the size of the array underlying the slice can be dynamically changed.*/
	formats := []string{"Hi,%v. Welcome!", "Nice to meet you, %v!", "Hail,%v! Well met!"}
	return formats[rand.Intn(len(formats))]
}

/*当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。*/
func add(x, y int) int {
	return x + y
}

/*Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
没有明确初始值的变量声明会被赋予它们的 零值。
数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/
var c, python, java bool

/*变量的初始化*/
var i, j int = 1, 2
var cX, pythonX, javaX = true, false, "no!"

/*
Go 的基本类型有
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
*/
/*
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。cannot use y (type int) as type float64 in assignment.
表达式 T(v) 将值 v 转换为类型 T。
i := 42
f := float64(i)
u := uint(f)
*/
/*
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
使用 const 关键字声明常量。
常量不能用 := 语法声明。
const Pi = 3.14
*/

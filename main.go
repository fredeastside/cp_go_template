package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type float = float32
type double = float64
type long = int64
type ulong = uint64

type Readable interface {
	bool | int | uint64 | int64 | double | float | byte
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func print(arg ...interface{})            { fmt.Fprint(out, arg...) }
func println(arg ...interface{})          { fmt.Fprintln(out, arg...) }
func printf(f string, arg ...interface{}) { fmt.Fprintf(out, f, arg...) }
func scanf(f string, arg ...interface{})  { fmt.Fscanf(in, f, arg...) }

func readLn() string {
	ln, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.Trim(ln, " \n\r")
}

func readString() string {
	return readLn()
}

func readStrings() []string {
	return strings.Split(readLn(), " ")
}

func read[T Readable](f func(string) (T, error), input ...string) T {
	if len(input) == 0 {
		input = []string{readLn()}
	}

	res, err := f(input[0])
	if err != nil {
		panic(err)
	}

	return res
}

func readMany[T Readable](f func(...string) T) []T {
	input := readStrings()
	arr := make([]T, len(input))
	for i := 0; i < len(arr); i++ {
		arr[i] = f(input[i])
	}

	return arr
}

func readBool(input ...string) bool {
	return read(strconv.ParseBool, input...)
}

func readInt(input ...string) int {
	return read(strconv.Atoi, input...)
}

func readByte(input ...string) byte {
	return byte(read(func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 8)
	}, input...))
}

func readDouble(input ...string) double {
	return read(func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}, input...)
}

func readFloat(input ...string) float {
	return float32(read(func(s string) (float64, error) {
		return strconv.ParseFloat(s, 32)
	}, input...))
}

func readLong(input ...string) long {
	return read(func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	}, input...)
}

func readULong(input ...string) ulong {
	return read(func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	}, input...)
}

func readBinary(input ...string) long {
	return read(func(s string) (int64, error) {
		return strconv.ParseInt(s, 2, 64)
	}, input...)
}

func readBools() []bool {
	return readMany(readBool)
}

func readInts() []int {
	return readMany(readInt)
}

func readBytes() []byte {
	return readMany(readByte)
}

func readDoubles() []double {
	return readMany(readDouble)
}

func readFloats() []float {
	return readMany(readFloat)
}

func readLongs() []long {
	return readMany(readLong)
}

func readULongs() []ulong {
	return readMany(readULong)
}

func main() {
	defer out.Flush()

	// solution
	
}

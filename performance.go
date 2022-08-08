package main

import (
	"log"
	"math"
	"reflect"
	"runtime"
	"time"
)

func intToAlphabet(num int) string {
	if num == 0 {
		return "a"
	}

	const ALPHA = "abcdefghijklmnopqrstuvwxyz"
	res := ""
	for num > 0 {
		res = string(ALPHA[num%len(ALPHA)]) + res
		num = int(math.Floor(float64(num / len(ALPHA))))
	}

	return res
}

func mapOperation(count int) interface{} {
	var dict map[string]int = make(map[string]int)
	for i := 0; i < count; i++ {
		dict[intToAlphabet(i)] = i
	}
	var sum int64 = 0
	for key := range dict {
		sum += int64(dict[key])
	}
	return sum
}

func test(operation func(int) interface{}, count int) {
	funcName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
	t0 := time.Now().UnixMilli()
	for i := 0; i < count; i++ {
		operation(3000000)
		log.Printf("[%s] Round %d done.", funcName, i)
	}
	log.Printf("[%s] Time used: %fs", funcName, (float32(time.Now().UnixMilli()-t0))/float32(1000))
}

func main() {
	test(mapOperation, 50)
}

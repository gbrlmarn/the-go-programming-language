// Exercise 1.3: Experiment to measure the difference in running time between our potential inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and section 11.4 shows how to write benchmark test for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func echo1() {
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}

func echo2() {
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
}

func echo3() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}

func GetFnName(i interface {}) string {
  return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func measureTime(fToMeasure func()) {
  starting_time := time.Now().Nanosecond()
  fn_name := GetFnName(fToMeasure)
  fToMeasure()
  result := time.Now().Nanosecond() - starting_time
  fmt.Println(
    fn_name, "took:", result, "nanoseconds to finish")
}

func main() {
  measureTime(echo1)
  measureTime(echo2)
  measureTime(echo3)
}

package main

import (
	"fmt"
)

var (
	Succ = Green
  Info = Teal
  Warn = Yellow
  Err = Red
	Spe = Magenta
)

var (
  Black   = Color("\033[1;30m%s\033[0m")
  Red     = Color("\033[1;31m%s\033[0m")
  Green   = Color("\033[1;32m%s\033[0m")
  Yellow  = Color("\033[1;33m%s\033[0m")
  Purple  = Color("\033[1;34m%s\033[0m")
  Magenta = Color("\033[1;35m%s\033[0m")
  Teal    = Color("\033[1;36m%s\033[0m")
  White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
  return func(args ...interface{}) string {
    return fmt.Sprintf(colorString, fmt.Sprint(args...))
  }
}

func statusCodeColor(code int) string {
	switch c := code; {
		case c >= 500:
			return Spe(code)
		case c >= 400:
			return Err(code)
		case c >= 300:
			return Warn(code)
		default:
			return Succ(code)
	}
}
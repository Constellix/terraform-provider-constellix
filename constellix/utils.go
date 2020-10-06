package constellix

import (
	"fmt"
	"log"
)

func toStringList(configured interface{}) []string {
	vs := make([]string, 0, 1)
	val, ok := configured.(string)
	if ok && val != "" {
		vs = append(vs, val)
	}
	return vs
}

func toListOfString(configured interface{}) []string {
	vs := make([]string, 0, 1)
	log.Println(configured.([]interface{}))
	for _, value := range configured.([]interface{}) {
		vs = append(vs, value.(string))
	}
	return vs
}

func toIntList(configured interface{}) []int {
	vs := make([]int, 0, 1)
	val, ok := configured.(int)
	if ok && val != 0 {
		vs = append(vs, val)
	}
	return vs
}

func toListOfInt(configured interface{}) []int {
	vs := make([]int, 0, 1)
	for _, value := range configured.([]interface{}) {
		vs = append(vs, value.(int))
	}
	return vs
}

// toString converts a value to its default string representation,
// and returns an empty string if the value is nil
func toString(i interface{}) string {
	if i != nil {
		return fmt.Sprintf("%v", i)
	}
	return ""
}

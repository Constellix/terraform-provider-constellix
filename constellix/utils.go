package constellix

import (
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


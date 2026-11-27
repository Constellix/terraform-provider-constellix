package constellix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func decodeJSONResponse(body []byte, target interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(body))
	// Keep numeric IDs out of float64 so their digits and decimal notation are preserved.
	decoder.UseNumber()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra interface{}
	if err := decoder.Decode(&extra); err != io.EOF {
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected data after JSON response")
	}
	return nil
}

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

func stripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

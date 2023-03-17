package jsonconv

import (
	"log"
	"testing"
)

func TestCamel2Case(t *testing.T) {
	str := "XxYYxxxYY +YY"
	log.Println("--->", str)

	cases := Camel2Case(str)
	log.Println("--->", cases)

	camel := Case2Camel(cases)
	log.Println("--->", camel)

	cases = Camel2Case(camel)
	log.Println("--->", cases)
}

func TestJsonToObject(t *testing.T) {

}

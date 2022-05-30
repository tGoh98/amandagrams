package main

import (
	"strings"
	"testing"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetAllCombinationsBase(t *testing.T) {
	expected := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	actual := getAllCombinations(expected, 1)
	if !Equal(actual, expected) {
		t.Fatalf(`getAllCombinations expected: %v but received %v`, expected, actual)
	}
}

func TestGetAllCombinationsN2(t *testing.T) {
	alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	actual := getAllCombinations(alphabet, 2)
	if len(actual) != 26*26 || actual[0] != "AA" {
		t.Fatalf(`getAllCombinations has len %v and first entry %v`, len(actual), actual[0])
	}
}

func TestStrSort(t *testing.T) {
	e1 := "dgo"
	a1 := sortStr("dog")
	e2 := "ACILLMOTYYZ"
	a2 := sortStr("ZYMOTICALLY")
	if e1 != a1 || e2 != a2 {
		t.Fatalf(`sortStr failed`)
	}
}

package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestStringContainedInRegex(t *testing.T) {
	s := "1d12+5"
	re := regexp.MustCompile(`[0-9]+d[0-9]+\+*[0-9]+`)
	fmt.Println(stringContainedInRegex(s, re))
	if !stringContainedInRegex(s, re) {
		t.Errorf("The wrong regexp was used in TestStringContainedInRegex\n String: %s, Regexp: %q\n", s, re)
	}

	failingStrIncorrectAdd := "1d12+5fq"
	if stringContainedInRegex(failingStrIncorrectAdd, re) {
		t.Errorf("failing string should not be contained in regexp\nfailingString: %s, regexp: %q", failingStrIncorrectAdd, re)
	}

	failingStrIncorrectNumDie := "1fd12"
	if stringContainedInRegex(failingStrIncorrectNumDie, re) {
		t.Errorf("failing string should not be contained in regexp\nfailingString: %s, regexp: %q", failingStrIncorrectNumDie, re)
	}

	failingStrIncorrectDieType := "1d12fkj"
	if stringContainedInRegex(failingStrIncorrectDieType, re) {
		t.Errorf("failing string should not be contained in regexp\nfailingString: %s, regexp: %q", failingStrIncorrectDieType, re)
	}
}

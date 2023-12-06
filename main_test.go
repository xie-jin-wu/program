package program

import (
	"fmt"
	"testing"
)

func TestProgramName(t *testing.T) {
	name, err := GetProgramName()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(name)
}

func TestProgramPath(t *testing.T) {
	path, err := GetProgramPath()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(path)
}

package main

import (
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestDrawingStringInStandardFont(t *testing.T) {
	expectedOutput :=` _    _  
| |  | | 
| |__| | 
|  __  | 
| |  | | 
|_|  |_| 
         
         
`

	output, err := exec.Command("go", "run", "main.go", "H").Output()
	if err != nil {
		t.Errorf("Failed to get terminal output: %v", err)
	}//does not match actual
	if expectedOutput != string(output) {
		t.Errorf("\nExpected\n%s\nExpected\n%s", expectedOutput, string(output))
	}
}
package gonap

import "testing"
import "fmt"

import "strconv"

func bad_type(shouldbe, got string) string {
	return " Type is " + got + " should be " + shouldbe
}

func bad_status(wanted, got int) string {
	return " StatusCode is " + strconv.Itoa(got) + " wanted " + strconv.Itoa(wanted)
}

func TestSetAuth(t *testing.T) {
	fmt.Println("Current Username ", Username)
	SetAuth("newuser", "newpass")
	fmt.Println("Applied Username ", Username)
}

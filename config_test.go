package gonap

import "testing"
import "fmt"

func TestSetAuth(t *testing.T) {

	SetAuth("newuser", "newpass")
	fmt.Println(Username)
}

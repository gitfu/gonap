package gonap

import "testing"
//import "fmt"

func mklocid() string {
	resp := ListLocations()

	locid := resp.Items[0].Id
	//fmt.Println(MkJson(resp))
	return locid
}

func TestListLocations(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListLocations()

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}



/**	Fails on apiary server.

func TestGetLocation(t *testing.T) {
	//t.Parallel()
	shouldbe := "location"
	want := 200
	locid := mklocid()
	resp := GetLocation(locid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
**/
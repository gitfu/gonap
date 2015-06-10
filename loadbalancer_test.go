// loadbalancer_test.go
package gonap

import "testing"
import "fmt"

var lbal_dcid = mkdcid()

func mklbalid(lbal_dcid string) string {
	var jason = []byte(`{
					"name":"Original Loadbalancer",
					}`)
	lbal := CreateLoadbalancer(lbal_dcid, jason)

	lbalid := lbal.Id
	return lbalid
}

func ExampleListLoadbalancers() {
	s := ListLoadbalancers(lbal_dcid)
	fmt.Println(s.Resp.StatusCode)
	// Output: 200

}

func TestListLoadbalancers(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	lbals := ListLoadbalancers(lbal_dcid)
	if lbals.Type != shouldbe {
		t.Errorf("ListLoadbalancers() type == %v, wanted %v", lbals.Type, shouldbe)
	}
	if lbals.Resp.StatusCode != want {
		t.Errorf("ListLoadbalancers() StatusCode == %v, wanted %v", lbals.Resp.StatusCode, want)
	}
}

func TestCreateLoadbalancer(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Goat",
					}`)
	newlbal := CreateLoadbalancer(lbal_dcid, jason)
	if newlbal.Resp.StatusCode != want {
		t.Errorf("CreateLoadbalancer() StatusCode == %v, wanted %v", newlbal.Resp.StatusCode, want)
	}
}

func TestGetLoadbalancer(t *testing.T) {
	t.Parallel()
	shouldbe := "loadbalancer"
	want := 200
	lbalid := mklbalid(lbal_dcid)
	lbal := GetLoadbalancer(lbal_dcid, lbalid)

	if lbal.Type != shouldbe {
		t.Errorf("GetLoadbalancer() type == %v, wanted %v", lbal.Type, shouldbe)
	}
	if lbal.Resp.StatusCode != want {
		t.Errorf("GetLoadbalancer() StatusCode == %v, wanted %v", lbal.Resp.StatusCode, want)
	}
}

func TestPatchLoadbalancer(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed Loadbalancer",
					}`)
	lbalid := mklbalid(lbal_dcid)
	lbal := PatchLoadbalancer(lbal_dcid, lbalid, jason_patch)
	if lbal.Resp.StatusCode != want {
		t.Errorf("PatchLoadbalancer() StatusCode == %v, wanted %v", lbal.Resp.StatusCode, want)
	}
}

func TestUpdateLoadbalancer(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Renamed Loadbalancer",
				
					}`)

	lbalid := mklbalid(lbal_dcid)
	lbal := UpdateLoadbalancer(lbal_dcid, lbalid, jason_update)
	if lbal.Resp.StatusCode != want {
		t.Errorf("UpdateLoadbalancer() StatusCode == %v, wanted %v", lbal.Resp.StatusCode, want)
	}
}

func TestDeleteLoadbalancer(t *testing.T) {
	t.Parallel()
	want := 202
	lbalid := mklbalid(lbal_dcid)
	lbal := DeleteLoadbalancer(lbal_dcid, lbalid)
	if lbal.StatusCode != want {
		t.Errorf("DeleteLoadbalancer() StatusCode == %v, wanted %v", lbal.StatusCode, want)
	}
}

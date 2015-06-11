// ipblock_test.go
package gonap

import "testing"

func mkipblkid() string {
	ipb := ReserveIpBlock([]byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`))

	ipblkid := ipb.Id
	return ipblkid
}

func TestListIpBlocks(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	ipbs := ListIpBlocks()

	if ipbs.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, ipbs.Type))
	}
	if ipbs.Resp.StatusCode != want {
		t.Errorf(bad_status(want, ipbs.Resp.StatusCode))
	}
}

func TestReserveIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`)
	ipb := ReserveIpBlock(jason)
	if ipb.Resp.StatusCode != want {
		t.Errorf(bad_status(want, ipb.Resp.StatusCode))

	}
}

func TestGetIpBlock(t *testing.T) {
	//t.Parallel()
	shouldbe := "ipblock"
	want := 200
	ipblkid := mkipblkid()
	ipb := GetIpBlock(ipblkid)

	if ipb.Type != shouldbe {
		t.Errorf("GetIpBlock() type == %v, wanted %v", ipb.Type, shouldbe)
	}
	if ipb.Resp.StatusCode != want {
		t.Errorf(bad_status(want, ipb.Resp.StatusCode))
	}
}

func TestReleaseIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	ipblkid := mkipblkid()
	ipb := ReleaseIpBlock(ipblkid)
	if ipb.StatusCode != want {
		t.Errorf(bad_status(want, ipb.StatusCode))
	}
}

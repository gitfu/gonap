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
	t.Parallel()
	shouldbe := "collection"
	want := 200
	ipbs := ListIpBlocks()

	if ipbs.Type != shouldbe {
		t.Errorf("ListIpblocks() type == %v, wanted %v", ipbs.Type, shouldbe)
	}
	if ipbs.Resp.StatusCode != want {
		t.Errorf("ListIpblocks() StatusCode == %v, wanted %v", ipbs.Resp.StatusCode, want)
	}
}

func TestReserveIpBlock(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`)
	ipb := ReserveIpBlock(jason)
	if ipb.Resp.StatusCode != want {
		t.Errorf("ReserveIpBlock() StatusCode == %v, wanted %v", ipb.Resp.StatusCode, want)
	}
}

func TestGetIpBlock(t *testing.T) {
	t.Parallel()
	shouldbe := "ipblock"
	want := 200
	ipblkid := mkipblkid()
	ipb := GetIpBlock(ipblkid)

	if ipb.Type != shouldbe {
		t.Errorf("GetIpBlock() type == %v, wanted %v", ipb.Type, shouldbe)
	}
	if ipb.Resp.StatusCode != want {
		t.Errorf("GetIpblock() StatusCode == %v, wanted %v", ipb.Resp.StatusCode, want)
	}
}

func TestReleaseIpBlock(t *testing.T) {
	t.Parallel()
	want := 202
	ipblkid := mkipblkid()
	ipb := ReleaseIpBlock(ipblkid)
	if ipb.Resp.StatusCode != want {
		t.Errorf("DeleteIpBlock() StatusCode == %v, wanted %v", ipb.Resp.StatusCode, want)
	}
}

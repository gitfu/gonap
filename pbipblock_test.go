// ipblock_test.go
package gonap

import "testing"

func mkipblkid() string {
	resp := ReserveIpBlock([]byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`))
	ipblk := ToIpblock(resp.Body)
	ipblkid := ipblk.Id
	return ipblkid
}

func TestListIpBlocks(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListIpBlocks()
	ipbs := ToIpblocks(resp.Body)
	if ipbs.Type != shouldbe {
		t.Errorf("ListIpblocks() type == %v, wanted %v", ipbs.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("ListIpblocks() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestReserveIpBlock(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`)
	resp := ReserveIpBlock(jason)
	if resp.StatusCode != want {
		t.Errorf("ReserveIpBlock() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestGetIpBlock(t *testing.T) {
	t.Parallel()
	shouldbe := "ipblock"
	want := 200
	ipblkid := mkipblkid()
	resp := GetIpBlock(ipblkid)
	ipb := ToIpblock(resp.Body)
	if ipb.Type != shouldbe {
		t.Errorf("GetIpBlock() type == %v, wanted %v", ipb.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("GetIpblock() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestReleaseIpBlock(t *testing.T) {
	t.Parallel()
	want := 202
	ipblkid := mkipblkid()
	resp := ReleaseIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf("DeleteIpBlock() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

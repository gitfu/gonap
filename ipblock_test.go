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
	resp := ListIpBlocks()

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestReserveIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`)
	resp := ReserveIpBlock(jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestGetIpBlock(t *testing.T) {
	//t.Parallel()
	shouldbe := "ipblock"
	want := 200
	ipblkid := mkipblkid()
	resp := GetIpBlock(ipblkid)

	
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestReleaseIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	ipblkid := mkipblkid()
	resp := ReleaseIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}


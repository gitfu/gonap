// ipblock_test.go
package goprofitbricks

import "testing"

func mkipblkid() string {
	ipb := ReserveIpBlock([]byte(`{
					"name":"Original IpBlock",
					"location":"us/las"
					}`))

	ipblkid := ipb.Id
	return ipblkid
}

func TestListIPBlocks(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListIPBlocks()

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
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
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}

}

func TestGetIpBlock(t *testing.T) {
	//t.Parallel()
	shouldbe := "ipblock"
	want := 200
	ipblkid := mkipblkid()
	resp := GetIpBlock(ipblkid)

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}

}

func TestReleaseIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	ipblkid := mkipblkid()
	resp := ReleaseIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}

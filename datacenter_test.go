package gonap

import "testing"

func mkdcid() string {
	dc := CreateDatacenter([]byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`))

	dcid := dc.Id
	return dcid
}

func TestListDatacenters(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	dcs := ListDatacenters()

	if dcs.Type != shouldbe {
		t.Errorf("ListDatacenters() type == %v, wanted %v", dcs.Type, shouldbe)
	}
	if dcs.Resp.StatusCode != want {
		t.Errorf("ListDatacenters() StatusCode == %v, wanted %v", dcs.Resp.StatusCode, want)
	}
}

func TestCreateDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`)
	newdc := CreateDatacenter(jason)
	if newdc.Resp.StatusCode != want {
		t.Errorf("CreateDatacenter() StatusCode == %v, wanted %v", newdc.Resp.StatusCode, want)
	}
}

func TestGetDatacenter(t *testing.T) {
	t.Parallel()
	shouldbe := "datacenter"
	want := 200
	dcid := mkdcid()
	dc := GetDatacenter(dcid)
	if dc.Type != shouldbe {
		t.Errorf("ListDatacenters() type == %v, wanted %v", dc.Type, shouldbe)
	}
	if dc.Resp.StatusCode != want {
		t.Errorf("GetDatacenter() StatusCode == %v, wanted %v", dc.Resp.StatusCode, want)
	}
}

func TestPatchDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed DC",
					}`)
	dcid := mkdcid()
	dc := PatchDatacenter(dcid, jason_patch)
	if dc.Resp.StatusCode != want {
		t.Errorf("PatchDatacenter() StatusCode == %v, wanted %v", dc.Resp.StatusCode, want)
	}
}

func TestUpdateDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"name":"Renamed DC",
					"location":"us/las"
					}`)

	dcid := mkdcid()
	dc := UpdateDatacenter(dcid, jason_update)
	if dc.Resp.StatusCode != want {
		t.Errorf("UpdateDatacenter() StatusCode == %v, wanted %v", dc.Resp.StatusCode, want)
	}
}

func TestDeleteDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	dcid := mkdcid()
	resp := DeleteDatacenter(dcid)
	if resp.StatusCode != want {
		t.Errorf("DeleteDatacenter() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

// New  format (method style) testing

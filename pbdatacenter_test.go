package gonap

import "testing"

func mkdcid() string {
	resp := CreateDatacenter([]byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`))
	dc := ToDatacenter(resp.Body)
	dcid := dc.Id
	return dcid
}

func TestListDatacenters(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListDatacenters()
	dcs := ToDatacenters(resp.Body)
	if dcs.Type != shouldbe {
		t.Errorf("ListDatacenters() type == %v, wanted %v", dcs.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("ListDatacenters() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestCreateDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original DC",
					"location":"us/las"
					}`)
	resp := CreateDatacenter(jason)
	if resp.StatusCode != want {
		t.Errorf("CreateDatacenter() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestGetDatacenter(t *testing.T) {
	t.Parallel()
	shouldbe := "datacenter"
	want := 200
	dcid := mkdcid()
	resp := GetDatacenter(dcid)
	dc := ToDatacenter(resp.Body)
	if dc.Type != shouldbe {
		t.Errorf("ListDatacenters() type == %v, wanted %v", dc.Type, shouldbe)
	}
	if resp.StatusCode != want {
		t.Errorf("GetDatacenter() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

func TestPatchDatacenter(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed DC",
					}`)
	dcid := mkdcid()
	resp := PatchDatacenter(dcid, jason_patch)
	if resp.StatusCode != want {
		t.Errorf("PatchDatacenter() StatusCode == %v, wanted %v", resp.StatusCode, want)
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
	resp := UpdateDatacenter(dcid, jason_update)
	if resp.StatusCode != want {
		t.Errorf("UpdateDatacenter() StatusCode == %v, wanted %v", resp.StatusCode, want)
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

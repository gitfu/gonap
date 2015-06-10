// image_test.go

package gonap

import "testing"

func mkimgid() string {
	imgs := ListImages()

	imgid := imgs.Items[0].Id
	return imgid
}

func TestListImages(t *testing.T) {
	t.Parallel()
	shouldbe := "collection"
	want := 200
	imgs := ListImages()

	if imgs.Type != shouldbe {
		t.Errorf("ListImages() type == %v, wanted %v", imgs.Type, shouldbe)
	}
	if imgs.Resp.StatusCode != want {
		t.Errorf("ListImages() StatusCode == %v, wanted %v", imgs.Resp.StatusCode, want)
	}
}

func TestGetImage(t *testing.T) {
	t.Parallel()
	shouldbe := "image"
	want := 200
	imgid := mkimgid()
	img := GetImage(imgid)
	if img.Type != shouldbe {
		t.Errorf("ListImages() type == %v, wanted %v", img.Type, shouldbe)
	}
	if img.Resp.StatusCode != want {
		t.Errorf("GetImage() StatusCode == %v, wanted %v", img.Resp.StatusCode, want)
	}
}

func TestPatchImage(t *testing.T) {
	t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed img",
					}`)
	imgid := mkimgid()
	img := PatchImage(imgid, jason_patch)
	if img.Resp.StatusCode != want {
		t.Errorf("PatchImage() StatusCode == %v, wanted %v", img.Resp.StatusCode, want)
	}
}

func TestUpdateImage(t *testing.T) {
	t.Parallel()
	want := 202
	jason_update := []byte(`{
					"size":80,
					
					}`)

	imgid := mkimgid()
	img := UpdateImage(imgid, jason_update)
	if img.Resp.StatusCode != want {
		t.Errorf("UpdateImage() StatusCode == %v, wanted %v", img.Resp.StatusCode, want)
	}
}

func TestDeleteImage(t *testing.T) {
	t.Parallel()
	want := 202
	imgid := mkimgid()
	resp := DeleteImage(imgid)
	if resp.StatusCode != want {
		t.Errorf("DeleteImage() StatusCode == %v, wanted %v", resp.StatusCode, want)
	}
}

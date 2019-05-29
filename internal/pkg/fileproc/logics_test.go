package fileproc

import (
	"testing"
)

func TestRandToken(t *testing.T) {
	rt := RandToken(10)
	if len(rt) != 20 {
		t.Error("wrong len of random token, expected 10, have:", len(rt))
	}

	rtN := RandToken(10)
	if len(rt) != 20 {
		t.Error("wrong len of random token, expected 10")
	}

	if rt == rtN {
		t.Error("tokens not unique")
	}
}

var casesFileType = []struct {
	In  string
	Exp bool
}{
	{
		In:  "kek",
		Exp: false,
	},
	{
		In:  "image/jpeg",
		Exp: true,
	},
	{
		In:  "image/jpg",
		Exp: true,
	},
	{
		In:  "image/png",
		Exp: true,
	},
}

func TestCheckFileType(t *testing.T) {
	for i, val := range casesFileType {
		r := CheckFileType(val.In)
		if r != val.Exp {
			t.Error("#", i, "expected:", val.Exp, ", have:", r)
		}
	}
}

var casesCreateResultFile = []struct {
	FileName      string
	FileExtension string
	Filetype      string
	FileBytes     []byte
	Exp           string
	ExpErr        string
}{
	{
		FileName:      "kek",
		FileExtension: ".jpg",
		Filetype:      "image",
		FileBytes:     make([]byte, 100),
		ExpErr:        "",
		Exp:           "kek.jpg",
	},
	{
		FileName:      "",
		FileExtension: "",
		Filetype:      "image",
		FileBytes:     make([]byte, 100),
		ExpErr:        "open : no such file or directory",
		Exp:           "kek.jpg",
	},
}

func TestCreateResultFile(t *testing.T) {
	for i, val := range casesCreateResultFile {
		r, err := CreateResultFile(val.FileName, val.FileExtension, val.Filetype, val.FileBytes)
		if err != nil {
			if err.Error() != val.ExpErr {
				t.Error("#", i, "ERROR expected:", val.ExpErr, ", have:", err)
				continue
			}
		} else {
			if r != val.Exp {
				t.Error("#", i, "expected:", val.Exp, ", have:", r)
				continue
			}
		}
	}
}

package osrelease

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		expected *Data
	}{
		{
			"alpine3_9", &Data{
				ID:         "alpine",
				Name:       "Alpine Linux",
				PrettyName: "Alpine Linux v3.9",
				VersionID:  "3.9.2",
			},
		},
		{
			"centos7", &Data{
				ID:         "centos",
				IDLike:     "rhel fedora",
				Name:       "CentOS Linux",
				PrettyName: "CentOS Linux 7 (Core)",
				Version:    "7 (Core)",
				VersionID:  "7",
			},
		},
		{
			"debian9", &Data{
				ID:         "debian",
				Name:       "Debian GNU/Linux",
				PrettyName: "Debian GNU/Linux 9 (stretch)",
				Version:    "9 (stretch)",
				VersionID:  "9",
			},
		},
		// todo: pull the files for these os'
		{"fedora29", &Data{}},
		{"rhel7", &Data{}},
		{
			"ubuntu-bionic", &Data{
				ID:              "ubuntu",
				IDLike:          "debian",
				Name:            "Ubuntu",
				PrettyName:      "Ubuntu 18.04.2 LTS",
				Version:         "18.04.2 LTS (Bionic Beaver)",
				VersionID:       "18.04",
				VersionCodename: "bionic",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs, err := ioutil.ReadFile(filepath.Join("testdata", tt.name))
			if err != nil {
				t.Errorf("cannot read datafile: %v", err)
			}

			info := Parse(string(bs))
			if !reflect.DeepEqual(tt.expected, info) {
				t.Errorf("expected %#v,\n got %#v", tt.expected, info)
			}
		})
	}
}

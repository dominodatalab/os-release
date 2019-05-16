package osrelease

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

var tests = []struct {
	name string
	data *Data
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
	{
		"fedora29", &Data{
			ID:         "fedora",
			Name:       "Fedora",
			PrettyName: "Fedora 29 (Container Image)",
			Version:    "29 (Container Image)",
			VersionID:  "29",
		},
	},
	{
		"rhel7", &Data{
			ID:         "rhel",
			IDLike:     "fedora",
			Name:       "Red Hat Enterprise Linux Server",
			PrettyName: "Red Hat Enterprise Linux 7.6",
			Version:    "7.6 (Maipo)",
			VersionID:  "7.6",
		},
	},
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

func TestParse(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs, err := ioutil.ReadFile(filepath.Join("testdata", tt.name))
			if err != nil {
				t.Errorf("cannot read datafile: %v", err)
			}

			info := Parse(string(bs))
			if !reflect.DeepEqual(tt.data, info) {
				t.Errorf("data %#v,\n got %#v", tt.data, info)
			}
		})
	}
}

func TestIsLikeDebian(t *testing.T) {
	exercisePredicate(t, "debian|ubuntu", func(data *Data) bool {
		return data.IsLikeDebian()
	})
}

func TestIsLikeFedora(t *testing.T) {
	exercisePredicate(t, "centos|fedora|rhel", func(data *Data) bool {
		return data.IsLikeFedora()
	})
}

func TestIsUbuntu(t *testing.T) {
	exercisePredicate(t, "ubuntu", func(data *Data) bool {
		return data.IsUbuntu()
	})
}

func TestIsRhel(t *testing.T) {
	exercisePredicate(t, "rhel", func(data *Data) bool {
		return data.IsRhel()
	})
}

func TestIsCentos(t *testing.T) {
	exercisePredicate(t, "centos", func(data *Data) bool {
		return data.IsCentos()
	})
}

func exercisePredicate(t *testing.T, idPattern string, fn func(*Data) bool) {
	for _, tt := range tests {
		matched, err := regexp.MatchString(idPattern, tt.data.ID)
		if err != nil {
			t.Fatal(err)
		}

		if matched != fn(tt.data) {
			t.Errorf("expected os %q to be %t", tt.data.Name, matched)
		}
	}
}

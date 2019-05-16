package osrelease

import "strings"

const (
	DebianID = "debian"
	FedoraID = "fedora"
	UbuntuID = "ubuntu"
)

// Data exposes the most common identification parameters.
type Data struct {
	ID              string
	IDLike          string
	Name            string
	PrettyName      string
	Version         string
	VersionID       string
	VersionCodename string
}

// Parse expects the contents of /etc/os-release and populates the fields of a Data object.
func Parse(contents string) *Data {
	info := map[string]string{}

	kvPairs := strings.Split(contents, "\n")
	for _, strPair := range kvPairs {
		if strPair != "" {
			kv := strings.Split(strPair, "=")
			info[kv[0]] = strings.Trim(kv[1], "\"")
		}
	}

	return &Data{
		ID:              info["ID"],
		IDLike:          info["ID_LIKE"],
		Name:            info["NAME"],
		PrettyName:      info["PRETTY_NAME"],
		Version:         info["VERSION"],
		VersionID:       info["VERSION_ID"],
		VersionCodename: info["VERSION_CODENAME"],
	}
}

// IsLikeDebian will return true for Debian and any other related OS, such as Ubuntu.
func (d *Data) IsLikeDebian() bool {
	return d.ID == DebianID || strings.Contains(d.IDLike, DebianID)
}

// IsLikeFedora will return true for Fedora and any other related OS, such as CentOS or RHEL.
func (d *Data) IsLikeFedora() bool {
	return d.ID == FedoraID || strings.Contains(d.IDLike, FedoraID)
}

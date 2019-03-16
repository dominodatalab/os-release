package osrelease

import "strings"

type Data struct {
	ID              string
	IDLike          string
	Name            string
	PrettyName      string
	Version         string
	VersionID       string
	VersionCodename string
}

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

//func (d *Data) IsLikeFedora() bool {
//	return false
//}
//
//func (d *Data) IsLikeDebian() bool {
//	return false
//}

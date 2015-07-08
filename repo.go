package goimjson

// Repo is an inteface for storing all changes
type Repo interface {
	Add(id string, blob []byte)
	Lookup(id string) []byte
}

type mapRepo struct {
	allVersions map[string][]byte
}

// NewMapRepo allocates a instance of mapRepo with Repo interface
func NewMapRepo() Repo {
	return &mapRepo{
		allVersions: make(map[string][]byte),
	}
}

func (m *mapRepo) Add(id string, blob []byte) {
	m.allVersions[id] = blob
}

func (m *mapRepo) Lookup(id string) []byte {
	return m.allVersions[id]
}

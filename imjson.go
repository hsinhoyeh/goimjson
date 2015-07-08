package goimjson

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/spaolacci/murmur3"
)

const (
	InvalidVersion = ""
)

// ImJSON wraps a json structure and a version id to provide an immutable json object
type ImJSON struct {

	// data is a simplejson data
	data *simplejson.Json

	// versionRepo holds all chagnes into a repo
	versionRepo Repo

	latestVersion string
}

// New returns a pointer to ImJSON
func New() (*ImJSON, error) {
	return &ImJSON{
		data:        simplejson.New(),
		versionRepo: NewMapRepo()}, nil
}

// NewWithBody allocates and returns an pointer of ImJSON with the marshaled data provided
func NewWithBody(body []byte) (*ImJSON, error) {
	d, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	return &ImJSON{
		data:        d,
		versionRepo: NewMapRepo(),
	}, nil
}

// Interface returns the underlying json data
func (i *ImJSON) Interface() interface{} {
	return i.data.Interface()
}

// Encode encodes the underlying json data into byte slice
func (i *ImJSON) Encode() ([]byte, error) {
	return i.data.Encode()
}

// Set modify the ImJSON by setting a key and a value to it
// The write operation will result in a version which can be used to lookup the original json data
func (i *ImJSON) Set(key string, val interface{}) (version string) {
	// we modify the json data first, and then update to allVersions
	i.data.Set(key, val)
	return i.addVersions()
}

// GetLatest get the latest version with the key
func (i *ImJSON) GetLatest(key string) (*ImJSON, error) {
	return i.Get(i.latestVersion, key)
}

// Get retrieves a key from a json object with a specific version
// the returned value will be a pointer of ImJSON
// TODO: do we need to pass all versions into the newly created ImJSON?
func (i *ImJSON) Get(version string, key string) (*ImJSON, error) {
	vBSlice := i.versionRepo.Lookup(version)
	if len(vBSlice) < 1 {
		return nil, nil // nil represents that the version or key is not found
	}
	imjson, err := NewWithBody(vBSlice)
	if err != nil {
		return nil, nil
	}
	// replace the data by the branch starting with the given key
	imjson.data = imjson.data.Get(key)
	return imjson, nil
}

func (i *ImJSON) addVersions() (verison string) {
	bslice, err := i.data.Encode()
	if err != nil {
		return InvalidVersion
	}
	ver := versionFromBytes(bslice)
	// TODO: detect the collision of versions
	// and need to pruge the old version
	i.versionRepo.Add(ver, bslice)
	i.latestVersion = ver
	return ver
}

func versionFromBytes(b []byte) string {
	hasher := murmur3.New64()
	hasher.Write(b)
	return fmt.Sprintf("%d", hasher.Sum64())
}

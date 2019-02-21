package store

type Store interface {
	Write(string, interface{}) error
	// Read2(string, *interface{}) error
	Read(string, string, interface{}) error
	ReadAll(string, interface{}) error
	Delete(string, string)
	Update(string, string, interface{}, interface{})
}

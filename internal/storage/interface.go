package storage

// Initializes the underlying storage system
type Initializer interface {
	Initialize() (int, error)
}

// Writes the file bytes contents to the underlying system
type Writer interface {
	Write(file_bytes []byte) (int, error)
	// TODO: need to define the model
	// Write(file_package Package) (int, error)
}

// Read the file byte contents from the underlying system
type Reader interface {
	Read(location string) ([]byte, error)
	// TODO: need to define the model
	// Read(location string) (Package, error)
}

type Terminator interface {
	Close() error
}

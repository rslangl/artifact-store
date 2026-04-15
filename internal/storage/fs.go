package storage

import (
	"os"
	//"sync"
)

type FileSystemStorage struct {
	f *os.File
	//mtx sync.Mutex
	// TODO: buffers, encoders, metrics, last_used timestamp etc.
}

func (fss *FileSystemStorage) Initialize() (int, error) {
	// TODO: ensure path exists, with read and write permisions
	return 0, nil
}

func (fss *FileSystemStorage) Write(file_bytes []byte) (int, error) {
	//fss.mtx.Lock()
	//defer fss.mtx.Unlock()
	return fss.f.Write(file_bytes)
}

func (fss *FileSystemStorage) Read(location string) (int, error) {
	//fss.mtx.Lock()
	//defer fss.mtx.Unlock()
	//return fss.f.Read(location)
	return 0, nil
}

func (fss *FileSystemStorage) Close() error {
	//fss.mtx.Lock()
	//defer fss.mtx.Unlock()
	return fss.f.Close()
}

// func Create() *FileSystemStorage {
// 	return &FileSystemStorage{
// 		nil,
// 		//nil,
// 	}
// }

package storage

type NasStorage struct {

}

func (fss *NasStorage) Initialize() (int, error) {
	return 0, nil
}

func (fss *NasStorage) Write(file_bytes []byte) (int, error) {
	return 0, nil
}

func (fss *NasStorage) Read(location string) (int, error) {
	return 0, nil
}

func (fss *NasStorage) Close() error {
	return nil
}

// func Create() *NasStorage {
// 	return &NasStorage{}
// }

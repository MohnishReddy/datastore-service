package services

type dataInterface interface {
	Compress(file *[]byte) (compressedFile *[]byte, err error)
	Decompress(compressedFile *[]byte) (file *[]byte, err error)
}

type data struct{}

func NewDataHandler() dataInterface {
	return &data{}
}

func (d *data) Compress(file *[]byte) (compressedFile *[]byte, err error) {
	compressedFile = file
	return
}

func (d *data) Decompress(compressedFile *[]byte) (file *[]byte, err error) {
	file = compressedFile
	return
}

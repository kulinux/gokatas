package circular

import "errors"

//Buffer is Circula Buffer
type Buffer struct {
	content []byte
	size    int
}

//ReadByte ...
func (b *Buffer) ReadByte() (byte, error) {
	if len(b.content) == 0 {
		return 0, errors.New("Empty")
	}
	res := b.content[0]
	b.content = b.content[1:]
	return res, nil
}

//WriteByte ...
func (b *Buffer) WriteByte(input byte) error {
	if b.size == len(b.content) {
		return errors.New("Full")
	}
	b.content = append(b.content, input)
	return nil
}

//Overwrite ...
func (b *Buffer) Overwrite(input byte) error {
	if len(b.content) == 0 {
		return errors.New("Empty")
	}
	if b.size == len(b.content) {
		b.ReadByte()

	}
	return b.WriteByte(input)
}

//Reset ...
func (b *Buffer) Reset() {
	b.content = make([]byte, 0)
}

//NewBuffer ...
func NewBuffer(size int) *Buffer {
	return &Buffer{
		content: make([]byte, 0),
		size:    size,
	}
}

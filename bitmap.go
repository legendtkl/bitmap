package bitmap

import (
	"errors"
)

type BitMap struct {
	size uint64
	data []uint64
}

// NewBitMap create a bitmap of sizeN
func NewBitMap(sizeN uint64) (*BitMap, error) {
	bitMap := &BitMap{
		size: sizeN,
	}
	bitMap.data = make([]uint64, (sizeN+1)/64+1)
	return bitMap, nil
}

// SetOne set posth bit to 1
func (bitMap *BitMap) SetOne(pos uint64) error {
	if pos > bitMap.size {
		return errors.New("pos out of size")
	}

	bitMap.data[pos/64] |= 1 << (pos % 64)
	return nil
}

// SetZero set posth bit to 0
func (bitMap *BitMap) SetZero(pos uint64) error {
	if pos > bitMap.size {
		return errors.New("pos out of size")
	}

	bitMap.data[pos/64] &= ^(1 << (pos % 64))
	return nil
}

// GetPosition get posth bit
func (bitMap *BitMap) GetPosition(pos uint64) (int64, error) {
	if pos > bitMap.size {
		return 0, errors.New("pos out of size")
	}
	if x := bitMap.data[pos/64] & (1 << (pos % 64)); x != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

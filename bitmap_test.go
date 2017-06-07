package bitmap

import (
	"testing"
)

func TestBasic(t *testing.T) {
	bitMap, _ := NewBitMap(100)
	val, err := bitMap.GetPosition(101)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("pos out of size Test Fail.")
	}

	err = bitMap.SetOne(101)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("pos out of size Test Fail.")
	}

	val, err = bitMap.GetPosition(99)
	if val != 0 {
		t.Error("Test Fail.")
	}

	err = bitMap.SetOne(99)
	if err != nil {
		t.Error("Test Fail.")
	}

	val, err = bitMap.GetPosition(99)
	if val != 1 {
		t.Error("Test Fail.")
	}

	err = bitMap.SetZero(99)
	if err != nil {
		t.Error("Test Fail.")
	}

	val, err = bitMap.GetPosition(99)
	if val != 0 {
		t.Error("Test Fail.")
	}
}

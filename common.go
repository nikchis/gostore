// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

const (
	storeArr0Index int = iota
	storeArr1Index
	storeArr2Index
	storeArr3Index
	storeArr4Index
	storeArr5Index
	storeArr6Index
	storeArr7Index
	storeArr8Index
	storeArr9Index
	storeArrAIndex
	storeArrBIndex
	storeArrCIndex
	storeArrDIndex
	storeArrEIndex
	storeArrFIndex
	storeArrLastIndex
)

type keyChar byte

func (c keyChar) storeArrIndex() int {
	switch c {
	case 48:
		return storeArr0Index
	case 49:
		return storeArr1Index
	case 50:
		return storeArr2Index
	case 51:
		return storeArr3Index
	case 52:
		return storeArr4Index
	case 53:
		return storeArr5Index
	case 54:
		return storeArr6Index
	case 55:
		return storeArr7Index
	case 56:
		return storeArr8Index
	case 57:
		return storeArr9Index
	case 65:
		return storeArrAIndex
	case 66:
		return storeArrBIndex
	case 67:
		return storeArrCIndex
	case 68:
		return storeArrDIndex
	case 69:
		return storeArrEIndex
	case 70:
		return storeArrFIndex
	case 97:
		return storeArrAIndex
	case 98:
		return storeArrBIndex
	case 99:
		return storeArrCIndex
	case 100:
		return storeArrDIndex
	case 101:
		return storeArrEIndex
	case 102:
		return storeArrFIndex
	}
	return storeArrLastIndex
}

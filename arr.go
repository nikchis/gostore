// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

const (
	arr0Index int = iota
	arr1Index
	arr2Index
	arr3Index
	arr4Index
	arr5Index
	arr6Index
	arr7Index
	arr8Index
	arr9Index
	arrAIndex
	arrBIndex
	arrCIndex
	arrDIndex
	arrEIndex
	arrFIndex
	arrGIndex
	arrHIndex
	arrIIndex
	arrJIndex
	arrKIndex
	arrLIndex
	arrMIndex
	arrNIndex
	arrOIndex
	arrPIndex
	arrQIndex
	arrRIndex
	arrSIndex
	arrTIndex
	arrUIndex
	arrVIndex
	arrWIndex
	arrXIndex
	arrYIndex
	arrZIndex
	arrLastIndex
)

type keyChar byte

func (c keyChar) arrIndex() int {
	switch c {
	case 48:
		return arr0Index
	case 49:
		return arr1Index
	case 50:
		return arr2Index
	case 51:
		return arr3Index
	case 52:
		return arr4Index
	case 53:
		return arr5Index
	case 54:
		return arr6Index
	case 55:
		return arr7Index
	case 56:
		return arr8Index
	case 57:
		return arr9Index
	case 65:
		return arrAIndex
	case 66:
		return arrBIndex
	case 67:
		return arrCIndex
	case 68:
		return arrDIndex
	case 69:
		return arrEIndex
	case 70:
		return arrFIndex
	case 71:
		return arrGIndex
	case 72:
		return arrHIndex
	case 73:
		return arrIIndex
	case 74:
		return arrJIndex
	case 75:
		return arrKIndex
	case 76:
		return arrLIndex
	case 77:
		return arrMIndex
	case 78:
		return arrNIndex
	case 79:
		return arrOIndex
	case 80:
		return arrPIndex
	case 81:
		return arrQIndex
	case 82:
		return arrRIndex
	case 83:
		return arrSIndex
	case 84:
		return arrTIndex
	case 85:
		return arrUIndex
	case 86:
		return arrVIndex
	case 87:
		return arrWIndex
	case 88:
		return arrXIndex
	case 89:
		return arrYIndex
	case 90:
		return arrZIndex
	case 97:
		return arrAIndex
	case 98:
		return arrBIndex
	case 99:
		return arrCIndex
	case 100:
		return arrDIndex
	case 101:
		return arrEIndex
	case 102:
		return arrFIndex
	case 103:
		return arrGIndex
	case 104:
		return arrHIndex
	case 105:
		return arrIIndex
	case 106:
		return arrJIndex
	case 107:
		return arrKIndex
	case 108:
		return arrLIndex
	case 109:
		return arrMIndex
	case 110:
		return arrNIndex
	case 111:
		return arrOIndex
	case 112:
		return arrPIndex
	case 113:
		return arrQIndex
	case 114:
		return arrRIndex
	case 115:
		return arrSIndex
	case 116:
		return arrTIndex
	case 117:
		return arrUIndex
	case 118:
		return arrVIndex
	case 119:
		return arrWIndex
	case 120:
		return arrXIndex
	case 121:
		return arrYIndex
	case 122:
		return arrZIndex
	}
	return arrLastIndex
}

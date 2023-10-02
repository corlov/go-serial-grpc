package main

func reverse(input []uint8) []uint8 {
    var output []uint8

    for i := len(input) - 1; i >= 0; i-- {
        output = append(output, input[i])
    }

    return output
}


func crc16(data []uint8) []uint8 {
	crc := uint16(0x0000)

	for i, _ := range data {
		a := 0
		temp := int((crc >> 8) << 8)
		bits := 0
		for bits < 8 {
			bits += 1
			if ((temp ^ a) & 0x8000 != 0) {
				a = (a << 1) ^ 0x1021
			} else {
				a <<= 1
			}
			temp <<= 1
		}
		crc = uint16(a) ^ (crc << 8) ^ uint16((data[i] & 0xFF))
	}

	// fmt.Println("crc16: ", crc)
	sequence := make([]uint8, 1)
	return intToSlice(crc & 0xFFFF, sequence)[:2]
}





func sliceToInt(s []uint8, storingDataType string) uint16 {
    res := uint16(0)
	shift := 8
    for i := 0; i < len(s); i++ {
		val := uint16(s[i])
		if (storingDataType == "BE") {
			res += uint16(val << uint8(shift*(i)))
		} else if (storingDataType == "LE") {
			res += uint16(val << uint8(shift*(len(s) - 1 - i)))
		} else {
			panic("sliceToInt, incorrect storingDataType value: " + storingDataType)
		}
		// fmt.Println("s[i]:", s[i], " i:", i, " shift: ", shift*(len(s) - 1 - i), " res:", uint16(val << uint16(shift*(len(s) - 1 - i))))
    }
	
    return res
}


func sliceToInt32(s []uint8, storingDataType string) uint32 {
    res := uint32(0)
	shift := 8
    for i := 0; i < len(s); i++ {
		val := uint32(s[i])
		if (storingDataType == "BE") {
			res += uint32(val << uint8(shift*(i)))
		} else if (storingDataType == "LE") {
			res += uint32(val << uint8(shift*(len(s) - 1 - i)))
		} else {
			panic("sliceToInt, incorrect storingDataType value: " + storingDataType)
		}
		// fmt.Println("s[i]:", s[i], " i:", i, " shift: ", shift*(len(s) - 1 - i), " res:", uint16(val << uint16(shift*(len(s) - 1 - i))))
    }
	
    return res
}


func intToSlice(n uint16, s []uint8) []uint8 {
    if n != 0 {
        i := uint8(n & 0xFF)
        s = append([]uint8{i}, s...)
        return intToSlice(n >> 8, s)
    }
    return s
}

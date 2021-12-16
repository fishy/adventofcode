package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	bitsPerByte = 8
	bitsPerHex  = 4

	signMask = 0b10000
)

var hexBitsMap = map[rune]bits{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

type bits []byte

func (b bits) String() string {
	var sb strings.Builder
	for i, bit := range b {
		sb.WriteString(fmt.Sprintf("%d", bit))
		if i < len(b)-1 && i%8 == 7 {
			sb.WriteString("_")
		}
	}
	return sb.String()
}

func (b *bits) next(n int) int {
	ret := (*b)[:n].toInt()
	*b = (*b)[n:]
	return ret
}

func (b *bits) readLiteral() int {
	var n int
	for {
		next := b.next(5)
		sign := next & signMask
		next = next - sign
		n <<= 4
		n += next
		if sign == 0 {
			return n
		}
	}
}

func (b bits) toInt() int {
	var n int
	for _, bit := range b {
		n <<= 1
		n += int(bit)
	}
	return n
}

func (b bits) available() int {
	return len(b)
}

func (b *bits) readSubPackets() (version int, values []int) {
	lt := b.next(1)
	switch lt {
	case 0:
		length := b.next(15)
		target := b.available() - length
		for b.available() > target {
			v, val := b.readPacket()
			version += v
			values = append(values, val)
		}
	case 1:
		length := b.next(11)
		for i := 0; i < length; i++ {
			v, val := b.readPacket()
			version += v
			values = append(values, val)
		}
	}
	return
}

func (b *bits) readPacket() (version int, value int) {
	version += b.next(3)
	operator := b.next(3)

	if operator == 4 {
		// literal
		return version, b.readLiteral()
	}

	v, values := b.readSubPackets()
	version += v
	switch operator {
	case 0:
		// sum
		for _, v := range values {
			value += v
		}
	case 1:
		// product
		value = 1
		for _, v := range values {
			value *= v
		}
	case 2:
		// min
		value = math.MaxInt
		for _, v := range values {
			if v < value {
				value = v
			}
		}
	case 3:
		// max
		for _, v := range values {
			if v > value {
				value = v
			}
		}
	case 5:
		// >
		if values[0] > values[1] {
			value = 1
		}
	case 6:
		// <
		if values[0] < values[1] {
			value = 1
		}
	case 7:
		// ==
		if values[0] == values[1] {
			value = 1
		}
	}
	return
}

func hexToBits(hex string) bits {
	b := make(bits, 0, len(hex)*bitsPerHex)
	for _, h := range hex {
		b = append(b, hexBitsMap[h]...)
	}
	return b
}

func main() {
	b := hexToBits(input)
	fmt.Println(b)

	fmt.Println(b.readPacket())
	fmt.Println("remaining:", b)
}

const input = `220D6448300428021F9EFE668D3F5FD6025165C00C602FC980B45002A40400B402548808A310028400C001B5CC00B10029C0096011C0003C55003C0028270025400C1002E4F19099F7600142C801098CD0761290021B19627C1D3007E33C4A8A640143CE85CB9D49144C134927100823275CC28D9C01234BD21F8144A6F90D1B2804F39B972B13D9D60939384FE29BA3B8803535E8DF04F33BC4AFCAFC9E4EE32600C4E2F4896CE079802D4012148DF5ACB9C8DF5ACB9CD821007874014B4ECE1A8FEF9D1BCC72A293A0E801C7C9CA36A5A9D6396F8FCC52D18E91E77DD9EB16649AA9EC9DA4F4600ACE7F90DFA30BA160066A200FC448EB05C401B8291F22A2002051D247856600949C3C73A009C8F0CA7FBCCF77F88B0000B905A3C1802B3F7990E8029375AC7DDE2DCA20C2C1004E4BE9F392D0E90073D31634C0090667FF8D9E667FF8D9F0C01693F8FE8024000844688FF0900010D8EB0923A9802903F80357100663DC2987C0008744F8B5138803739EB67223C00E4CC74BA46B0AD42C001DE8392C0B0DE4E8F660095006AA200EC198671A00010E87F08E184FCD7840289C1995749197295AC265B2BFC76811381880193C8EE36C324F95CA69C26D92364B66779D63EA071008C360098002191A637C7310062224108C3263A600A49334C19100A1A000864728BF0980010E8571EE188803D19A294477008A595A53BC841526BE313D6F88CE7E16A7AC60401A9E80273728D2CC53728D2CCD2AA2600A466A007CE680E5E79EFEB07360041A6B20D0F4C021982C966D9810993B9E9F3B1C7970C00B9577300526F52FCAB3DF87EC01296AFBC1F3BC9A6200109309240156CC41B38015796EABCB7540804B7C00B926BD6AC36B1338C4717E7D7A76378C85D8043F947C966593FD2BBBCB27710E57FDF6A686E00EC229B4C9247300528029393EC3BAA32C9F61DD51925AD9AB2B001F72B2EE464C0139580D680232FA129668`

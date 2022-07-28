package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func IsValidIp(addr string) bool {
	if addr == "" {
		return false
	}
	segs := 0
	chcnt := 0
	accum := 0

	for _, ch := range addr {
		if ch == '.' {
			if chcnt == 0 {
				return false
			}

			if segs++; segs == 4 {
				return false
			}
			chcnt = 0
			accum = 0
			continue
		}
		cn, _ := strconv.Atoi(string(ch))
		if (cn < 0) || (cn > 9) {
			return false
		}
		if accum = accum*10 + cn - 0; accum > 255 {
			return false
		}
		chcnt++
	}

	if segs != 3 || chcnt == 0 {
		return false
	}
	return true
}

func Tokenize(addr string) []uint {
	segments := strings.Split(addr, ".")
	tokens := make([]uint, len(segments))
	for index, t := range segments {
		i, _ := strconv.Atoi(t)
		tokens[index] = uint(i)
	}
	return tokens
}

func SimpleTransform(format string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i == 3 {
			result += fmt.Sprintf(format, tokens[i])
		} else {
			result += fmt.Sprintf(format+".", tokens[i])
		}
	}
	return result
}

func ConditionalTransform(cond int, format, fallbackFormat string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i >= cond {
			if i == 3 {
				result += fmt.Sprintf(format, tokens[i])
			} else {
				result += fmt.Sprintf(format+".", tokens[i])
			}
		} else {
			result += fmt.Sprintf(fallbackFormat+".", tokens[i])
		}
	}
	return result
}

func TransformLeftShift(shift int, format, fallbackFormat string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i < shift {
			result += fmt.Sprintf(format+".", tokens[i])
		} else {
			result += fmt.Sprintf(fallbackFormat, (tokens[2]<<8)|tokens[3])
			break
		}
	}
	return result
}

func ConditionalTransformLeftShift(shift int, format, fallbackFormat string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i >= shift {
			result += fmt.Sprintf(format+".", tokens[i])
		} else {
			result += fmt.Sprintf(fallbackFormat, (tokens[2]<<8)|tokens[3])
			break
		}
	}
	return result
}

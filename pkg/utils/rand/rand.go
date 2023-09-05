package rand

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 8, 9}
	l := len(numeric)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[r.Intn(l)])
	}
	return sb.String()
}

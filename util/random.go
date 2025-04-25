package util

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvxyz"

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func formatIntoPgTypeText(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  true,
	}
}

func RandomPgText(n int) pgtype.Text {
	return pgtype.Text{
		String: RandomString(n),
		Valid:  true,
	}
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomPhoneNumber() pgtype.Text {

	areaCode := rand.Intn(900) + 100  // 100-999
	prefix := rand.Intn(900) + 100    // 100-999
	lineNum := rand.Intn(9000) + 1000 // 1000-9999

	// Format the phone number
	phoneStr := fmt.Sprintf("%d-%d-%d", areaCode, prefix, lineNum)

	// Return as pgtype.Text with Valid set to true
	return pgtype.Text{
		String: phoneStr,
		Valid:  true,
	}
}

func RandomEmail(name string) string {
	return fmt.Sprintf("%s@email.com", name)
}

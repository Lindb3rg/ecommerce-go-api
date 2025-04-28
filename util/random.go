package util

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabetCapital = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int, capitalLetters bool) string {

	var sb strings.Builder
	if capitalLetters {

		k := len(alphabetCapital)
		for i := 0; i < n; i++ {
			c := alphabetCapital[rand.Intn(k)]
			sb.WriteByte(c)
		}
	} else {

		k := len(alphabet)
		for i := 0; i < n; i++ {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}

	}

	return sb.String()
}

func RandomPgTypeString(n int, capitalLetters bool) pgtype.Text {

	var sb strings.Builder
	if capitalLetters {

		k := len(alphabetCapital)
		for i := 0; i < n; i++ {
			c := alphabetCapital[rand.Intn(k)]
			sb.WriteByte(c)
		}
	} else {

		k := len(alphabet)
		for i := 0; i < n; i++ {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}

	}

	return pgtype.Text{
		String: sb.String(),
		Valid:  true,
	}
}

func FormatIntoPgTypeText(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  true,
	}
}

func GetRandomCountry() pgtype.Text {
	countries := []string{
		"Argentina", "Austria", "Belgium", "Brazil", "Canada",
		"Denmark", "Finland", "France", "Germany", "Ireland",
		"Italy", "Mexico", "Norway", "Poland", "Portugal",
		"Spain", "Sweden", "Switzerland", "UK", "USA", "Venezuela",
	}

	randomIndex := rand.Intn(len(countries))
	country := countries[randomIndex]

	return pgtype.Text{
		String: country,
		Valid:  true,
	}
}

func RandomPhoneNumber() pgtype.Text {

	areaCode := rand.Intn(900) + 100  // 100-999
	prefix := rand.Intn(900) + 100    // 100-999
	lineNum := rand.Intn(9000) + 1000 // 1000-9999

	// Format the phone number
	phone := fmt.Sprintf("%d-%d-%d", areaCode, prefix, lineNum)

	return pgtype.Text{
		String: phone,
		Valid:  true,
	}
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail(name string) string {
	return fmt.Sprintf("%s@email.com", name)
}

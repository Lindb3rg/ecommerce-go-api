package util

import (
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
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

// func RandomPgTypeString(n int, capitalLetters bool) pgtype.Text {

// 	var sb strings.Builder
// 	if capitalLetters {

// 		k := len(alphabetCapital)
// 		for i := 0; i < n; i++ {
// 			c := alphabetCapital[rand.Intn(k)]
// 			sb.WriteByte(c)
// 		}
// 	} else {

// 		k := len(alphabet)
// 		for i := 0; i < n; i++ {
// 			c := alphabet[rand.Intn(k)]
// 			sb.WriteByte(c)
// 		}

// 	}

// 	return pgtype.Text{
// 		String: sb.String(),
// 		Valid:  true,
// 	}
// }

func FormatIntoPgTypeText(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  true,
	}
}

func RandomContactName() pgtype.Text {
	name := gofakeit.Name()
	return pgtype.Text{
		String: name,
		Valid:  true,
	}

}

func RandomAddress() *gofakeit.AddressInfo {
	address := gofakeit.Address()
	return address

}

func RandomCompanyName() string {
	company := gofakeit.Company()
	return company
}

func RandomContactTitle() pgtype.Text {
	contactTitles := []string{
		"Accounting Manager",
		"Assistant Sales Agent",
		"Assistant Sales Representative",
		"Marketing Assistant",
		"Marketing Manager",
		"Order Administrator",
		"Owner",
		"Owner/Marketing Assistant",
		"Sales Agent",
		"Sales Associate",
		"Sales Manager",
		"Sales Representative",
	}

	randomIndex := rand.Intn(len(contactTitles))
	contactTitle := contactTitles[randomIndex]

	return pgtype.Text{
		String: contactTitle,
		Valid:  true,
	}
}

func RandomCountry() pgtype.Text {
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

func RandomRegion() pgtype.Text {
	regions := []string{
		"AK",
		"BC",
		"CA",
		"Co. Cork",
		"DF",
		"ID",
		"Isle of Wight",
		"Lara",
		"MT",
		"NM",
		"Nueva Esparta",
		"OR",
		"Québec",
		"RJ",
		"SP",
		"Táchira",
		"WA",
		"WY",
	}

	randomIndex := rand.Intn(len(regions))
	region := regions[randomIndex]

	return pgtype.Text{
		String: region,
		Valid:  true,
	}
}

func RandomPhoneNumber() pgtype.Text {

	phone := gofakeit.Phone()
	return pgtype.Text{
		String: phone,
		Valid:  true,
	}

}

func RandomEmail() pgtype.Text {
	email := gofakeit.Email()
	return pgtype.Text{
		String: email,
		Valid:  true,
	}
}

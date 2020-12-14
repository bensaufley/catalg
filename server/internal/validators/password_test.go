package validators_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bensaufley/catalg/server/internal/validators"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		it      string
		inputs  []string
		wantErr error
	}{
		{
			it: "rejects blank passwords",
			inputs: []string{""},
			wantErr: errors.New("password is blank"),
		},
		{
			it: "rejects passwords that are too short",
			inputs: []string{"asdf123", "a b c"},
			wantErr: errors.New("password is too short"),
		},
		{
			it: "rejects passwords that are too long",
			inputs: []string{
				"oaiweurpowjerpkeawrpoi hofpnaewporfn ewioprjewpior ewpioripe rpoaewpri uaewipru eawpi rpaew rpiaoieuroiewur pewrio uaer12\\3321[32 ",
				"1u2321o u02u 0u0u d-f uwa-iru a9023 84-9032 490ew9r8-9uf 90e-90wdf9-0ewuf9-0 wear90-8ewr-9ew!!11#21 iuiu padifj pwiofe iuewrio euwrewr  iawruioewu rioewuripoeawu rpioaewrioehwjfhewoifha pf",
			},
			wantErr: errors.New("password is too long"),
		},
		{
			it: "rejects passwords that contain no letters",
			inputs: []string{
				"123 41240218 12412 4!!![][]!",
				"[]|\\\" , ,./;/, .; 123 12312 3122131 \\[]\\[",
			},
			wantErr: errors.New("password contains no letters"),
		},
		{
			it: "rejects passwords that contain no symbols",
			inputs: []string{
				"fjpaiwejrpwirheopiraewnvoeanvewopifnaewoirne",
				"iewfuenvodnd",
				"hdfhnkafhoaewnfijeawnrpoaernpeonvoihdafdsklsfjndsfkjhcvanvownrewprneawroerkfmnxcvckvncxovnsironwekrnewraeruiyweorheroieerahfnbwo",
			},
			wantErr: errors.New("password contains no numbers or symbols"),
		},
		{
			it: "accepts valid passwords",
			inputs: []string{
				"correct horse battery staple",
				"short0987612",
				"very long sentence that has, if you will, a parenthetical clause, & just - barely - bumps up exactly against the character limit",
			},
		},
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range tests {
			testCase := tc
			for _, i := range testCase.inputs {
				input := i
				g.Run(testCase.it + ": " + input, func(test *testing.T) {
					test.Parallel()

					err := validators.Password(input)

					if testCase.wantErr == nil {
						assert.NoError(test, err)
					} else {
						assert.EqualError(test, err, testCase.wantErr.Error())
					}
				})
			}
		}
	})
}

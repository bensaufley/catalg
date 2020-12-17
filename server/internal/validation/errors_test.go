package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bensaufley/catalg/server/internal/validation"
)

func TestNewValidationError(t *testing.T) {
	type args struct {
		field  string
		errors []string
	}
	testCases := []struct {
		name string
		args args
		want *validation.Error
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc

			g.Run(tc.name, func(test *testing.T) {
				got := validation.NewError(tc.args.field, tc.args.errors...)

				assert.Equal(test, testCase.want, got)
			})
		}
	})
}

func TestCollectErrors(t *testing.T) {
	testCases := []struct {
		it      string
		errors  []*validation.Error
		wantErr *validation.Error
	}{}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				err := validation.CollectErrors(testCase.errors...)

				if testCase.wantErr == nil {
					assert.Nil(test, err)
				} else if assert.NotNil(test, err) {
					assert.EqualError(test, err, testCase.wantErr.Error())
					assert.Equal(test, testCase.wantErr.Errors(), err.Errors())
				}
			})
		}
	})
}

func TestValidationError_Error(t *testing.T) {
	type args struct {
		field  string
		errors []string
	}
	testCases := []struct {
		name    string
		args    args
		wantErr string
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(tc.name, func(test *testing.T) {
				err := validation.NewError(testCase.args.field, testCase.args.errors...)

				assert.EqualError(test, err, testCase.wantErr)
			})
		}
	})
}

func TestValidationError_Errors(t *testing.T) {
	type args struct {
		field  string
		errors []string
	}
	testCases := []struct {
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(tc.name, func(test *testing.T) {
				err := validation.NewError(testCase.args.field, testCase.args.errors...)

				assert.Equal(test, testCase.want, err.Errors())
			})
		}
	})
}

func TestValidationError_Merge(t *testing.T) {
	testCases := []struct {
		name string
		err1 *validation.Error
		err2 *validation.Error
		want *validation.Error
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc

			g.Run(tc.name, func(test *testing.T) {
				err := testCase.err1.Merge(testCase.err2)

				if testCase.want == nil {
					assert.Nil(test, err)
				} else if assert.NotNil(test, err) {
					assert.EqualError(test, err, testCase.want.Error())
					assert.Equal(test, testCase.want.Errors(), err.Errors())
				}
			})
		}
	})
}

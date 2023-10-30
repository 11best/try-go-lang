package calculator

import "errors"

// using `go mod init calculator` for made that be a local package
// https://stackoverflow.com/questions/17539407/how-to-import-local-packages-without-gopath/65055987#65055987?newreg=e83e2ebffc1d45598dc1ca09bb0e2ef7

// NAMING !!!
// in golang if naming using Capitalize(first letter of function name is upper) => public
// if lower => private (can use only it own file)

// have to handle error by yourself
func Divide(x float64, y float64) (float64, error) {
	if y == 0. {
		return 0., errors.New("division by zero")
	}
	return x / y, nil
}

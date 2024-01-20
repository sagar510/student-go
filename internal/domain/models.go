package domain

import "math/big"

type students struct {
	id           big.Int
	student_name string
	phone        string
	address      string

	user_id  big.Int
	admin_id big.Int
}

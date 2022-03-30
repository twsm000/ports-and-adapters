package api

import (
	"appstruct/internal/ports"
)

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:    db,
		arith: arith,
	}
}

func (adp Adapter) GetAdd(a int32, b int32) (int32, error) {
	result, err := adp.arith.Add(a, b)
	if err != nil {
		return 0, err
	}

	err = adp.db.AddHistory(result, "addition")
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (adp Adapter) GetSub(a int32, b int32) (int32, error) {
	result, err := adp.arith.Sub(a, b)
	if err != nil {
		return 0, err
	}

	err = adp.db.AddHistory(result, "subtraction")
	if err != nil {
		return 0, err
	}

	return result, nil
}
func (adp Adapter) GetMulti(a int32, b int32) (int32, error) {
	result, err := adp.arith.Multi(a, b)
	if err != nil {
		return 0, err
	}

	err = adp.db.AddHistory(result, "multiplication")
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (adp Adapter) GetDiv(a int32, b int32) (int32, error) {
	result, err := adp.arith.Div(a, b)
	if err != nil {
		return 0, err
	}

	err = adp.db.AddHistory(result, "division")
	if err != nil {
		return 0, err
	}

	return result, nil
}

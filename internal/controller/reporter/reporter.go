package reporter

import (
	"fmt"
)

type client interface {
	Generate(length uint) (string, error)
}

type parentheses interface {
	IsBalanced(str string) bool
}

type reporter struct {
	parentheses parentheses
	client      client
}

func New(p parentheses, client client) reporter {
	return reporter{parentheses: p, client: client}
}

func (r reporter) Do() error {
	lengths := [3]uint{2, 4, 8}
	for _, length := range lengths {
		if err := r.report(length); err != nil {
			return err
		}
	}
	return nil
}

func (r reporter) report(length uint) error {
	var balancedNum int
	for i := 0; i < 1000; i++ {
		resp, err := r.client.Generate(length)
		if err != nil {
			return err
		}
		if r.parentheses.IsBalanced(resp) {
			balancedNum++
		}
	}
	percentBalanced := float64(balancedNum) / 10
	fmt.Printf("For length %d, %f percent of generated strings are balanced\n", length, percentBalanced)
	return nil
}

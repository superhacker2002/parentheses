package reporter

import (
	"fmt"
	"github.com/superhacker2002/parentheses/internal/parentheses"
)

type client interface {
	Generate(length uint) (string, error)
}

type Reporter struct {
	client client
}

func New(client client) Reporter {
	return Reporter{client: client}
}

func (r Reporter) Do() error {
	lengths := [3]uint{2, 4, 8}
	for _, length := range lengths {
		if err := r.report(length); err != nil {
			return err
		}
	}
	return nil
}

func (r Reporter) report(length uint) error {
	var balancedNum int
	for i := 0; i < 1000; i++ {
		resp, err := r.client.Generate(length)
		if err != nil {
			return err
		}
		if parentheses.IsBalanced(resp) {
			balancedNum++
		}
	}
	percentBalanced := float64(balancedNum) / 10
	fmt.Printf("For length %d, %f percent of generated strings are balanced\n", length, percentBalanced)
	return nil
}

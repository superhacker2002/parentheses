package reporter

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

type client interface {
	Generate(length uint) (string, error)
}

type parenthesesChecker interface {
	IsBalanced(str string) bool
}

type reporter struct {
	parentheses parenthesesChecker
	client      client
}

func New(p parenthesesChecker, client client) reporter {
	return reporter{parentheses: p, client: client}
}

func (r reporter) Do() error {
	lengths := [3]uint{2, 4, 8}
	g := new(errgroup.Group)
	for _, length := range lengths {
		g.Go(func() error {
			return r.report(length)
		})
	}
	return g.Wait()
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

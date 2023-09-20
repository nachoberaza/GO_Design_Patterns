package concrete_strategies

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/strategy/internal/strategy"
)

type Fifo struct {
}

func (l *Fifo) Evict(_ *strategy.Cache) {
	fmt.Println("Evicting by fifo strategy")
}

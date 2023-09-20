package concrete_strategies

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/strategy/internal/strategy"
)

type Lru struct {
}

func (l *Lru) Evict(_ *strategy.Cache) {
	fmt.Println("Evicting by lru strategy")
}

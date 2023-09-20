package concrete_strategies

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/strategy/internal/strategy"
)

type Lfu struct {
}

func (l *Lfu) Evict(_ *strategy.Cache) {
	fmt.Println("Evicting by lfu strategy")
}

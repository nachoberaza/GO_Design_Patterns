package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/strategy/internal/concrete_strategies"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/strategy/internal/strategy"
)

func main() {
	lfu := &concrete_strategies.Lfu{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &concrete_strategies.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &concrete_strategies.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}

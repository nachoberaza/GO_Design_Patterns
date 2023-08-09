package main

import (
	"fmt"
	"github.com/baldurdevs/go_desing_patterns/patterns/creational/factory/internal/factories"
)

func main() {
	developerFactory := factories.NewEmployeeFactory("Developer", 6000)
	leaderFactory := factories.NewEmployeeFactory("Leader", 123000)
	bossFactory := factories.NewEmployeeFactory("CEO", 1000000)

	developer := developerFactory.Create("Nacho")
	fmt.Println(developer)

	leader := leaderFactory.Create("Fabi")
	fmt.Println(leader)

	boss := bossFactory.Create("Marquitos")
	fmt.Println(boss)
}

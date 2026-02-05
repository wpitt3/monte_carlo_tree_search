package mcts

type FlatPolicy struct{}

func (_ FlatPolicy) DefinePolicy(_ State[Action], actions []Action) []float32 {
	var policy []float32
	for i := 0; i < len(actions); i++ {
		policy = append(policy, 1.0)
	}
	return policy
}

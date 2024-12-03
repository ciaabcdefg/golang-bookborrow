package array

func Map[X, Y any](inputs []X, function func(X) Y) []Y {
	outputs := make([]Y, len(inputs))
	for i := range inputs {
		outputs[i] = function(inputs[i])
	}
	return outputs
}

func Filter[X any](inputs []X, predicate func(X) bool) []X {
	outputs := []X{}
	for i := range inputs {
		if predicate(inputs[i]) {
			outputs = append(outputs, inputs[i])
		}
	}
	return outputs
}

package main

import (
	"fmt"
	"os"
	"practice/input"
	"practice/optimizer"
)

func main() {
	files := []string{
		"a_an_example.in.txt",
		"b_basic.in.txt",
		"c_coarse.in.txt",
		"d_difficult.in.txt",
		"e_elaborate.in.txt",
	}

	for _, f := range files {
		process(f)
	}
}

func process(file string) {
	data, err := input.FromFile(file)
	if err != nil {
		panic(err)
	}

	res := optimizer.Solve(data)

	output := prettyPrint(res)

	_ = os.WriteFile(fmt.Sprintf("%v.output.txt", file), []byte(output), 0644)
}

func prettyPrint(res []string) string {
	result := ""

	result += fmt.Sprintf("%v", len(res))
	for _, v := range res {
		result += fmt.Sprintf(" %v", v)
	}

	return result
}

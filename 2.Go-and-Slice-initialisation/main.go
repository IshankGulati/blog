package main

import "strconv"

func MapAppend(in []int) []string {
	var out []string
	for _, s := range in {
		out = append(out, strconv.Itoa(s))
	}
	return out
}

func MapMakeAppend(in []int) []string {
	out := make([]string, 0, len(in))
	for _, s := range in {
		out = append(out, strconv.Itoa(s))
	}
	return out
}

func MapMakeLoop(in []int) []string {
	out := make([]string, len(in), len(in))
	for j, s := range in {
		out[j] = strconv.Itoa(s)
	}
	return out
}

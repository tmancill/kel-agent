package main

import "strings"

type sliceFlag []string

func (i *sliceFlag) String() string {
	return "my string representation"
}

func (i *sliceFlag) Set(value string) error {
	tokens := strings.Split(value, ",")
	for _, t := range tokens {
		*i = append(*i, t)
	}
	return nil
}

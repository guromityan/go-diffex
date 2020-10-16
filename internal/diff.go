package internal

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
)

func Diff(origin, target []Cell) {
	if diff := cmp.Diff(origin, target); diff != "" {
		for _, s := range strings.Split(diff, "\n") {
			if strings.HasPrefix(s, "+") {
				color.Green("%v", s)
				continue
			}

			if strings.HasPrefix(s, "-") {
				color.Red("%v", s)
				continue
			}

			trimed := strings.TrimSpace(s)
			if strings.HasPrefix(trimed, "{") && strings.HasSuffix(trimed, "},") {
				continue
			}

			fmt.Println(s)
		}
	}
}

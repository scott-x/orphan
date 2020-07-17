/*
* @Author: scottxiong
* @Date:   2020-07-06 18:43:21
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-07-06 19:46:14
 */
package orphan

import (
	_"fmt"
	"strings"
)

func Do(parents *[]string, children *[]string) []string {
	var result []string

	//initialization
	m := make(map[string]bool, 0)

	for _, p := range *parents {
		for _, c := range *children {
			if _, ok := m[c]; ok {
				continue
			} else {
				if strings.Contains(p, c) {
					m[c] = true
				}
			}
		}
	}

	//filter
	for _, y := range *children {
		if !m[y] {
			result = append(result, y)
		}
	}
	return result
}

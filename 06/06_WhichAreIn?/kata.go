/*
Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are substrings of strings of a2.

Example 1:
a1 = ["arp", "live", "strong"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns ["arp", "live", "strong"]

Example 2:
a1 = ["tarp", "mice", "bull"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns []

Notes:
Arrays are written in "general" notation. See "Your Test Cases" for examples in your language.
In Shell bash a1 and a2 are strings. The return is a string where words are separated by commas.
Beware: r must be without duplicates.
*/

package kata

import (
	"sort"
	"strings"
)

/* !!!!!!! WTF !!!!! This code:
Log
Expected
    <[]string | len:3, cap:8>: ["live", "arp", "strong"]
to equal
    <[]string | len:3, cap:3>: ["arp", "live", "strong"]
*/

func InArray(array1 []string, array2 []string) []string {
	tmp := []string{}
	for _, v := range array1 {
		for _, val := range array2 {
			switch {
			case strings.Contains(val, v):
				tmp = append(tmp, v)
			}
		}
	}

	for i := 0; i < len(tmp); i++ {
		if contains(tmp[i+1:], tmp[i]) {
			tmp = remove(tmp, i)
			i--
		}
	}

	sort.Strings(tmp)
	return tmp
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

/* Best practice:
package kata

import (
  "sort"
  "strings"
)

func InArray(array1 []string, array2 []string) []string {
  seen := make(map[string]struct{})
  result := []string{}
  for _, s1 := range array1 {
    if _, ok := seen[s1]; ok {
      continue
    }
    seen[s1] = struct{}{}
    for _, s2 := range array2 {
      if strings.Contains(s2, s1) {
        result = append(result, s1)
        break
      }
    }
  }
  sort.Strings(result)
  return result
}
*/

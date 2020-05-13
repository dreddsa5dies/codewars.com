/*
Wikipedia: The Baum–Sweet sequence is an infinite automatic sequence of 0s and 1s defined by the rule:

bn = 1 if the binary representation of n contains no block of consecutive 0s of odd length;
bn = 0 otherwise;

for n ≥ 0.

Define a generator function BaumSweet that sequentially generates the values of this sequence.

It will be tested for up to 1 000 000 values.

Note that the binary representation of 0 used here is not 0 but the empty string ( which does not contain any blocks of consecutive 0s ).
*/

package kata

import (
	"fmt"
	"strings"
)

func BaumSweet(ch chan<- int) {
	i := 0
	for {
		s := fmt.Sprintf("%b", i)
		arrS := strings.Split(s, "1")
		b := false
		for _, v := range arrS {
			if len(v)%2 != 0 {
				b = true
			}
		}
		if i == 0 {
			ch <- 1
		} else {
			if !b {
				ch <- 1
			} else {
				ch <- 0
			}
		}
		i++
	}
}

package kata

import "strings"

func Rot90Counter(s string) string {
	a := strings.Split(s, "\n")
	n := len(a)
	ret := []string{}
	for i := n - 1; i >= 0; i-- {
		line := ""
		for j := 0; j < n; j++ {
			line += string(a[j][i])
		}
		ret = append(ret, line)
	}
	return strings.Join(ret, "\n")
}

func Diag2Sym(s string) string {
	a := strings.Split(s, "\n")
	n := len(a)
	ret := []string{}
	for i := n - 1; i >= 0; i-- {
		line := ""
		for j := n - 1; j >= 0; j-- {
			line += string(a[j][i])
		}
		ret = append(ret, line)
	}
	return strings.Join(ret, "\n")
}

func SelfieDiag2Counterclock(s string) string {
	a1 := strings.Split(s, "\n")
	a2 := strings.Split(Diag2Sym(s), "\n")
	a3 := strings.Split(Rot90Counter(s), "\n")
	n := len(a1)
	ret := []string{}
	for i := 0; i < n; i++ {
		ret = append(ret, a1[i]+"|"+a2[i]+"|"+a3[i])
	}
	return strings.Join(ret, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

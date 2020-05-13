/*
Clients place orders to a stockbroker as strings. The order can be simple or multiple.

Type of a simple order: Quote/white-space/Quantity/white-space/Price/white-space/Status

where Quote is formed of non-whitespace character, Quantity is an int, Price a double (with mandatory decimal point "." ), Status is represented by the letter B (buy) or the letter S (sell).

Example:

"GOOG 300 542.0 B"

A multiple order is the concatenation of simple orders with a comma between each.

Example:

"ZNGA 1300 2.66 B, CLH15.NYM 50 56.32 B, OWW 1000 11.623 B, OGG 20 580.1 B"

or (C)

"ZNGA 1300 2.66 B,CLH15.NYM 50 56.32 B,OWW 1000 11.623 B,OGG 20 580.1 B"

To ease the stockbroker your task is to produce a string of type

"Buy: b Sell: s" where b and s are 'double' formatted with no decimal, b representing the total price of bought stocks and s the total price of sold stocks.

Example:

"Buy: 294990 Sell: 0"

Unfortunately sometimes clients make mistakes. When you find mistakes in orders, you must pinpoint these badly formed orders and produce a string of type:

"Buy: b Sell: s; Badly formed nb: badly-formed 1st simple order ;badly-formed nth simple order ;"

where nb is the number of badly formed simple orders, b representing the total price of bought stocks with correct simple order and s the total price of sold stocks with correct simple order.

Examples:

"Buy: 263 Sell: 11802; Badly formed 2: CLH16.NYM 50 56 S ;OWW 1000 11 S ;"

"Buy: 100 Sell: 56041; Badly formed 1: ZNGA 1300 2.66 ;"

Notes:

Due to Codewars whitespace differences will not always show up in test results.
With Golang use a format with "%.0f" for "Buy" and "Sell".
*/

package kata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func BalanceStatement(lst string) string {
	var (
		b, s     float64
		mistakes []string
	)

	str := strings.Split(lst, ", ")

	for _, v := range str {
		arr := strings.Split(v, " ")

		switch {
		case len(arr) != 4:
			mistakes = append(mistakes, v)
			continue
		case strings.ContainsAny(arr[1], "."):
			mistakes = append(mistakes, v)
			continue
		case !strings.ContainsAny(arr[2], "."):
			mistakes = append(mistakes, v)
			continue
		case arr[3] != "B" && arr[3] != "S":
			mistakes = append(mistakes, v)
			continue
		}

		f, err := strconv.ParseFloat(arr[2], 64)
		if err != nil {
			log.Fatalln(err)
		}

		i, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch {
		case arr[3] == "B":
			b += f * float64(i)
		case arr[3] == "S":
			s += f * float64(i)
		}
	}

	fin := fmt.Sprintf("Buy: %.0f Sell: %.0f", b, s)
	// ; Badly formed 1: CSCO 250.0 29 B ;
	if len(mistakes) != 0 && mistakes[0] != "" {
		fin += fmt.Sprintf("; Badly formed %v: ", len(mistakes))
		for _, v := range mistakes {
			fin += v + " ;"
		}
	}
	return fin
}

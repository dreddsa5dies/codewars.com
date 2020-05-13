/*
dataand data1 are two strings with rainfall records of a few cities for months from January to December. The records of towns are separated by \n. The name of each town is followed by :.

data and towns can be seen in "Your Test Cases:".

Task:
function: mean(town, strng) should return the average of rainfall for the city town and the strng data or data1 (In R and Julia this function is called avg).
function: variance(town, strng) should return the variance of rainfall for the city town and the strng data or data1.
Examples:
mean("London", data), 51.19(9999999999996) 
variance("London", data), 57.42(833333333374)
Notes:
if functions mean or variance have as parameter town a city which has no records return -1 or -1.0 (depending on the language)
Don't truncate or round: the tests will pass if abs(your_result - test_result) <= 1e-2 or abs((your_result - test_result) / test_result) <= 1e-6 depending on the language.
Shell tests only variance
A ref: http://www.mathsisfun.com/data/standard-deviation.html
data and data1 (can be named d0 and d1 depending on the language; see "Sample Tests:") are adapted from: http://www.worldclimate.com
*/

package kata

import("strings";"strconv")

func Mean(town string, strng string) float64 {
    fin := pr(town, strng)
    if len(fin) == 0 {
      return -1
    }
    
    total := 0.0

    for _, v := range fin {
      total += v
    }
    
    return total / float64(len(fin))
}
    
func Variance(town string, strng string) float64 {
    fin := pr(town, strng)
    if len(fin) == 0 {
      return -1
    }
    
    mean := Mean(town, strng)
    
    total := 0.0
    for _, number := range fin {
        total += (number-mean)*(number-mean)
    }
    return total / float64(len(fin))
}

func pr(town string, strng string) []float64 {
  m := map[string][]float64{}
  
  str := strings.Split(strng, "\n")
  
  for _, v := range str {
    x := strings.Split(v, ":")
    y := strings.Split(x[1], ",")
    for _, val := range y {
      t := strings.Split(val, " ")
      if s, err := strconv.ParseFloat(t[1], 64); err == nil {
        m[x[0]] = append(m[x[0]], s)
      }
    }
  }
  return m[town]
}
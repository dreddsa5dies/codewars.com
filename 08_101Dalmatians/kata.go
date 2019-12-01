/*
Your friend has been out shopping for puppies (what a time to be alive!)... He arrives back with multiple dogs, and you simply do not know how to respond!

By repairing the function provided, you will find out exactly how you should respond, depending on the number of dogs he has.

The number of dogs will always be a number and there will always be at least 1 dog.
*/

package kata

func HowManyDalmatians(number int) string {
	dogs := []string{"Hardly any",
		"More than a handful!",
		"Woah that's a lot of dogs!",
		"101 DALMATIONS!!!"}

	var str string
	if number <= 10 {
		str = dogs[0]
	} else if number <= 50 {
		str = dogs[1]
	} else if number <= 100 {
		str = dogs[2]
	} else if number == 101 {
		str = dogs[3]
	}
	return str
}

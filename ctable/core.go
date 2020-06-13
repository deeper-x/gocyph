package gtable

import "math/rand"

const minVal = 32
const maxVal = 255

// createTable create one map from two lists
func createTable(strList []string, intList []int) map[int]string {
	var retMap = map[int]string{}

	for k, v := range intList {
		retMap[v] = strList[k]
	}

	return retMap
}

// getASCIIelems return all ASCII symbols used in input fields
func getASCIIelems() []string {
	var retVal []string

	// consider only useful symbols
	for i := minVal; i <= maxVal; i++ {
		retVal = append(retVal, string(i))
	}

	return retVal
}

// getRandomInt generate random unique integer
func getRandomInt() int {
	return rand.Intn(100000-1) + 1
}

// alreadyIn check if elements is already created
func alreadyIn(inVal int, randomInts []int) bool {
	for _, v := range randomInts {
		if v == inVal {
			return true
		}
	}
	return false
}

// newRandomInts build a list of unique random integers
func newRandomInts() []int {
	var randomInts []int

	size := 0

	for size <= (maxVal - minVal) {
		newInt := getRandomInt()

		if !alreadyIn(newInt, randomInts) {
			randomInts = append(randomInts, newInt)
			size++
		}
	}

	return randomInts
}

// New create the chiperTable
func New() map[int]string {
	ASCIIlist := getASCIIelems()
	intsList := newRandomInts()

	return createTable(ASCIIlist, intsList)
}

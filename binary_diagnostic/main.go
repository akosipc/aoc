package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readLines("data.txt")

	if err != nil {
		fmt.Println(err)
	}

	gammaBits, epsilonBits := []string{}, []string{}

	for x := range input[0] {
		gamma, epsilon := calcGammaEpsilon(buildBits(input, x))

		gammaBits = append(gammaBits, gamma)
		epsilonBits = append(epsilonBits, epsilon)
	}
	gamma := convertStringBinaryToInt(gammaBits)
	epsilon := convertStringBinaryToInt(epsilonBits)

	fmt.Println(gamma * epsilon)

	oxy, car := bitCriteria(input)

	oxygen := convertStringBinaryToInt(oxy)
	carbon := convertStringBinaryToInt(car)

	fmt.Println(oxygen * carbon)

}

func bitCriteria(array []string) ([]string, []string) {
	i, y := 0, 0
	oxygenList, carbonList := array, array

	for i < 12 {
		oxygen, _ := calcGammaEpsilon(buildBits(oxygenList, i))

		if oxygen == "" {
		} else {
			oxygenList = filterByBitAndPosition(oxygenList, oxygen, i)
		}

		if len(oxygenList) == 1 {
			i = 12
		} else {
			i++
		}
	}

	for y < 12 {
		_, carbon := calcGammaEpsilon(buildBits(carbonList, y))

		if carbon == "" {
		} else {
			carbonList = filterByBitAndPosition(carbonList, carbon, y)
		}

		if len(carbonList) == 1 {
			y = 12
		} else {
			y++
		}
	}

	return oxygenList, carbonList
}

func filterByBitAndPosition(array []string, bit string, position int) []string {
	filtered := []string{}

	for i := range array {
		if fmt.Sprintf("%c", array[i][position]) == bit {
			filtered = append(filtered, array[i])
		}
	}

	return filtered
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func buildBits(array []string, index int) []string {
	var byteArray []string

	for _, element := range array {
		byteArray = append(byteArray, fmt.Sprintf("%c", element[index]))
	}

	return byteArray
}

func calcGammaEpsilon(array []string) (string, string) {
	zeroes, ones := []string{}, []string{}

	for i := range array {
		if array[i] == "1" {
			ones = append(ones, array[i])
		} else {
			zeroes = append(zeroes, array[i])
		}
	}

	if len(zeroes) > len(ones) {
		return "0", "1"
	} else if len(zeroes) < len(ones) {
		return "1", "0"
	} else {
		return "1", "0"
	}
}

func convertStringBinaryToInt(str []string) int64 {
	i, err := strconv.ParseInt(strings.Join(str, ""), 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	return i
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T := toInt(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		A := splitToFloat64(scanner.Text())
		scanner.Scan()
		P := splitToFloat64(scanner.Text())
		calcWairo(A, P)
	}
}

func calcWairo(A, P []float64) {
	var total float64
	var sum float64
	for i, v := range A {
		ii := (float64)(i + 1)
		total += ii * v
		sum += v
	}
	avg := total / sum
	if avg >= 3 {
		fmt.Printf("0")
		return

	}

	// 賄賂計算
	cost_map := make([]float64, 0)

	for _, ii := range []float64{4, 5} {
		// ii 点のレビューを何件買うか
		p := P[(int)(ii-1)]
		var num_buy int
		for {
			num_buy += 1
			tmp_avg := (total + ii*float64(num_buy)) / (sum + float64(num_buy))
			// fmt.Printf("ii: %f, point %f, total: %f, tmp_avg: %f, num_buy: %d, avg: %f\n", ii, p, total, tmp_avg, num_buy, avg)
			if tmp_avg >= 3 {
				cost_map = append(cost_map, p*float64(num_buy))
				break
			}
		}
	}

	fmt.Printf("%d\n", (int)(ArrayMin(cost_map)))
}

func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func toFloat64(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

func splitToInt(s string) []int {
	str_arr := strings.Split(s, " ")
	int_arr := make([]int, len(str_arr))
	for i, v := range str_arr {
		int_arr[i] = toInt(v)
	}

	return int_arr
}

func splitToFloat64(s string) []float64 {
	str_arr := strings.Split(s, " ")
	arr := make([]float64, len(str_arr))
	for i, v := range str_arr {
		arr[i] = toFloat64(v)
	}

	return arr
}

func ArrayMin(arr []float64) float64 {
	min := arr[0]
	for _, v := range arr {
		min = math.Min(min, v)
	}

	return min
}

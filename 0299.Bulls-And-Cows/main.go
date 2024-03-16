package main

import "fmt"

func main() {
	testCases := []struct {
		secret string
		guess  string
		result string
	}{
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
		{"1234", "5678", "0A0B"},
		{"1122", "2211", "0A4B"},
		{"1122", "1222", "3A0B"},
	}

	for _, tc := range testCases {
		output := getHint(tc.secret, tc.guess)
		if output != tc.result {
			fmt.Printf("FAIL: For secret=%s, guess=%s, expected %s but got %s\n。", tc.secret, tc.guess, tc.result, output)
		}
	}
}

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	dpSecret := make([]int, 10)
	dpGuess := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		// to caculate bulls
		if secret[i] == guess[i] {
			bulls++
			continue
		}
		dpSecret[int(secret[i]-'0')]++
		dpGuess[int(guess[i])-'0']++
	}

	for i := 0; i < 10; i++ {
		// choose the minimum one
		cows += min(dpSecret[i], dpGuess[i])
	}

	return fmt.Sprint(bulls, "A", cows, "B")
}

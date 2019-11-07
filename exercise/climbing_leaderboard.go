package exercise

func ClimbingLeaderBoard(scores []int, personScores []int) []int {
	var leaderBoard, scoreGroup []int
	ranks := make(map[int]int)

	highestScore := scores[0]
	lowestScore := scores[len(scores)-1]
	scoreGroup = append(scoreGroup, highestScore)

	ranks[highestScore] = 1
	for i := 1; i < len(scores); i++ {
		if scores[i] < scores[i-1] {
			ranks[scores[i]] = ranks[scores[i-1]] + 1
			scoreGroup = append(scoreGroup, scores[i])
		}
	}

	for _, score := range personScores {
		if score >= highestScore {
			leaderBoard = append(leaderBoard, 1)
			continue
		}

		if score < lowestScore {
			leaderBoard = append(leaderBoard, len(ranks)+1)
			continue
		}

		if score == lowestScore {
			leaderBoard = append(leaderBoard, len(ranks))
			continue
		}

		if ranks[score] != 0 {
			leaderBoard = append(leaderBoard, ranks[score])
			continue
		}

		flag := binarySearch(scoreGroup, score)
		leaderBoard = append(leaderBoard, ranks[flag])
	}

	return leaderBoard
}

func binarySearch(arr []int, x int) int {
	var flag int
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < x {
			high = median - 1
		} else {
			low = median + 1
		}

		flag = median
	}

	return arr[flag]
}

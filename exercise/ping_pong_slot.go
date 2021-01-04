package exercise

/**
The engineer love playing table tennis and there is only one table available.
Given the engineer's availability to play table tennis, we would like to determine the number of hours where they will be playing singles / doubles.
Condition:
The working hours are 9am to 6pm.
There is only one table.
When there are two to three people, they can only play singles.
When there are  over four people, they can play doubles.
Input:
The first line of the input givens the number of players T, T players available time follow.
Each consists of one line with Integer S E. S represents start time, and E represents end time.
Ex:
2
9 10
9 12
From Mohamed Siddiq to Everyone: (19:12)
 It means 2 players can play ping pong. The 1st player can play from 9am  to 10am and the 2nd player can play ping pong from 9am to 12pm.

Output:
The output would be a one liner containing x y. where x is singles hours count and y is double hours.
The above input means player 1 can play ping pong from 9am to 10am, and player 2 can play ping pong from 9am to 12pm.

The output of above example would be
1 0
*/

type PingPongSlot struct {
	Start int
	End   int
}

type PingPongTime struct {
	Min int
	Max int
}

func PingPongMatch(time PingPongTime, avail []PingPongSlot) (int, int) {
	slot := make(map[int]int)
	for i := time.Min; i < time.Max; i++ {
		slot[i] = 0
	}

	for _, p := range avail {
		for x := p.Start; x < p.End; x++ {
			slot[x] = slot[x] + 1
		}
	}

	single, double := 0, 0
	for _, s := range slot {
		if s > 1 && s < 4 {
			single++
		}
		if s > 3 {
			double++
		}
	}

	return single, double
}

// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
//
// The trick with this problem is solving it in a performant way.
//
// Need to sort the user's input along with the scores and iterate through both arrays
// at the same time.
//
// The loop goes through the player scores and iterates through the high score list
// until the player's score fits.  Once we've exhausted the high scores, the player is
// always the last rank + 1
//
// The user's score and rank is stored in a map.  At the end of the function the
// original player list is evaluated against the map to return the value.
//
// While it's true that I could have iterated the player scores in reverse the problem
// doesn't say that scores are sorted.  Thus I don't assume they are prior to the method
// being run.
//


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

// Scores represents a list of player scores
type Scores struct {
	scores []int32
}

func NewScores(scores []int32) *Scores {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	return &Scores{scores: scores}
}

func (ranking *Scores) Get(idx int) int32 {
	return ranking.scores[idx]
}

func (ranking *Scores) HasIdx(idx int) bool {
	return idx < len(ranking.scores)
}


// DenseRankingIter iterates through a sorted score list and returns the rankings
type DenseRankingIter struct {
	ranking *Scores

	index int
	rank int32
	score int32
}

// NewDenseRankingIter creates a new Dense Ranking iterator for the scores
func (ranking *Scores) NewDenseRankingIter() *DenseRankingIter {
	return &DenseRankingIter{ranking: ranking, index: -1, rank: 0, score: 0 }
}

// Next ranking from the iter
func (iter *DenseRankingIter) Next() (hasMore bool, rank int32, score int32) {

	if iter.ranking.HasIdx( iter.index + 1 ) {
		iter.index++
		score := iter.ranking.Get(iter.index)
		if score != iter.score {
			iter.rank++
		}
		iter.score = score
	}

	hasMore = iter.ranking.HasIdx( iter.index + 1 )
	rank = iter.rank
	score = iter.score
	return
}


// Main function that solves the problem
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Sort the ranked from highest to lowest
	// Problem is vague whether these are always sorted
	// Although in practice from the test cases it appears they are
	//
	ranking := NewScores(ranked)
	rankIter := ranking.NewDenseRankingIter()

	sortedPlayer := make([]int32, len(player))
	copy(sortedPlayer, player)

	// Sort the player ranks from highest to lowest as well
	sort.Slice(sortedPlayer, func(i, j int) bool {
		return sortedPlayer[i] > sortedPlayer[j]
	})

	// Create an array with the player ranks to return
	playerRanks := make(map[int32]int32)
	hasMoreRanks, rank, rankScore := rankIter.Next()

	// Loop through all the player ranking
	for _, playerScore := range sortedPlayer {

		// Loop while the player score is less than the current ranking score
		for playerScore < rankScore && hasMoreRanks {
			hasMoreRanks, rank, rankScore = rankIter.Next()
		}

		// If the player score is greater than or equal to the current rank score
		// Then it should have replace that rank or be in that rank
		if playerScore >= rankScore {
			playerRanks[playerScore] = rank
			continue
		}

		// No more ranks, player score is last
		if !hasMoreRanks {
			playerRanks[playerScore] = rank + 1
		}
	}

	// Prepare the return value based on the scores
	retValue := make([]int32, len(player))
	for idx, playerScore := range player {
		retValue[idx] = playerRanks[playerScore]
	}
	return retValue
}

// This stuff was provided as scaffolding by HackerRank
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	writer := bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
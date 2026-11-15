package service

import "testing"

func TestGetLeaderboardIncludesAllPlayers(t *testing.T) {
	game := Game{
		Players: []*Player{
			{Name: "Second", Points: 30},
			{Name: "Fourth", Points: 10},
			{Name: "First", Points: 40},
			{Name: "Third", Points: 20},
		},
	}

	leaderboard := game.getLeaderboard()
	want := []LeaderboardEntry{
		{Name: "First", Points: 40},
		{Name: "Second", Points: 30},
		{Name: "Third", Points: 20},
		{Name: "Fourth", Points: 10},
	}
	if len(leaderboard) != len(want) {
		t.Fatalf("got %d leaderboard entries, want %d", len(leaderboard), len(want))
	}

	for i, entry := range leaderboard {
		if entry != want[i] {
			t.Errorf("entry %d = {%q, %d}, want {%q, %d}", i, entry.Name, entry.Points, want[i].Name, want[i].Points)
		}
	}
}

func TestGetPointsRewardCountsOnlyCorrectAnswers(t *testing.T) {
	game := Game{
		Time: 0,
		Players: []*Player{
			{Answered: true, AnsweredCorrectly: true},
			{Answered: true, AnsweredCorrectly: false},
		},
	}

	if got, want := game.getPointsReward(), 800; got != want {
		t.Errorf("getPointsReward() = %d, want %d", got, want)
	}
}

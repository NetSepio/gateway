package OperatorEventActivities

import (
	"time"

	"github.com/NetSepio/gateway/util/pkg/logwrapper"
)

func StartLeaderboardUpdateScheduler() {
	go func() {
		for {
			if err := UpdateLeaderboardFromIndexer(); err != nil {
				logwrapper.Error("Failed to update leaderboard from indexer", err)
			}
			time.Sleep(1 * time.Hour) // Update every hour
		}
	}()
}

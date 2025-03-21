package main

import (
	"time"

	"github.com/potatozhb/bookingapp/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `json:"apikey"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Apikey:    dbUser.Apikey,
	}
}

type Feed struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
}

func dbFeedToFeed(dbUser database.Feed) Feed {
	return Feed{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Url:       dbUser.Url,
		UserID:    dbUser.UserID,
	}
}

func dbFeedsToFeeds(feeds []database.Feed) []Feed {
	var result []Feed
	for _, dbfeed := range feeds {
		result = append(result, dbFeedToFeed(dbfeed))
	}
	return result
}

package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nishantxoxo/rssag/internal/database"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	Name      string		`json:"name"`
	APIkey 		string 		`json:"api_key"`
}


func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:  dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIkey: dbUser.ApiKey,
	}
}




type Feed struct {
	ID        uuid.UUID		`json:"id"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	Name      string		`json:"name"`
	Url       string		`json:"url"`
	UserID    uuid.UUID		`json:"user_id"`
}


func databaseFeedToFeed(dbUser database.Feed) Feed {
	return Feed{
		ID:  dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		Url: dbUser.Url,
		UserID: dbUser.UserID,
	}
}


func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed {}

	for _, dbFeed := range dbFeed {
		feeds = append(feeds, databaseFeedToFeed(dbFeed) )
	}

	return feeds
}
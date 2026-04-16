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


type FeedFollows struct {
	ID        uuid.UUID		`json:"id"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`

	UserID    uuid.UUID		`json:"user_id"`
	FeedID    uuid.UUID		`json:"feed_id"`

}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow ) FeedFollows {
	return FeedFollows{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,

	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollows {
	feeds := []FeedFollows {}

	for _, dbFeed := range dbFeedFollows {
		feeds = append(feeds, databaseFeedFollowToFeedFollow(dbFeed) )
	}

	return feeds
}



type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time 	`json:"updated_at"`
	Title       string		`json:"title"`
	Description *string	`json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbp database.Post ) Post {

	var descc *string 
	if dbp.Description.Valid {
		descc = &dbp.Description.String
	}

	return Post{
		ID: dbp.ID,
		CreatedAt: dbp.CreatedAt,
		UpdatedAt: dbp.UpdatedAt,
		Title: dbp.Title,
		Description: descc,
		PublishedAt: dbp.PublishedAt,
		Url: dbp.Url,
		FeedID: dbp.FeedID,
	}
}


func databasePostsToPosts(dbPosts []database.Post) []Post {
	Posts := []Post {}

	for _, dbpost := range dbPosts {
		Posts = append(Posts, databasePostToPost(dbpost) )
	}

	return Posts
}
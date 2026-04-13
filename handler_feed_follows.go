package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/nishantxoxo/rssag/internal/database"
	// "structs"
)

func (apiCfg *apiConfig )handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	 
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
		// URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:=	decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing josn: %v", err))
		return
	}
	

	FeedFollow, err :=	apiCfg.DB.CreatefeedFollow(r.Context(), database.CreatefeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC() ,
		UpdatedAt: time.Now().UTC(),
		
		UserID: user.ID,
		FeedID: params.FeedID,
		
	})

	if err!=nil {
				respondWithError(w, 400, fmt.Sprintf("couldnt create feed follow: %v", err))
		return
	}

	// respondWithJSON(w, 200)1

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(FeedFollow))
} 


func (apiCfg *apiConfig )handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	 

	FeedFollow, err :=	apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err!=nil {
				respondWithError(w, 400, fmt.Sprintf("couldnt get feed follows: %v", err))
		return
	}

	// respondWithJSON(w, 200)1

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(FeedFollow))
}



func (apiCfg *apiConfig )handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	 
	feedFollowIDstr:=	chi.URLParam(r,"feedFollowID")
	feedFollowID , err := uuid.Parse(feedFollowIDstr)
	if err !=nil {
		respondWithError(w, 400, fmt.Sprintf("couldnt parse feed follows ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})	

		if err !=nil {
		respondWithError(w, 400, fmt.Sprintf("couldnt delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}


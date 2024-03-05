package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"restapi/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

// @Summary UserVote
// @Security ApiKeyAuth
// @Tags votes
// @Description user vote
// @Accept  json
// @Produce  json
// @Param input body domain.UserRating true "vote info"
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/posts/:id/vote [post]
func (h *Handler) UserVote(c *gin.Context) {
	var input domain.UserRating

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postID, err := getID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, ok := c.Get("id")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, roleDoesNotExist)
		return
	}

	userID := id.(int)

	currentTime := time.Now().UTC()

	postUserID, err := h.service.PostService.GetUserID(postID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if postUserID == userID {
		NewErrorResponse(c, http.StatusInternalServerError, errSelfVote.Error())
		return
	}

	rating, err := h.service.RatingService.GetVotes(userID, postID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	postVote := domain.VoteSumUpdate{
		Vote:   input.Vote,
		PostID: postID,
	}

	if rating.FromUserID == userID && rating.ToPostID == postID {
		if !currentTime.After(rating.WhenVoted.Add(1 * time.Second)) {
			NewErrorResponse(c, http.StatusInternalServerError, errCannotVoteYet.Error())
			return
		}

		if input.Vote == rating.UserVote {
			postVote = domain.VoteSumUpdate{
				Vote:   -input.Vote,
				PostID: postID,
			}
		}

		err = h.service.RatingService.UpdateVote(input, rating)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = h.service.VoteService.UpdatePostVotes(postVote)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"errResponse": err,
			"message":     "user vote updated",
		})

		return
	}

	err = h.service.RatingService.InsertVote(userID, postID, input.Vote)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.VoteService.UpdatePostVotes(postVote)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"errResponse": err,
		"message":     "user voted",
	})

}

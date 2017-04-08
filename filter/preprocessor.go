package filter

import "../models"

func Preprocess(post *models.Post) *models.ProcessedPost {
	processed := new(models.ProcessedPost)
	processed.PostLength = uint(len(post.Raw_body_text))
	processed.LikeCount = post.Like_count
	processed.FollowersCount = post.Author_followers_count
	processed.FriendCount = post.Author_friends_count
	processed.HashtagCount = uint(len(post.Hashtags))
	processed.RetweetCount = post.Retweet_count
	processed.IsRetweet = post.Is_reshare
	processed.KloutScore = post.Author_klout_score
	exclamationCount := 0
	for _, c := range post.Raw_body_text {
		character := string(c)
		if character == "!" {
			exclamationCount++
		}
	}
	processed.ExclaimationCount = uint(exclamationCount)

	return processed
}
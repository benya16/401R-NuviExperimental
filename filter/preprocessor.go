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

	//take the raw body text, run it through the filter

	dangerWords := Dictionary.isDangerousSentance(post.Cleaned_body_text)

	processed.Shooter = dangerWords.Contains("shooter")
	processed.ActiveShooter = dangerWords.Contains("active shooter")
	processed.Attack = dangerWords.Contains("attack")
	processed.Bomb = dangerWords.Contains("bomb")
	processed.BombThreat = dangerWords.Contains("bomb threat")
	processed.Breaking = dangerWords.Contains("breaking")
	processed.Danger = dangerWords.Contains("danger")
	processed.Dead = dangerWords.Contains("dead")
	processed.Gunman = dangerWords.Contains("gunman")
	processed.Killing = dangerWords.Contains("killing")
	processed.Rape = dangerWords.Contains("rape")
	processed.Shooting = dangerWords.Contains("shooting")
	processed.Stabbing = dangerWords.Contains("stabbing")
	processed.Terrorist = dangerWords.Contains("terrorist")
	processed.Warning = dangerWords.Contains("warning")


	return processed
}
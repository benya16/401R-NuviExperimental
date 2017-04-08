package models

type Post struct {
	Activity_url string
	Author_favorites_count uint
	Author_followers_count uint
	Author_friends_count uint
	Author_klout_score uint
	Author_picture_url string
	Author_posts_count uint
	Author_profile_url string
	Author_real_name string
	Author_username string
	Bio string
	Cleaned_body_text string
	Country string
	Country_code string
	Embedded_urls []string
	Hashtags []string
	Historical_search  bool
	Is_reshare bool
	Language string
	Latitude float64
	Like_count uint
	Location_display_name string
	Longitude float64
	Mentions  []string
	Meta_data []interface{}
	Network string
	Normalized_urls []string
	Parent struct{
		Parent_author   string
		Parent_author_profile_url string
		Parent_author_reach int
		Parent_body_text string
		Parent_created_at string
		Parent_social_source_id string
	}
	Post_created_at string
	Post_media  struct {
		Media_url string
		Url string
		Display_url string
		Media_type string
		Video_url string
	}
	Raw_body_text string
	Region string
	Region_code string
	Retweet_count uint
	Social_monitor_sources []struct {
		Company_uid string
		Monitor_uid string
		Keywords map[string]uint
	}
	Social_monitor_uids []string
	Social_source_uid string
	Source string
	Topic_monitor_id string
}

type ProcessedPost struct {
	PostLength        uint
	LikeCount         uint
	FollowersCount    uint
	FriendCount       uint
	HashtagCount      uint
	RetweetCount      uint
	IsRetweet         bool
	KloutScore        uint
	ExclaimationCount uint
	Shooting	bool
	Shooter		bool
	ActiveShooter	bool
	Gunman		bool
	Warning		bool
	Danger		bool
	Breaking	bool
	BombThreat	bool
	Killing		bool
	Dead		bool
	Stabbing	bool
	Attack		bool
	Attacker	bool
	Terrorist	bool
	Bomb		bool
	Rape		bool
}

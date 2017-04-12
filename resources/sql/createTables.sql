create table post
(
  UUID varchar(128) NOT NULL PRIMARY KEY,
  COLLECTED timestamp DEFAULT now(),
  threat boolean,
  POST bytea
);

create table processed_post
(
  UUID varchar(128) REFERENCES post(UUID) ON DELETE CASCADE NOT NULL,
  post_length int,
  like_count int,
  followers_count int,
  friend_count int,
  hashtag_count int,
  retweet_count int,
  is_retweet BOOLEAN,
  klout_score int,
  exclaimation_count int,
  active_shooter BOOLEAN DEFAULT FALSE,
  attack BOOLEAN DEFAULT FALSE,
  attacker BOOLEAN DEFAULT FALSE,
  bomb BOOLEAN DEFAULT FALSE,
  bomb_threat BOOLEAN DEFAULT FALSE,
  breaking BOOLEAN DEFAULT FALSE,
  danger BOOLEAN DEFAULT FALSE,
  dead BOOLEAN DEFAULT FALSE,
  gunman BOOLEAN DEFAULT FALSE,
  killing BOOLEAN DEFAULT FALSE,
  rape BOOLEAN DEFAULT FALSE,
  shooter BOOLEAN DEFAULT FALSE,
  shooting BOOLEAN DEFAULT FALSE,
  stabbing BOOLEAN DEFAULT FALSE,
  terrorist BOOLEAN DEFAULT FALSE,
  warning BOOLEAN DEFAULT FALSE,
  class_threat BOOLEAN
);
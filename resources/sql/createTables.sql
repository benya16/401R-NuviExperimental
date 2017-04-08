create table post
(
  UUID varchar(128) DEFAULT uuid_generate_v1() PRIMARY KEY,
  COLLECTED timestamp DEFAULT now(),
  threat boolean,
  POST bytea
);

create table processed_post
(
  post_length int,
  like_count int,
  followers_count int,
  friend_count int,
  hashtag_count int,
  retweet_count int,
  is_retweet BOOLEAN,
  klout_score int,
  exclaimation_count int
);
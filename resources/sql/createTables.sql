create table post
(
  UUID varchar(128) DEFAULT uuid_generate_v1() PRIMARY KEY,
  COLLECTED timestamp DEFAULT now(),
  threat boolean,
  POST bytea
)
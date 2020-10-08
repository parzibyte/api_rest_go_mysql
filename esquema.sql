CREATE TABLE video_games
(
    id    bigint unsigned not null primary key auto_increment,
    name  VARCHAR(255)    NOT NULL,
    genre VARCHAR(255)    NOT NULL,
    year  INTEGER         NOT NULL
);

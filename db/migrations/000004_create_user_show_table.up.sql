CREATE TABLE IF NOT EXISTS user_shows (
	user_id bigint NOT NULL,
	show_id bigint NOT NULL
);

CREATE INDEX user_shows_idx ON user_shows (user_id, show_id);

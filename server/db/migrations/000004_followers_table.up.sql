CREATE TABLE IF NOT EXISTS followers(
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    CONSTRAINT fk_followers_follower_id FOREIGN KEY(follower_id) REFERENCES users(id),
    CONSTRAINT fk_followers_followee_id FOREIGN key(followee_id) REFERENCES users(id)
);
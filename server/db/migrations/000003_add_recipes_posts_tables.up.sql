CREATE TABLE IF NOT EXISTS recipes(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    instructions TEXT
);

CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    recipe_id INTEGER NOT NULL,
    description VARCHAR(255),
    CONSTRAINT fk_posts_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_posts_recipe_id FOREIGN KEY(recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS recipe_photos(
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    post_id INTEGER,
    photo_link VARCHAR(100),
    CONSTRAINT fk_recipe_photos_rid FOREIGN KEY(recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    CONSTRAINT fk_recipe_photos_pid FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);
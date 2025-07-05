ALTER TABLE users ADD subscribe BOOLEAN DEFAULT TRUE;
CREATE INDEX user_subscribe ON users (subscribe);
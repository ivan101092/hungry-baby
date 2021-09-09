
CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS files (
	id SERIAL PRIMARY KEY,
	type TEXT NOT NULL CHECK (char_length(type) <= 50),
	url TEXT NOT NULL CHECK (char_length(url) <= 255),
	user_upload TEXT NOT NULL CHECK (char_length(user_upload) <= 100),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER files BEFORE UPDATE ON files FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO files ("id", "type", "url", "user_upload", "deleted_at")
VALUES (0, '', '', '', now());

CREATE TABLE IF NOT EXISTS countries (
	id SERIAL PRIMARY KEY,
	country_code TEXT NOT NULL CHECK (char_length(country_code) <= 10),
  	name TEXT NOT NULL CHECK (char_length(name) <= 100),
	status BOOLEAN NOT NULL DEFAULT 'false',
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER countries BEFORE UPDATE ON countries FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO countries ("id", "country_code", "name", "deleted_at")
VALUES (0, '', '', now());
INSERT INTO countries (country_code, name, status, created_at, updated_at)
VALUES ('62', 'Indonesia', true, current_timestamp, current_timestamp);
ALTER SEQUENCE countries_id_seq RESTART WITH 2;

CREATE TABLE IF NOT EXISTS provinces (
	id SERIAL PRIMARY KEY,
	country_id INTEGER NOT NULL REFERENCES countries(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 50),
	name TEXT NOT NULL CHECK (char_length(name) <= 50),
	status BOOLEAN NOT NULL DEFAULT 'false',
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER provinces BEFORE UPDATE ON provinces FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO provinces ("id", "country_id", "code", "name", "deleted_at")
VALUES (0, 0, '', '', now());

CREATE TABLE IF NOT EXISTS cities (
	id SERIAL PRIMARY KEY,
	province_id INTEGER NOT NULL REFERENCES provinces(id),
	code TEXT NOT NULL CHECK (char_length(code) <= 50),
	name TEXT NOT NULL CHECK (char_length(name) <= 50),
	status BOOLEAN NOT NULL DEFAULT 'false',
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER cities BEFORE UPDATE ON cities FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO cities ("id", "province_id", "code", "name", "deleted_at")
VALUES (0, 0, '', '', now());

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    code TEXT NOT NULL CHECK (char_length(code) <= 50),
    email TEXT NOT NULL CHECK (char_length(email) <= 100),
    name TEXT NOT NULL CHECK (char_length(name) <= 100),
    profile_image_id INTEGER REFERENCES files(id) DEFAULT 0,
    gender TEXT NOT NULL CHECK (char_length(gender) <= 50) DEFAULT '',
    phone TEXT NOT NULL CHECK (char_length(phone) <= 100) DEFAULT '',
    city_id INTEGER REFERENCES cities(id) DEFAULT 0,
    address TEXT NOT NULL CHECK (char_length(address) <= 500) DEFAULT '',
    settings JSONB NOT NULL CHECK (char_length(settings::TEXT) <= 1000) DEFAULT '{}',
    status BOOLEAN NOT NULL DEFAULT 'false',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO users ("id", "code", "email", "name", "deleted_at")
VALUES (0, '', '', '', now());

CREATE TABLE IF NOT EXISTS user_credentials (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    type TEXT NOT NULL CHECK (char_length(type) <= 50),
    email TEXT NOT NULL CHECK (char_length(email) <= 100),
    email_valid_at TIMESTAMP WITH TIME ZONE,
    password TEXT NOT NULL CHECK (char_length(password) <= 500),
    registration_details JSONB NOT NULL CHECK (char_length(registration_details::TEXT) <= 1000) DEFAULT '{}',
    status BOOLEAN NOT NULL DEFAULT 'false',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER user_credentials BEFORE UPDATE ON user_credentials FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO user_credentials ("id", "user_id", "type", "email", "email_valid_at", "password", "deleted_at")
VALUES (0, 0, '', '', now(), '', now());

CREATE TABLE IF NOT EXISTS meal_plans (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) DEFAULT 0,
    name TEXT NOT NULL CHECK (char_length(name) <= 100),
    min_age NUMERIC(20,3) NOT NULL DEFAULT 0,
	max_age NUMERIC(20,3) NOT NULL DEFAULT 0,
	interval NUMERIC(20,3) NOT NULL DEFAULT 0, -- in minutes
	suggestion_quantity NUMERIC(20,3) NOT NULL DEFAULT 0,
    unit TEXT NOT NULL CHECK (char_length(unit) <= 20) DEFAULT '',
	status TEXT NOT NULL CHECK (char_length(status) <= 20) DEFAULT '',
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER meal_plans BEFORE UPDATE ON meal_plans FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO meal_plans ("id", "name", "deleted_at")
VALUES (0, '', now());

CREATE TABLE IF NOT EXISTS user_childs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    name TEXT NOT NULL CHECK (char_length(name) <= 100),
    gender TEXT NOT NULL CHECK (char_length(gender) <= 20),
    birth_date TIMESTAMP WITH TIME ZONE,
    birth_length NUMERIC(20,3) NOT NULL DEFAULT 0,
    birth_weight NUMERIC(20,3) NOT NULL DEFAULT 0,
    head_circumference NUMERIC(20,3) NOT NULL DEFAULT 0,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER user_childs BEFORE UPDATE ON user_childs FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO user_childs ("id", "user_id", "name", "gender", "birth_date", "deleted_at")
VALUES (0, 0, '', '', now(), now());

CREATE TABLE IF NOT EXISTS user_child_meals (
    id SERIAL PRIMARY KEY,
    user_child_id INTEGER NOT NULL REFERENCES user_childs(id),
    meal_plan_id INTEGER NOT NULL REFERENCES meal_plans(id) DEFAULT 0,
    name TEXT NOT NULL CHECK (char_length(name) <= 100),
    suggestion_quantity NUMERIC(20,3) NOT NULL DEFAULT 0,
	quantity NUMERIC(20,3) NOT NULL DEFAULT 0,
	unit TEXT NOT NULL CHECK (char_length(unit) <= 20) DEFAULT '',
	scheduled_at TIMESTAMP WITH TIME ZONE,
	finish_at TIMESTAMP WITH TIME ZONE,
	status TEXT NOT NULL CHECK (char_length(status) <= 20) DEFAULT '', -- pending, done
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER user_child_meals BEFORE UPDATE ON user_child_meals FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO user_child_meals ("id", "user_child_id", "name", "finish_at", "deleted_at")
VALUES (0, 0, '', now(), now());

CREATE TABLE IF NOT EXISTS workshops (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (char_length(name) <= 100),
    description TEXT NOT NULL CHECK (char_length(description) <= 500),
    quota NUMERIC(20,3) NOT NULL DEFAULT 0,
	start_at TIMESTAMP WITH TIME ZONE,
	end_at TIMESTAMP WITH TIME ZONE,
	meet_info JSONB NOT NULL CHECK (char_length(meet_info::TEXT) <= 1000) DEFAULT '{}',
    status BOOLEAN NOT NULL DEFAULT 'false',
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER workshops BEFORE UPDATE ON workshops FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO workshops ("id", "name", "description", "start_at", "end_at", "deleted_at")
VALUES (0, 0, '', now(), now(), now());

CREATE TABLE IF NOT EXISTS workshop_participants (
    id SERIAL PRIMARY KEY,
    workshop_id INTEGER NOT NULL REFERENCES workshops(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER workshop_participants BEFORE UPDATE ON workshop_participants FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO workshop_participants ("id", "workshop_id", "user_id", "deleted_at")
VALUES (0, 0, 0, now());

CREATE TABLE IF NOT EXISTS follows (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    user_follow_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER follows BEFORE UPDATE ON follows FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO follows ("id", "user_id", "user_follow_id", "deleted_at")
VALUES (0, 0, 0, now());

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    follow_id INTEGER NOT NULL REFERENCES follows(id),
    type TEXT NOT NULL CHECK (char_length(type) <= 100),
    text TEXT NOT NULL CHECK (char_length(text) <= 100),
    file_id INTEGER NOT NULL REFERENCES files(id) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE TRIGGER chats BEFORE UPDATE ON chats FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
INSERT INTO chats ("id", "follow_id", "type", "text", "deleted_at")
VALUES (0, 0, '', '', now());

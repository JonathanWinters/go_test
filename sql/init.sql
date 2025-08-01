CREATE TABLE IF NOT EXISTS "level"(id SERIAL PRIMARY KEY,levelid TEXT NOT NULL,map bytea NOT NULL,position bytea NOT NULL,playerhitpoints INT,created timestamp DEFAULT NOW());

-- game
-- INSERT INTO level(id, levelid)
--     VALUES(1, 'test')
--     ON CONFLICT DO NOTHING;


--!INFO EFC: bytea vs json
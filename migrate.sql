CREATE SEQUENCE IF NOT EXISTS cities_id_seq START 1;

CREATE TABLE IF NOT EXISTS cities (
    id INTEGER DEFAULT nextval('cities_id_seq') PRIMARY KEY,
    name VARCHAR(255),
    country VARCHAR(255),
    lat FLOAT,
    lon FLOAT
);

INSERT INTO cities (name, country, lat, lon) VALUES ('London', 'GB', 51.5073219, -0.1276474);
INSERT INTO cities (name, country, lat, lon) VALUES ('Liverpool', 'GB', 53.4071991, -2.99168);
INSERT INTO cities (name, country, lat, lon) VALUES ('Manchester', 'GB', 53.4794892, -2.2451148);
INSERT INTO cities (name, country, lat, lon) VALUES ('Southampton', 'GB', 50.9025349, -1.404189);
INSERT INTO cities (name, country, lat, lon) VALUES ('Sheffield', 'GB', 53.3806626, -1.4702278);
INSERT INTO cities (name, country, lat, lon) VALUES ('Bristol', 'GB', 51.4538022, -2.5972985);
INSERT INTO cities (name, country, lat, lon) VALUES ('Leicester', 'GB', 52.6362, -1.1331969);
INSERT INTO cities (name, country, lat, lon) VALUES ('Coventry', 'GB', 52.4081812, -1.510477);
INSERT INTO cities (name, country, lat, lon) VALUES ('Nottingham', 'GB', 52.9534193, -1.1496461);
INSERT INTO cities (name, country, lat, lon) VALUES ('Sunderland', 'GB', 54.9058512, -1.3828727);
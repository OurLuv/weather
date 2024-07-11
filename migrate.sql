CREATE SEQUENCE IF NOT EXISTS cities_id_seq START 1;

CREATE TABLE IF NOT EXISTS cities (
    id INTEGER DEFAULT nextval('cities_id_seq') PRIMARY KEY,
    name VARCHAR(255),
    country VARCHAR(255),
    lat FLOAT,
    lon FLOAT
);

CREATE TABLE IF NOT EXISTS forecasts ( 
  city_id INTEGER NOT NULL, 
  temp FLOAT8 NOT NULL, 
  dt INTEGER NOT NULL, 
  json JSONB 
);

ALTER TABLE forecasts ADD CONSTRAINT forecasts_city_id_dt_key UNIQUE (city_id, dt);

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

INSERT INTO cities (name, country, lat, lon) VALUES ('Berlin', 'DE', 52.5170365, 13.3888599);
INSERT INTO cities (name, country, lat, lon) VALUES ('Hamburg', 'DE', 53.550341, 10.000654);
INSERT INTO cities (name, country, lat, lon) VALUES ('Munich', 'DE', 48.1371079, 11.5753822);
INSERT INTO cities (name, country, lat, lon) VALUES ('Stuttgart', 'DE', 48.7784485, 9.1800132);
INSERT INTO cities (name, country, lat, lon) VALUES ('Leipzig', 'DE', 51.3406321, 12.3747329);
INSERT INTO cities (name, country, lat, lon) VALUES ('Dortmund', 'DE', 51.5142273, 7.4652789);
INSERT INTO cities (name, country, lat, lon) VALUES ('Bremen', 'DE', 53.0758196, 8.8071646);
INSERT INTO cities (name, country, lat, lon) VALUES ('Dresden', 'DE', 51.0493286, 13.7381437);
INSERT INTO cities (name, country, lat, lon) VALUES ('Hanover', 'DE', 52.3744779, 9.7385532);
INSERT INTO cities (name, country, lat, lon) VALUES ('Mainz', 'DE', 50.0012314, 8.2762513);

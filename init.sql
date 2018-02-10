CREATE TABLE IF NOT EXISTS event (
  id         INT NOT NULL AUTO_INCREMENT,
  theme      VARCHAR(256),
  event_date DATE,
  created    DATE,
  starter    VARCHAR(512),
  main_dish  VARCHAR(512),
  dessert    VARCHAR(512),
  infotext   VARCHAR(2048),
  image      BLOB,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS participant (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(255),
  created  DATE,
  menu     INT,
  event_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE IF NOT EXISTS comment (
  id       INT NOT NULL AUTO_INCREMENT,
  content  VARCHAR(1024),
  name     VARCHAR(256),
  created  DATE,
  event_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE IF NOT EXISTS images (
  id       INT NOT NULL AUTO_INCREMENT,
  event_id INT,
  picture  BLOB,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE IF NOT EXISTS user (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(256) UNIQUE,
  salt     VARCHAR(256),
  password VARCHAR(256),
  PRIMARY KEY (id)
)

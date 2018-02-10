CREATE TABLE event (
  id        INT NOT NULL AUTO_INCREMENT,
  theme     VARCHAR(256),
  created   DATE,
  starter   VARCHAR(512),
  main_dish VARCHAR(512),
  dessert   VARCHAR(512),
  infotext  VARCHAR(2048),
  image     BLOB,
  PRIMARY KEY (id)
);

CREATE TABLE participant (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(255),
  created  DATE,
  menu     INT,
  event_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE comment (
  id       INT NOT NULL AUTO_INCREMENT,
  content  VARCHAR(1024),
  name     VARCHAR(256),
  created  DATE,
  event_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE images (
  id       INT NOT NULL AUTO_INCREMENT,
  event_id INT,
  picture  BLOB,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE user (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(256),
  salt     VARCHAR(256),
  password VARCHAR(256),
  PRIMARY KEY (id)
)

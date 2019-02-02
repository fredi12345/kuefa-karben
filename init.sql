#TODO Remove Drop Table after testing
DROP TABLE IF EXISTS kuefa_karben.participant;
DROP TABLE IF EXISTS kuefa_karben.comment;
DROP TABLE IF EXISTS kuefa_karben.images;
DROP TABLE IF EXISTS kuefa_karben.user;
DROP TABLE IF EXISTS kuefa_karben.event;

CREATE TABLE IF NOT EXISTS event (
  event_id     INT NOT NULL AUTO_INCREMENT,
  theme        VARCHAR(256),
  event_date   DATETIME,
  created_date DATETIME,
  starter      VARCHAR(512),
  main_dish    VARCHAR(512),
  dessert      VARCHAR(512),
  infotext     VARCHAR(2048),
  image_name    VARCHAR(256),
  PRIMARY KEY (event_id)
);

CREATE TABLE IF NOT EXISTS participant (
  id                  INT NOT NULL AUTO_INCREMENT,
  name                VARCHAR(255),
  message             VARCHAR(1024),
  participant_created DATETIME,
  menu                INT,
  event_id            INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (event_id)
);

CREATE TABLE IF NOT EXISTS comment (
  id              INT NOT NULL AUTO_INCREMENT,
  content         VARCHAR(1024),
  name            VARCHAR(256),
  comment_created DATETIME,
  event_id        INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (event_id)
);

CREATE TABLE IF NOT EXISTS images (
  id        INT NOT NULL AUTO_INCREMENT,
  event_id  INT,
  image_name VARCHAR(256),
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (event_id)
);

CREATE TABLE IF NOT EXISTS user (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(256) UNIQUE,
  salt     VARCHAR(256),
  password VARCHAR(256),
  PRIMARY KEY (id)
);

# TODO Remove after Testing is done:
INSERT INTO user (name, salt, password)
VALUES ('test', '1Y6LMQth5V', '0d16233affc52371347cbb20123ff8157158e8589bcfeabff2e44d30891fc32a');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('testtheme',
        Now(),
        NOW(),
        'teststarter',
        'testmaindish',
        'testdessert',
        'testinfotext',
        'public/images/first-event.png');
INSERT INTO participant (name, participant_created, menu, event_id)
VALUES ('testname', Now(), 1, (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO comment (content, name, comment_created, event_id)
VALUES ('testcontent', 'testname', Now(), (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO images (event_id, image_name)
VALUES ((SELECT event_id FROM event ORDER BY event_id LIMIT 1), 'public/images/first-event.png');

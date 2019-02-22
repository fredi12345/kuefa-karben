# Following SQL is only to create a Testinstance. !!!CURRENT DATABASE WILL BE ERASED!!!
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
  FOREIGN KEY (event_id) REFERENCES event (event_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comment (
  id              INT NOT NULL AUTO_INCREMENT,
  content         VARCHAR(1024),
  name            VARCHAR(256),
  comment_created DATETIME,
  event_id        INT,
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (event_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS images (
  id        INT NOT NULL AUTO_INCREMENT,
  event_id  INT,
  image_name VARCHAR(256),
  PRIMARY KEY (id),
  FOREIGN KEY (event_id) REFERENCES event (event_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user (
  id       INT NOT NULL AUTO_INCREMENT,
  name     VARCHAR(256) UNIQUE,
  salt     VARCHAR(256),
  password VARCHAR(256),
  PRIMARY KEY (id)
);

#User
INSERT INTO user (name, salt, password)
VALUES ('test', '2s5rcqth1r', '870ba668965e5b822160c1b4d044c43edded546e7c9538e2339b7e6cf9374612');

#Events
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Testevent',
        '2020-01-01T00:30',
        '2018-07-19T22:12',
        'Eine leckere Vorspeise',
        'Ein leckeres Hauptgericht',
        'Ein leckerer Nachtisch',
        'Wir freuen uns auf eure Teilnahme.',
        'Cover-Test.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Herbstessen',
        '2016-10-16T17:30',
        '2018-07-19T22:12',
        'Suppe',
        'Braten mit Klößen und Gemüse',
        'Nachtisch',
        'Das erste Herbstessen der Küfa Karben findet im Jugendcafé im JuKuZ statt. Wir freuen uns auf ein gemeinsames Essen in entspannter Atmosphäre!',
        'Cover-2016-10-16.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Weihnachtsessen',
        '2016-11-27T17:30',
        '2018-07-19T22:12',
        'Heiße Bouillon',
        'Gulasch mit Rotkohl und Bandnudeln vegane Variante: Champignongulasch',
        'Spekulatius-Orangen-Creme',
        'Gibt es einen geeigneteren Anlass als den 1. Advent, um zum Essen einzuladen? Vermutlich, aber wir nehmen mit diesem gerne Vorlieb.',
        'Cover-2016-11-27.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Frühlingsessen',
        '2017-03-19T17:30',
        '2018-07-19T22:12',
        'Spinatsalat mit Kartoffelcroutons',
        'Hähnchen (bzw. Tofuspieße) mit Gnocchis und Bärlauchpesto',
        'Blätterteig mit Ziegenkäse, Honig und Nüssen',
        'Nach langer Pause folgt nun endlich das dritte gemeinsame Essen der Küfa! Wir freuen uns auf euch!',
        'Cover-2017-03-19.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('4. Küfa-Essen',
        '2017-04-23T18:00',
        '2018-07-19T22:12',
        'Frühlingsgemüse in Bierteig ausgebacken, Kräuterdip',
        'Frankfurter Kartoffelwurst im Speckmantel',
        'Starkbiersoße an Kartoffeln und Gemüse., warme Malzbierbonbons mit Vanillesoße',
        'Am 23. April ist Tag des deutschen Bieres und ganz mottogetreu wird auch unser Menü rund ums Bier konzipiert sein. Natürlich gibt es für unsere jüngeren Gäste auch bierfreie Varationen! :)',
        'Cover-2017-04-23.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Gemeinsames Grillen',
        '2017-06-11T18:00',
        '2018-07-19T22:12',
        'Gegrilltes Capresebrot',
        'Teriyaki-Spieße mit gegrilltem Gemüse und Erdnusssoße',
        'Gegrillter Pfirsich mit Pistazieneis und Minze',
        'Wir wollen gemeinsam mit euch die Grillzeit einleiten! Kommt vorbei, wir freuen uns. Es gibt Hansis Grill-Menü!',
        'Cover-2017-06-11.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Burgertime',
        '2017-07-16T17:30',
        '2018-07-19T22:12',
        'Süßkartoffelchips mit dreierlei Dip',
        'Selbstgemachte Burger mit Krautsalat',
        'S´mores mit Heidelbeeren',
        'Es gibt Burger, klassisch, vegetarisch und vegan. Der Termin steht! Wir freuen uns auf euch!',
        'Cover-2017-07-16.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('1 Jahr Küfa - Immer wieder kommt ein neuer Herbst',
        '2017-10-29T17:30',
        '2018-07-19T22:12',
        'Tomatensuppe mit Petersilienpesto und selbstgebackenem Brot',
        'Mediteranes Hähnchen natürlich wird es auch eine vegane Variante geben',
        'Gebratene Grießnocken mit heißen Kirschen',
        'Ein Jahr ist vorbei und auch dieses Mal haben wir uns ein herbstliches Menü für euch überlegt',
        'Cover-2017-10-29.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Novemberessen',
        '2017-11-26T17:30',
        '2018-07-19T22:12',
        'Flädlesuppe auf Gemüse- bzw. Rinderbrühbasis',
        'Kohlrouladen mit Kartoffelpüree',
        'Erdnusskürbispralinen mit Haferkrokant',
        'Wir freuen uns auf euch!',
        'Cover-2017-11-26.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Asiatisch aufgetischt!',
        '2018-02-25T17:30',
        '2018-07-19T22:12',
        'Frühlingsrollen mit dreierlei Dips',
        'Asiatisches Curry mit Reis und Hähnchen- oder Tofuspießen',
        'Gebackene Banane mit Vanilleeis und süßem Honig oder Ahornsirup',
        'Kommt vorbei und genießt zu den Speisen die Getränke zum Jugendcafépreis. Wir freuen uns auf euch!',
        'Cover-2018-02-25.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Herzlich Hessisch Uffgetischt',
        '2018-05-20T17:30',
        '2018-07-19T22:12',
        'Handkässalat',
        'Grüne Soße mit Kartoffeln und Schnitzel (natürlich auch als vegane Variante',
        'Apple Crumble',
        'Wir freuen uns auf euch!',
        'Cover-2018-05-20.png');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name)
VALUES ('Italienischer Sommer',
        '2018-06-24T17:30',
        '2018-07-19T22:12',
        'Tomaten-Melonen-Salat',
        'Gnocchi an Pesto Penne Bolognese Spaghetti Aglio e Olio',
        'Ricotta-Cannelloni mit warmer Himbeersoße',
        'Wir freuen uns auf euch!',
        'Cover-2018-06-24.png');

#Participants
INSERT INTO participant (name, participant_created, menu, message, event_id)
VALUES ('Tester 1', Now(), 0, '', (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO participant (name, participant_created, menu, message, event_id)
VALUES ('Tester 2', Now(), 1, 'Ich bring meine Kinder mit', (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO participant (name, participant_created, menu, message, event_id)
VALUES ('Tester 3',
        Now(),
        2,
        'Ich bringe meine Kinder nicht mit',
        (SELECT event_id FROM event ORDER BY event_id LIMIT 1));

#Comments
INSERT INTO comment (content, name, comment_created, event_id)
VALUES ('Test 1 Hallo testtesttest', 'Tester 1', Now(), (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO comment (content, name, comment_created, event_id)
VALUES ('Test 2 Hallo ÜÄßö /\\#ashkl%20iqu', 'Tester 2', Now(), (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO comment (content, name, comment_created, event_id)
VALUES ('testcontent', 'Tester 3', Now(), (SELECT event_id FROM event ORDER BY event_id LIMIT 1));

#Pictures
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');
INSERT INTO images (event_id, image_name)
VALUES (1, 'Cover-Test.png');

#Covers
#INSERT INTO images (event_id, image_name) VALUES ((SELECT event_id FROM event ORDER BY event_id LIMIT 1), 'public/images/first-event.png' );
INSERT INTO images (event_id, image_name)
VALUES (2, 'Cover-2016-10-16.png');
INSERT INTO images (event_id, image_name)
VALUES (3, 'Cover-2016-11-27.png');
INSERT INTO images (event_id, image_name)
VALUES (4, 'Cover-2017-03-19.png');
INSERT INTO images (event_id, image_name)
VALUES (5, 'Cover-2017-04-23.png');
INSERT INTO images (event_id, image_name)
VALUES (6, 'Cover-20117-06-11.png');
INSERT INTO images (event_id, image_name)
VALUES (7, 'Cover-2017-07-16.png');
INSERT INTO images (event_id, image_name)
VALUES (8, 'Cover-2017-10-29.png');
INSERT INTO images (event_id, image_name)
VALUES (9, 'Cover-2017-11-26.png');
INSERT INTO images (event_id, image_name)
VALUES (10, 'Cover-2018-02-25.png');
INSERT INTO images (event_id, image_name)
VALUES (11, 'Cover-2018-05-20.png');
INSERT INTO images (event_id, image_name)
VALUES (12, 'Cover-2018-06-24.png');

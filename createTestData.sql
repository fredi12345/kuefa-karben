# Following SQL is only to create some Testdata

# INSERT INTO user ( name, salt, password) VALUES ('test','1Y6LMQth5V','0d16233affc52371347cbb20123ff8157158e8589bcfeabff2e44d30891fc32a');
INSERT INTO event (theme, event_date, created_date, starter, main_dish, dessert, infotext, image_url) VALUES ('testtheme',Now(), NOW(),'teststarter','testmaindish','testdessert','testinfotext','public/images/first-event.png' );
INSERT INTO participant (name, participant_created, menu, event_id) VALUES ('testname', Now(), 1, (SELECT event_id FROM event ORDER BY  event_id LIMIT 1));
INSERT INTO comment (content, name, comment_created, event_id) VALUES ('testcontent', 'testname', Now(), (SELECT event_id FROM event ORDER BY event_id LIMIT 1));
INSERT INTO images (event_id, image_url) VALUES ((SELECT event_id FROM event ORDER BY  event_id LIMIT 1), 'public/images/first-event.png' );

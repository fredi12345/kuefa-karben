## Nur die nehmen, die auch vorhanden sind (die die nicht gelb hinterlegt sind)....
ALTER TABLE participant DROP FOREIGN KEY participant_ibfk_1;
ALTER TABLE images DROP FOREIGN KEY images_ibfk_1;
ALTER TABLE comment DROP FOREIGN KEY comment_ibfk_1;
ALTER TABLE participant DROP FOREIGN KEY participant_ibfk_2;
ALTER TABLE images DROP FOREIGN KEY images_ibfk_2;
ALTER TABLE comment DROP FOREIGN KEY comment_ibfk_2;
ALTER TABLE participant DROP FOREIGN KEY participant_ibfk_3;
ALTER TABLE images DROP FOREIGN KEY images_ibfk_3;
ALTER TABLE comment DROP FOREIGN KEY comment_ibfk_3;
ALTER TABLE participant DROP FOREIGN KEY participant_ibfk_4;
ALTER TABLE images DROP FOREIGN KEY images_ibfk_4;
ALTER TABLE comment DROP FOREIGN KEY comment_ibfk_4;
ALTER TABLE participant DROP FOREIGN KEY participant_ibfk_5;
ALTER TABLE images DROP FOREIGN KEY images_ibfk_5;
ALTER TABLE comment DROP FOREIGN KEY comment_ibfk_5;

## Erst die Statements oben, dann EINMAL diese:
ALTER TABLE participant
ADD FOREIGN KEY (event_id)
REFERENCES event(event_id)
ON DELETE CASCADE;

ALTER TABLE images
ADD FOREIGN KEY (event_id)
REFERENCES event(event_id)
ON DELETE CASCADE;

ALTER TABLE comment
ADD FOREIGN KEY (event_id)
REFERENCES event(event_id)
ON DELETE CASCADE;

## Zeigt Version an:
SELECT @@Version;


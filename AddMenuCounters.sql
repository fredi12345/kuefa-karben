ALTER TABLE participant
    ADD classic_count int DEFAULT 0,
    ADD vegetarian_count int DEFAULT 0,
    ADD vegan_count int DEFAULT 0;

UPDATE participant set participant.classic_count = 1 where participant.menu=0;
UPDATE participant set participant.vegetarian_count = 1 where participant.menu=1;
UPDATE participant set participant.vegan_count = 1 where participant.menu=2;

#achtung
ALTER table participant drop menu;

# zum löschen von den testkommentaren, einfach die eventIDs raussuchen und hiermit löschen
#DELETE from comment where event_id = 2

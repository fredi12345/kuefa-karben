alter table event
    add closing_date datetime null;
UPDATE event set event.closing_date = event_date;
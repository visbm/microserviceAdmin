ALTER TABLE BOOKING
paid             BOOLEAN NOT NULL;

UPDATE BOOKING SET paid = 'true' WHERE id = 1;
UPDATE BOOKING SET paid = 'true' WHERE id = 2;
UPDATE BOOKING SET paid = 'false' WHERE id = 3;
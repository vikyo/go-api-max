-- SQLite
SELECT id, email, password, created_at, updated_at
FROM users;

DELETE from users;

SELECT Id, name, `desc`, location, dateTime, user_id
FROM events;

DELETE from events;

SELECT id, user_id, event_id
FROM registrations;
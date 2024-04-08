-- users
INSERT INTO users ("id","name", "role_id", "status")
VALUES ('b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','kohki', 'b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','true');

-- todos
INSERT INTO todos ("id","user_id","name","target_date","done_date","status")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93f872f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task1','2021-01-01 00:00:00','2021-01-01 00:00:00','true'),
        ('91d1a74f-8d0e-7287-b218-31e8fdaf8ad5','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task2', '2021-01-01 00:00:00', '2021-01-01 00:00:00','true');

-- calendar
INSERT INTO calendar ("id","calendar_name")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93f852f','parsonal');

-- events
INSERT INTO events ("id","user_id","tag_id","title","start_date","end_date")
VALUES ('e1d388da-47e6-ca1d-2722-0a41d93f872f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','e1d388da-47e0-ca1d-2722-0a41d93f872h','event1','2024-05-01 12:00:00','2024-05-01 15:00:00'),
        ('e1d388da-47e0-ca1d-2722-0a41d93f872g','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','g1d388da-47e0-ca1d-2722-0a41d93f872f','event2','2024-05-02 00:00:00','2024-05-06 00:00:00');

-- roles
INSERT INTO roles ("id","role_name")
VALUES ('b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2f5','admin');

-- tags
INSERT INTO tags ("id","user_id","tag_name","tag_color")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93f872h','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','school','#ff0000'),
        ('e1d388da-47e0-ca1d-2722-0a41d93f872g','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','part-time','#00ff00');

-- user_calendar
INSERT INTO user_calendar ("id","user_id","calendar_id","access_level")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93g872f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','e1d388da-47e0-ca1d-2722-0a41d93f852f','admin');

-- calendar_events
INSERT INTO calendar_events ("id", "calendar_id","event_id")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d73f872f','e1d388da-47e0-ca1d-2722-0a41d93f852f','e1d388da-47e6-ca1d-2722-0a41d93f872f'),
        ('e1d388da-47e0-ca1d-2722-0a41d93f872g','e1d388da-47e0-ca1d-2722-0a41d93f852f','e1d388da-47e0-ca1d-2722-0a41d93f872g');
```
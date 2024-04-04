-- users
INSERT INTO users ("id","name", "role_id", "status")
VALUES ('b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','kohki', 'b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','true');

-- todos
INSERT INTO todos ("id","user_id","name","target_date","done_date","status")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93f872f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task1','2021-01-01 00:00:00','2021-01-01 00:00:00','true'),
        ('91d1a74f-8d0e-7287-b218-31e8fdaf8ad5','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task2', '2021-01-01 00:00:00', '2021-01-01 00:00:00','true');
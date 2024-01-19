BEGIN;

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('6f8191bf-7f23-4050-b0f6-a606e0b99eb2', 'admin', '管理员',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('695d06d1-b946-467c-a874-d1f3c94fea89', 'zhangsan', '张三',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('4022032a-fe06-41c2-bcb3-820170f83568', 'lisi', '李四',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('bdf2d865-7a49-46ed-b4de-b3e342cf2753', 'wangwu', '王五',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('b6133b82-0070-42c5-aa22-648ae31510ff', 'zhaoliu', '赵六',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

INSERT INTO public."user" (id, username, nick_name, password, status, tel, email, create_time, update_time,
                           deleted)
VALUES ('d2d4ff16-ca61-4ddf-b1ba-8b627c871f05', 'sunqi', '孙七',
        '123456', 3, '', '',
        '2022-12-07 03:35:33.490506 +00:00',
        '2023-08-15 10:23:29.923340 +00:00', 0);

COMMIT;



show DATABASES


create database GinOrmDemo;
use GinOrmDemo;


CREATE TABLE usertable (
    id INT NOT NULL PRIMARY KEY,
    username VARCHAR(255),
    age TINYINT,
    email VARCHAR(255),
    add_time INT
);

-- 插入第一条模拟数据
INSERT INTO usertable (id, username, age, email, add_time)
VALUES (1, 'john_doe', 28, 'john@example.com', 1620000000);

-- 插入第二条模拟数据
INSERT INTO usertable (id, username, age, email, add_time)
VALUES (2, 'jane_smith', 32, 'jane@example.com', 1620100000);

-- 插入第三条模拟数据
INSERT INTO usertable (id, username, age, email, add_time)
VALUES (3, 'mike_johnson', 45, 'mike@example.com', 1620200000);


select * from usertable;
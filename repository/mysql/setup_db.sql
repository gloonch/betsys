CREATE TABLE server (
    id int primary key AUTO_INCREMENT,
    name varchar(255) not null unique ,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);
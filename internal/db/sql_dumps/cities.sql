create table cities
(
    id   serial
        primary key,
    name varchar(255) not null
);

comment on table cities is 'Города';
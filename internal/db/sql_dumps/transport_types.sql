create table transport_types
(
    id   serial
        primary key,
    name varchar(255) not null
);

comment on table transport_types is 'Виды транспорта';
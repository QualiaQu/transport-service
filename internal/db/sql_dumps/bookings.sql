create table bookings
(
    id           serial
        primary key,
    user_id      integer                             not null,
    route_id     integer                             not null
        references routes,
    booking_time timestamp default CURRENT_TIMESTAMP not null
);

comment on table bookings is 'Таблица для хранения информации о бронированиях';

comment on column bookings.user_id is 'ID пользователя';

comment on column bookings.route_id is 'ID рейса';

comment on column bookings.booking_time is 'Время бронирования';

alter table bookings
    owner to postgres;


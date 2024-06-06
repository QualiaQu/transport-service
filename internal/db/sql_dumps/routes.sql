create table routes
(
    id                 serial
        primary key,
    origin             integer
        references cities,
    destination        integer
        references cities,
    transport_type     integer
        references transport_types,
    price              numeric(10, 2)        not null,
    departure_datetime timestamp             not null,
    arrival_datetime   timestamp             not null,
    is_booked          boolean default false not null
);

comment on table routes is 'Рейсы';

comment on column routes.origin is 'Откуда';

comment on column routes.destination is 'Куда';

comment on column routes.transport_type is 'Вид транспорта';

comment on column routes.price is 'Цена';

comment on column routes.departure_datetime is 'Дата-время отправления';

comment on column routes.arrival_datetime is 'Дата-время прибытия';

comment on column routes.is_booked is 'Наличие брони';
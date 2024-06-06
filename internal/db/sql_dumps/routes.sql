CREATE TABLE routes (
                        id SERIAL PRIMARY KEY,
                        origin INT REFERENCES cities(id),
                        destination INT REFERENCES cities(id),
                        date DATE NOT NULL,
                        transport_type INT REFERENCES transport_types(id),
                        price DECIMAL(10, 2) NOT NULL,
                        departure_datetime TIMESTAMP NOT NULL,
                        arrival_datetime TIMESTAMP NOT NULL
);

CREATE TABLE videojuegos
(
    id     bigint unsigned not null primary key auto_increment,
    nombre VARCHAR(255)    NOT NULL,
    genero VARCHAR(255)    NOT NULL,
    anio   INTEGER         NOT NULL
);

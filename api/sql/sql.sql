/* CREATE DATABASE IF NOT EXISTS dbgors;
USE dbgors;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment PRIMARY KEY,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    senha varchar(20) NOT NULL UNIQUE,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB; 

CREATE TABLE seguidores(
	usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    
    primary key(usuario_id, seguidor_id)
)ENGINE=INNODB;

CREATE TABLE publicacoes(
    id INT auto_increment PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas INT default 0,
    criadaEm timestamp default current_timestamp
    
)ENGINE=INNODB;

*/
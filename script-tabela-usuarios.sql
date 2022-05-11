/*
create database bdgo;
use bdgo


show tables;

create table usuarios(
id int auto_increment primary key,
nome varchar(50) not null,
email varchar(50) not null) ENGINE=INNODB;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON bdgo.* TO 'golang'@'localhost';
---------------------------------------------------------
CREATE DATABASE IF NOT EXISTS dbgors;
USE dbgors;
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

-- tabelas
CREATE TABLE usuarios(
    id int auto_increment PRIMARY KEY,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    senha varchar(100) NOT NULL,
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

-- dados
insert into usuarios (nome, nick, email, senha)
VALUES
("Usuario 1","usuario_1","usuario1@email.com","$2a$10$7Ap7AjVrPKQ5ErPhCjsR9.43XGHttQiD56BnOUWhU03LVRnlPgukG"),
("Usuario 2","usuario_2","usuario2@email.com","$2a$10$7Ap7AjVrPKQ5ErPhCjsR9.43XGHttQiD56BnOUWhU03LVRnlPgukG"),
("Usuario 3","usuario_3","usuario3@email.com","$2a$10$7Ap7AjVrPKQ5ErPhCjsR9.43XGHttQiD56BnOUWhU03LVRnlPgukG");


INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1,2),
(3,1),
(1,3);

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

USE dbgors;

desc usuarios;
desc seguidores;
desc publicacoes;
select * from usuarios;
select * from seguidores;
select * from publicacoes;
-- delete from publicacoes where id = 4
-- delete from seguidores where usuario_id in (1,2,3);

insert into publicacoes(titulo, conteudo, autor_id) values
("Publicacao do usuario 1", "Texto da publicacao do usuario1", 1),
("Publicacao do usuario 2", "Texto da publicacao do usuario2", 2),
("Publicacao do usuario 3", "Texto da publicacao do usuario3", 3);


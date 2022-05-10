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


insert into publicacoes(titulo, conteudo, autor_id) values
("Publicacao do usuario 1", "Texto da publicacao do usuario1", 1),
("Publicacao do usuario 2", "Texto da publicacao do usuario2", 2),
("Publicacao do usuario 3", "Texto da publicacao do usuario3", 3);
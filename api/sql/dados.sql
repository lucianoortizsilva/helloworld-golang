insert into usuarios (nome, nick, email, senha)
values
("Ronaldo", "Fenômeno", "ronaldo@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario1
("Ronaldinho", "Gaúcho", "ronaldinho@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario2
("Pelé", "Rei", "pele@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Ronaldo Nazário", "Eu sou o Ronaldo!", 1),
("Ronaldinho Gaúcho", "Eu sou Ronaldinho Gaúcho!", 2),
("Pelé", "Eu sou Pelé! Oba!", 3);
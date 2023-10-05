insert into roles(kode, nama) values 
('ADMIN', 'Admin'),
('PENILAI', 'Penilai'),
('PETUGASBPN', 'Petugas BPN'),
('MASYARAKAT', 'Masyarakat');

insert into user_datas(uuid, nama) values 
('f0e0c976-2066-4d71-9bd0-fa1490164281', 'Admin Aplikasi');

insert into users(email, password, role, user_data) values 
('admin@gmail.com', '$2a$10$Ff86pYrOwzQvwxbKnrMpXu2iGCxeViyUTbxqCDbur6XwQ98iso5zW', 'ADMIN', 'f0e0c976-2066-4d71-9bd0-fa1490164281')
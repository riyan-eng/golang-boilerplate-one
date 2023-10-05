create table user_datas(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    nama varchar(255),
    nik varchar(255),
    nomor_telepon varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(nik),
    unique(nomor_telepon)
);
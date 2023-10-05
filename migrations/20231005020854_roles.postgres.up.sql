create table roles(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    kode varchar(255),
    nama varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(kode)
);
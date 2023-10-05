create table users(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    email varchar(255),
    "role" varchar(255),
    "password" text,
    user_data uuid,
    is_aktif boolean not null default true,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(email),
    constraint fk_role foreign key(role) references roles(kode) on delete set null,
    constraint fk_user_data foreign key(user_data) references user_datas(uuid) on delete set null
);
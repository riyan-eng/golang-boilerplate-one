create table objects(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    bucket varchar(255),
    nama varchar(255),
    "owner" uuid,
    "size" int,
    mime_type varchar(255),
    "path" varchar,
    "url" varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(bucket, nama),
    constraint fk_bucket foreign key(bucket) references buckets(nama) on delete cascade
);
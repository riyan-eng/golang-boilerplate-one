create table buckets(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    nama varchar(255),
    "owner" uuid,
    public boolean not null default false,
    file_size_limit int,
    allowed_mime_types varchar(255) array,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(nama)
)
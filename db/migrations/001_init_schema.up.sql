create table if not exists sources (
    id              integer         primary key not null,
    name            varchar(255)    not null,
    date_col        integer         not null,
    descr_col       integer         not null,
    debit_col       integer         not null,
    credit_col      integer         not null,
    pii_cols        text,
    separator       varchar(255)    not null, -- to check this

    created_at      datetime        not null,
    updated_at      datetime        not null
);

create table if not exists categories (
    id              integer         primary key not null,
    name            varchar(255)    not null,

    created_at      datetime        not null,
    updated_at      datetime        not null
);

create table if not exists matchers (
    id              integer         primary key not null,
    enabled         boolean         default true,
    source          integer,
    matcher         text            not null,
    category        integer,
    related         text,
    num_matches     integer         default 0,

    created_at      datetime        not null,
    updated_at      datetime        not null,

    foreign key(source)             references sources(id),
    foreign key(category)           references categories(id)
);

create table if not exists cities (
    city            varchar(255)    not null,
    state_code      char(2)         not null,
    state           varchar(255)    not null

    primary key(city, state_code)
);

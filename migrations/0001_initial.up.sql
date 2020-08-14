create table if not exists matchers (
    id          integer         primary key not null,
    enabled     boolean         default true,
    matcher     text            not null,
    account     varchar(255)    not null,
    related     text,
    num_matches integer         default 0,

    created     datetime        not null,
    updated     datetime        not null
);

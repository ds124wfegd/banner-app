CREATE TABLE users
(
    id                 serial       not null unique,
    name               varchar(255) not null,
    username           varchar(255) not null unique,
    password_hash      varchar(255) not null
);

CREATE TABLE admins
(
    id                 int          not null unique,
    username           varchar(255) not null unique,
    password_hash      varchar(255) not null,
    admin_status       boolean      not null default true,
    system_password    varchar(255) not null
);

CREATE TABLE banners
(
    id                serial                not null unique,
    is_active         int                   not null,
    feature_id        int                   not null,
    tag_id_1          int                   not null,
    tag_id_2          int,
    tag_id_3          int
);

CREATE TABLE content
(
    id          serial                          not null unique,
    banner_id   int,
	title       varchar(255),
	some_title  varchar(255),
	text        varchar(255),
	some_text   varchar(255),
	some_url    varchar(255)
);


CREATE TABLE feature_banner
(
    feature_id  int         not null,
    banner_id   int[]       not null
);

CREATE TABLE tag_banner
(
    tag_id_1          int          not null,
    tag_id_2          int,
    tag_id_3          int,
    banner_id         int[]        not null
);



create table `derma`.`user`
(
    `id`               bigint                                                                       not null,
    `username`         varchar(255)                                                                 not null unique,
    `password`         varchar(255)                                                                 not null,
    `avatar`           varchar(255) default 'https://files.ozline.icu/images/avatar.jpg'            not null comment 'url',
    `signature`        varchar(255) default 'NOT NULL BUT SEEMS NULL'                               not null comment '255charmax',
    `email`            varchar(255)                                                                 not null,
    `phone`            bigint                                                                       not null,
    `birth`            date                                                                         not null,
    `created_at`       timestamp    default current_timestamp                                       not null,
    `updated_at`       timestamp    default current_timestamp                                       not null on update current_timestamp comment 'update profile time',
    `deleted_at`       timestamp    default null null,
    constraint `id`
        primary key (`id`)
) engine=InnoDB default charset=utf8mb4;
create table `derma`.`user`
(
    `id`               bigint                                                                       not null,
    `username`         varchar(255)                                                                 not null unique,
    `password`         varchar(255)                                                                 not null,
    `avatar`           varchar(255) default 'https://s2.loli.net/2024/01/02/bhezZY4OPAFDHd2.jpg'    not null comment 'url',
    `signature`        varchar(255) default 'NOT NULL BUT SEEMS NULL'                               not null comment '255charmax',
    `email`            varchar(255)                                                                 not null,
    `phone`            bigint                                                                       not null,
    `birth`            date         default '2006-03-02'                                                                    not null,
    `created_at`       timestamp    default current_timestamp                                       not null,
    `updated_at`       timestamp    default current_timestamp                                       not null on update current_timestamp comment 'update profile time',
    `deleted_at`       timestamp    default null null,
    constraint `id`
        primary key (`id`)
) engine=InnoDB default charset=utf8mb4;

create table `derma`.`picture`
(
    `id`              bigint                              not null,
    `user_id`         bigint                              not null,
    `picture_url`     varchar(255)                        not null comment 'url',
    `percent`         varchar(255)                        not null,
    `describe`        varchar(255)                        not null,
    `created_at`      timestamp default current_timestamp not null,
    `updated_at`      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at`      timestamp default null null,
    constraint `id`
        primary key (`id`),
    constraint `picture_user`
        foreign key (`user_id`)
            references `derma`.`user` (`id`)
            on delete cascade
)engine=InnoDB default charset=utf8mb4;
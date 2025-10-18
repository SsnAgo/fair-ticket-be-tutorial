create schema `fair-ticket-tutorial` collate utf8mb4_general_ci;

use `fair-ticket-tutorial`;

create table participant
(
    id                 bigint auto_increment
        primary key,
    name               varchar(255) not null,
    address            varchar(255) not null,
    lucky_num          bigint       not null,
    project_id         bigint       not null,
    project_onchain_id bigint       not null,
    win                tinyint(1)   not null,
    merkle_proof       text         not null comment '数组字符串',
    created_at         datetime     not null,
    updated_at         datetime     not null,
    deleted_at         datetime     null
);

create index participant_address_index
    on participant (address);

create table project
(
    id           int auto_increment comment '自增id'
        primary key,
    onchain_id   bigint       not null comment '智能合约中的project id',
    fingerprint  varchar(255) not null comment '链上传入，emit事件带出，用于关联链上链下',
    name         varchar(255) not null,
    description  text         null,
    image_url    text         not null,
    total_supply int          not null,
    lottery_num  bigint       not null,
    merkle_root  text         not null,
    status       varchar(100) not null,
    owner        varchar(255) not null,
    created_at   datetime     not null,
    updated_at   datetime     not null,
    deleted_at   datetime     null
);

create index project_fingerprint_index
    on project (fingerprint);

create index project_onchain_id_index
    on project (onchain_id);
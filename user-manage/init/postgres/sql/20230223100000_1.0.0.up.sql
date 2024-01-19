BEGIN;

create table if not exists public."user"
(
    id          uuid not null
        primary key,
    username    text,
    nick_name   text,
    password    text,
    status      integer,
    tel         text,
    email       text,
    create_time timestamp with time zone,
    update_time timestamp with time zone,
    deleted     integer default 0
);

comment on table public."user" is '用户表';

comment on column public."user".username is '用户名';

comment on column public."user".nick_name is '昵称';

comment on column public."user".password is '密码';

comment on column public."user".status is '状态，1：正常，2：禁用，3：锁定';

comment on column public."user".tel is '联系电话';

comment on column public."user".email is '邮箱';

comment on column public."user".create_time is '创建时间';

comment on column public."user".update_time is '更新时间';

comment on column public."user".deleted is '删除时间';

create unique index if not exists user_nick_name_deleted_uindex
    on public."user" (nick_name, deleted);

create index if not exists user_nick_name_index
    on public."user" (nick_name);

create unique index if not exists user_username_deleted_uindex
    on public."user" (username, deleted);

create index if not exists user_username_index
    on public."user" (username);


COMMIT;



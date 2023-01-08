create database contacts_db;

create schema account;

create table account.account
(
    id       uuid not null
        constraint account_id_pk
            primary key,
    login    text not null
        constraint account_login_pk
            unique,
    password_hash text not null,
    created_at    timestamp with time zone default now() not null,
    updated_at    timestamp with time zone default now() not null
);

create table account.contact
(
    id         uuid                      not null
        constraint contact_id_pk
            primary key,
    account_id uuid
        constraint contact_account_id_fk
            references account.account (id),
    full_name  text                      not null,
    email      text                      not null,
    phone      text                      not null,
    address    text                      not null,
    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null
);

create table account.item
(
    id          uuid    not null
        constraint item_id_pk
            primary key,
    name        text    not null,
    description text    not null,
    price       numeric not null,
    contact_id  uuid    not null
        constraint item_contact_id_fk
            references account.contact (id)
);



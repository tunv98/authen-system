create database if not exists authen_system;

### users
create table users (
                       id bigint unsigned auto_increment primary key,
                       created_at datetime(3) null,
                       updated_at datetime(3) null,
                       deleted_at datetime(3) null,
                       full_name longtext not null,
                       phone_number varchar(191) null,
                       email varchar(191) null,
                       user_name varchar(191) null,
                       pass_word longtext not null,
                       birthday longtext null,
                       latest_login datetime(3) null,
                       constraint email unique (email),
                       constraint phone_number unique (phone_number),
                       constraint user_name unique (user_name)
);
create index idx_users_deleted_at on users (deleted_at);

### campaigns
create table campaigns (
                           id bigint unsigned auto_increment primary key,
                           created_at datetime(3) null,
                           updated_at datetime(3) null,
                           deleted_at datetime(3) null,
                           campaign_name varchar(191) null,
                           total_vouchers bigint unsigned not null,
                           discount_value double null,
                           start_date datetime(3) null,
                           end_date datetime(3) null,
                           constraint campaign_name unique (campaign_name)
);
create index idx_campaigns_deleted_at on campaigns (deleted_at);

### vouchers
create table vouchers (
                          id bigint unsigned auto_increment primary key,
                          created_at datetime(3) null,
                          updated_at datetime(3) null,
                          deleted_at datetime(3) null,
                          code varchar(191) null,
                          campaign_id bigint unsigned null,
                          user_id bigint unsigned null,
                          expired_time datetime(3) null,
                          status enum ('active', 'used', 'expired') null,
                          constraint code unique (code),
                          constraint fk_vouchers_campaign foreign key (campaign_id) references campaigns (id),
                          constraint fk_vouchers_user foreign key (user_id) references users (id)
);
create index idx_vouchers_deleted_at on vouchers (deleted_at);


## inserts

insert into campains (campaign_name, total_vouchers, discount_value, start_date, end_date)
values ('TOPUP_FIRST_LOGIN', 100, 0.3, '2024-03-08 02:35:51.000', '2024-03-15 03:36:09.000')

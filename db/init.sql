
-- add uuid-ossp extension
create extension "uuid-ossp";

-- create tables

create table "user" (
	"uuid" uuid primary key default uuid_generate_v4(),

	"nickname" varchar(60) not null,

	"email" varchar(60) not null unique,
	"username" varchar(45) not null unique,

	"password" text not null,

	"created_at" timestamp not null default now(),
	"updated_at" timestamp,
	"deleted_at" timestamp
);

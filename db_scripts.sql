
drop table csv_column_mappings;

drop table csv_file_configurations;

create table csv_file_configurations(
  id serial PRIMARY KEY,
  name varchar not null unique,
  file_pattern varchar not null,
  created_at TIMESTAMP WITH TIME ZONE not null,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE default null,
)

create table csv_column_mappings(
  id serial PRIMARY KEY,
  name varchar not null,
  ordinal int not null,
  type varchar not null,
  args jsonb not null default '{}'::jsonb,
  created_at TIMESTAMP WITH TIME ZONE not null,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE default null,
  file_configuration_id int not null,
  FOREIGN KEY (file_configuration_id) REFERENCES csv_file_configurations (id)
)

CREATE TABLE activity_sources(
   id serial PRIMARY KEY,
   name VARCHAR (255) UNIQUE NOT null,
   created_at timestamp not null,
   updated_at timestamp,
   deleted_at timestamp
);

create table activity_files(
  id serial primary key,
  name varchar not null,
  type varchar not null,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE default null,
  deleted_at TIMESTAMP WITH TIME ZONE default null,
  extension varchar,
  is_processed boolean default false,
  processed_at TIMESTAMP WITH TIME ZONE default null
)

create table activities(
	id serial primary key,
	transaction_date timestamp not null,
	posted_date timestamp not null,
	created_at timestamp not null,
	updated_at timestamp,
	deleted_at timestamp,
	description varchar,
	amount money not null,
	category varchar,
	source_id int,
	code varchar not null,
	UNIQUE(code, source_id),
	FOREIGN KEY (source_id)
      REFERENCES activity_sources (id)
);

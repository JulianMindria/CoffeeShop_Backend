CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT uuid_generate_v4(),
	username varchar(255) NULL DEFAULT NULL::character varying,
	email varchar(255) NOT NULL,
	pass varchar(255) NOT NULL,
	phone varchar(255) NOT NULL,
	image_file varchar NULL,
	roles varchar NOT NULL DEFAULT 'user'::character varying,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id_user),
	CONSTRAINT users_un UNIQUE (username),
	CONSTRAINT users_username_key UNIQUE (username),
	CONSTRAINT users_username_key1 UNIQUE (username)
);
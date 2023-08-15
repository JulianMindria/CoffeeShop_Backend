CREATE TABLE public.products (
	id_product uuid NOT NULL DEFAULT uuid_generate_v4(),
	product_name varchar(255) NOT NULL,
	description varchar NOT NULL,
	stock int4 NOT NULL DEFAULT 0,
	price int4 NOT NULL,
	image_file varchar NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	categories varchar NULL DEFAULT 'food'::character varying,
	isfavorite bool NOT NULL DEFAULT false,
	CONSTRAINT product_pk PRIMARY KEY (id_product)
);
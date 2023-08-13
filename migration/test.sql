CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    username varchar(255) DEFAULT NULL,
    email varchar(255) NOT NULL,
    pass varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    user_image varchar DEFAULT NULL, 
    roles varchar NOT NULL DEFAULT 'user',
    created_at timestamp without time zone NULL DEFAULT now(),
	updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (id_user)
)


CREATE TABLE public.products (
    id_product uuid NULL DEFAULT uuid_generate_v4(),
    product_name varchar(255) NOT NULL,
    "description" varchar NOT NULL,
    stock int NOT NULL DEFAULT 0,
    price int NOT NULL,
    product_image varchar DEFAULT NULL,
    created_at timestamp without time zone NULL DEFAULT now(),
	updated_at timestamp without time zone NULL,
    CONSTRAINT product_pk PRIMARY KEY (id_product)
)


CREATE TABLE public.favorite_products (
    id_favorite_product uuid NULL DEFAULT uuid_generate_v4(),
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    id_product uuid NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NULL DEFAULT now(),
	updated_at timestamp without time zone NULL,
    CONSTRAINT fav_pk PRIMARY KEY (id_favorite_product),
    CONSTRAINT user_favorite_fk FOREIGN KEY (id_user) REFERENCES public.users(id_user) ON DELETE CASCADE,
    CONSTRAINT prod_favorite_fk FOREIGN KEY (id_product) REFERENCES public.products(id_product) ON DELETE CASCADE
)


INSERT INTO users(username,email,pass,phone) VALUES('julian', 'mindria', 'pass', '081')


SELECT * FROM users

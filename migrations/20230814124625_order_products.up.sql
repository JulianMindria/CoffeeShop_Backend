CREATE TABLE public.order_products (
	order_id uuid NOT NULL,
	id_products uuid NULL,
	id_user uuid NULL,
	status bool NULL DEFAULT false,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT fav_pk PRIMARY KEY (order_id),
	CONSTRAINT order_products_fk FOREIGN KEY (id_user) REFERENCES public.users(id_user) ON DELETE CASCADE,
	CONSTRAINT order_products_fk_1 FOREIGN KEY (id_products) REFERENCES public.products(id_product) ON DELETE CASCADE
);
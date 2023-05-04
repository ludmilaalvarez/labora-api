
CREATE TABLE IF NOT EXISTS public.items
(
    id serial,
    customer_name character varying(255) NOT NULL,
    order_date date NOT NULL,
    product character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price numeric NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (id)
)


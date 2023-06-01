CREATE TABLE IF NOT EXISTS public.wallets
(
    id SERIAL,
    person_id CHARACTER VARYING(25) NOT NULL,
	date DATE NOT NULL,
    country CHARACTER VARYING(30) NOT NULL,
    state CHARACTER VARYING(25) NOT NULL,
	amount DOUBLE PRECISION NOT NULL,
    name CHARACTER VARYING (50) NOT NULL,
    CONSTRAINT wallets_pkey PRIMARY KEY (id)
);



CREATE TABLE IF NOT EXISTS public.solicitud
(
    id SERIAL,
    state CHARACTER VARYING(25) NOT NULL,
    date DATE NOT NULL,
    status CHARACTER VARYING(25) NOT NULL,
	person_id CHARACTER VARYING(25) NOT NULL,
	country CHARACTER VARYING (30) NOT NULL,
	wallet_id INTEGER,
    type_transaction CHARACTER VARYING (30) NOT NULL,
	
    CONSTRAINT solicitud_pkey PRIMARY KEY (id),
	CONSTRAINT FK_solicitud_wallets FOREIGN KEY (wallet_id) REFERENCES wallets (id)
);


CREATE TABLE IF NOT EXISTS public.transaction
(
    id SERIAL,
    amount double precision NOT NULL,
    type character varying(25) NOT NULL,
    date date NOT NULL,
    sender_wallet_id integer,
    receiver_wallet_id integer,
    sender_id character varying(25),
    receiver_id character(25),
    CONSTRAINT transaction_pkey PRIMARY KEY (id),
    CONSTRAINT fk_solicitud_wallets FOREIGN KEY (sender_wallet_id) REFERENCES public.wallets (id)
);


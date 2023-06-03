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
    sender_id CHARACTER VARYING (20),
	receiver_id CHARACTER VARYING (20),
    amount double precision NOT NULL,
    type character varying(25) NOT NULL,
    date date NOT NULL
);

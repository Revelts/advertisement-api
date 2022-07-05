CREATE TABLE public.company (
    company_id serial PRIMARY KEY,
    company_name VARCHAR(255),
    company_balance numeric
)

CREATE TABLE public.ads (
    advertisement_id serial PRIMARY KEY,
    advertisement_owner int default 0 NOT NULL,
    advertisement_name VARCHAR(255),
    advertisement_category int,
    advertisement_baseprice numeric,
    advertisement_created_by int,
    advertisement_created_at timestamp without time zone default now(),
    advertisement_updated_at timestamp without time zone default now(),
)

CREATE TABLE public.log (
    id serial PRIMARY KEY,
    company_id int,
    advertisement_id int,
    type int,
    cost numeric
)

CREATE TABLE public.transaction (
   trx_id serial PRIMARY KEY,
   advertisement_id int,
   company_id int,
   start_advertising timestamp without time zone,
   end_advertising timestamp without time zone,
   advertisement_location VARCHAR(255),
   fee numeric
)

CREATE TABLE public.transaction_status (
   trx_id int,
   status int,
   created_at timestamp without time zone default now()
)
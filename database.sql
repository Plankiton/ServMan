--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: check_validate_token(); Type: FUNCTION; Schema: public; Owner: plankiton
--

CREATE FUNCTION public.check_validate_token() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN TRUNC(DATE_PART('day', tg_argv[0].last_log_time::timestamp - NOW()::timestamp))>=5;
END;
$$;


ALTER FUNCTION public.check_validate_token() OWNER TO plankiton;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addrs; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.addrs (
    id text NOT NULL,
    street text,
    state text,
    number text,
    code text,
    city text,
    neightbourn text
);


ALTER TABLE public.addrs OWNER TO plankiton;

--
-- Name: farms; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.farms (
    id text NOT NULL,
    person_id text,
    address_id text,
    name text,
    create_time timestamp with time zone,
    update_time timestamp with time zone
);


ALTER TABLE public.farms OWNER TO plankiton;

--
-- Name: people; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.people (
    id text NOT NULL,
    doc_value text,
    doc_type text DEFAULT 'cpf'::text,
    phone text,
    name text,
    pass_hash text,
    create_time timestamp with time zone,
    update_time timestamp with time zone,
    telephone text,
    type text
);


ALTER TABLE public.people OWNER TO plankiton;

--
-- Name: role_ships; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.role_ships (
    role_id bigint NOT NULL,
    person_id text
);


ALTER TABLE public.role_ships OWNER TO plankiton;

--
-- Name: role_ships_role_id_seq; Type: SEQUENCE; Schema: public; Owner: plankiton
--

CREATE SEQUENCE public.role_ships_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.role_ships_role_id_seq OWNER TO plankiton;

--
-- Name: role_ships_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: plankiton
--

ALTER SEQUENCE public.role_ships_role_id_seq OWNED BY public.role_ships.role_id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.roles OWNER TO plankiton;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: plankiton
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO plankiton;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: plankiton
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: servs; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.servs (
    id text NOT NULL,
    description text,
    price numeric,
    begin_photo text,
    end_photo text,
    employee_id text,
    farm_id text,
    begin_time timestamp with time zone,
    end_time timestamp with time zone,
    stoped boolean
);


ALTER TABLE public.servs OWNER TO plankiton;

--
-- Name: tokens; Type: TABLE; Schema: public; Owner: plankiton
--

CREATE TABLE public.tokens (
    id text NOT NULL,
    person_id text,
    create_time timestamp with time zone,
    last_log_time timestamp with time zone
);


ALTER TABLE public.tokens OWNER TO plankiton;

--
-- Name: role_ships role_id; Type: DEFAULT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.role_ships ALTER COLUMN role_id SET DEFAULT nextval('public.role_ships_role_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Data for Name: addrs; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.addrs (id, street, state, number, code, city, neightbourn) FROM stdin;
\.


--
-- Data for Name: farms; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.farms (id, person_id, address_id, name, create_time, update_time) FROM stdin;
\.


--
-- Data for Name: people; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.people (id, doc_value, doc_type, phone, name, pass_hash, create_time, update_time, telephone, type) FROM stdin;
b6589fc6ab0dc82cf12099d1c2d40ab994e8410c	0	cpf	0	ROOT	$2a$04$ON2Fs0piZiyQjmOI/gwedeb.gEo9194oNL8/gbRJsd3iSeayNI/eS	2021-01-19 12:29:02.634474-03	2021-01-19 12:29:02.634474-03	\N	\N
\.


--
-- Data for Name: role_ships; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.role_ships (role_id, person_id) FROM stdin;
8	b6589fc6ab0dc82cf12099d1c2d40ab994e8410c
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.roles (id, name) FROM stdin;
5	root
6	employee
7	client
8	admin
\.


--
-- Data for Name: servs; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.servs (id, description, price, begin_photo, end_photo, employee_id, farm_id, begin_time, end_time, stoped) FROM stdin;
\.


--
-- Data for Name: tokens; Type: TABLE DATA; Schema: public; Owner: plankiton
--

COPY public.tokens (id, person_id, create_time, last_log_time) FROM stdin;
\.


--
-- Name: role_ships_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: plankiton
--

SELECT pg_catalog.setval('public.role_ships_role_id_seq', 1, false);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: plankiton
--

SELECT pg_catalog.setval('public.roles_id_seq', 8, true);


--
-- Name: addrs addrs_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.addrs
    ADD CONSTRAINT addrs_pkey PRIMARY KEY (id);


--
-- Name: farms farms_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.farms
    ADD CONSTRAINT farms_pkey PRIMARY KEY (id);


--
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


--
-- Name: role_ships role_ships_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.role_ships
    ADD CONSTRAINT role_ships_pkey PRIMARY KEY (role_id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: servs servs_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.servs
    ADD CONSTRAINT servs_pkey PRIMARY KEY (id);


--
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: plankiton
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (id);


--
-- Name: idx_addrs_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_addrs_id ON public.addrs USING btree (id);


--
-- Name: idx_farms_address_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE UNIQUE INDEX idx_farms_address_id ON public.farms USING btree (address_id);


--
-- Name: idx_farms_create_time; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_farms_create_time ON public.farms USING btree (create_time);


--
-- Name: idx_farms_name; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_farms_name ON public.farms USING btree (name);


--
-- Name: idx_farms_update_time; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_farms_update_time ON public.farms USING btree (update_time);


--
-- Name: idx_people_doc_value; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE UNIQUE INDEX idx_people_doc_value ON public.people USING btree (doc_value);


--
-- Name: idx_people_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_people_id ON public.people USING btree (id);


--
-- Name: idx_people_name; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_people_name ON public.people USING btree (name);


--
-- Name: idx_people_type; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_people_type ON public.people USING btree (type);


--
-- Name: idx_role_ships_person_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_role_ships_person_id ON public.role_ships USING btree (person_id);


--
-- Name: idx_roles_name; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_roles_name ON public.roles USING btree (name);


--
-- Name: idx_servs_begin_time; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_servs_begin_time ON public.servs USING btree (begin_time);


--
-- Name: idx_servs_employee_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_servs_employee_id ON public.servs USING btree (employee_id);


--
-- Name: idx_servs_end_time; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_servs_end_time ON public.servs USING btree (end_time);


--
-- Name: idx_servs_farm_id; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_servs_farm_id ON public.servs USING btree (farm_id);


--
-- Name: idx_servs_price; Type: INDEX; Schema: public; Owner: plankiton
--

CREATE INDEX idx_servs_price ON public.servs USING btree (price);


--
-- PostgreSQL database dump complete
--

INSERT INTO people(id, doc_value, doc_type, phone, name, pass_hash)
VALUES('b6589fc6ab0dc82cf12099d1c2d40ab994e8410c', '0', 'cpf', '0', 'root', '$2a$04$ON2Fs0piZiyQjmOI/gwedeb.gEo9194oNL8/gbRJsd3iSeayNI/eS');

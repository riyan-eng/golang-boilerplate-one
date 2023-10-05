--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: example; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.example (
    id integer NOT NULL,
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    nama character varying(255),
    detail text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.example OWNER TO riyan;

--
-- Name: example_id_seq; Type: SEQUENCE; Schema: public; Owner: riyan
--

CREATE SEQUENCE public.example_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.example_id_seq OWNER TO riyan;

--
-- Name: example_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: riyan
--

ALTER SEQUENCE public.example_id_seq OWNED BY public.example.id;


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.permissions (
    p_type character varying(32) DEFAULT ''::character varying NOT NULL,
    v0 character varying(255) DEFAULT ''::character varying NOT NULL,
    v1 character varying(255) DEFAULT ''::character varying NOT NULL,
    v2 character varying(255) DEFAULT ''::character varying NOT NULL,
    v3 character varying(255) DEFAULT ''::character varying NOT NULL,
    v4 character varying(255) DEFAULT ''::character varying NOT NULL,
    v5 character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.permissions OWNER TO riyan;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    kode character varying(255),
    nama character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.roles OWNER TO riyan;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: riyan
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO riyan;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: riyan
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO riyan;

--
-- Name: user_datas; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.user_datas (
    id integer NOT NULL,
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    nama character varying(255),
    nik character varying(255),
    nomor_telepon character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.user_datas OWNER TO riyan;

--
-- Name: user_datas_id_seq; Type: SEQUENCE; Schema: public; Owner: riyan
--

CREATE SEQUENCE public.user_datas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_datas_id_seq OWNER TO riyan;

--
-- Name: user_datas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: riyan
--

ALTER SEQUENCE public.user_datas_id_seq OWNED BY public.user_datas.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: riyan
--

CREATE TABLE public.users (
    id integer NOT NULL,
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255),
    role character varying(255),
    password text,
    user_data uuid,
    is_aktif boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO riyan;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: riyan
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO riyan;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: riyan
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: example id; Type: DEFAULT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.example ALTER COLUMN id SET DEFAULT nextval('public.example_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: user_datas id; Type: DEFAULT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.user_datas ALTER COLUMN id SET DEFAULT nextval('public.user_datas_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: example example_pkey; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.example
    ADD CONSTRAINT example_pkey PRIMARY KEY (id);


--
-- Name: example example_uuid_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.example
    ADD CONSTRAINT example_uuid_key UNIQUE (uuid);


--
-- Name: roles roles_kode_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_kode_key UNIQUE (kode);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: roles roles_uuid_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_uuid_key UNIQUE (uuid);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: user_datas user_datas_nik_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.user_datas
    ADD CONSTRAINT user_datas_nik_key UNIQUE (nik);


--
-- Name: user_datas user_datas_nomor_telepon_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.user_datas
    ADD CONSTRAINT user_datas_nomor_telepon_key UNIQUE (nomor_telepon);


--
-- Name: user_datas user_datas_pkey; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.user_datas
    ADD CONSTRAINT user_datas_pkey PRIMARY KEY (id);


--
-- Name: user_datas user_datas_uuid_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.user_datas
    ADD CONSTRAINT user_datas_uuid_key UNIQUE (uuid);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_uuid_key; Type: CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_uuid_key UNIQUE (uuid);


--
-- Name: idx_permissions; Type: INDEX; Schema: public; Owner: riyan
--

CREATE INDEX idx_permissions ON public.permissions USING btree (p_type, v0, v1);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: riyan
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users fk_role; Type: FK CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_role FOREIGN KEY (role) REFERENCES public.roles(kode) ON DELETE SET NULL;


--
-- Name: users fk_user_data; Type: FK CONSTRAINT; Schema: public; Owner: riyan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_user_data FOREIGN KEY (user_data) REFERENCES public.user_datas(uuid) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--


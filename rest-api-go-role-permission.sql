--
-- PostgreSQL database dump
--

-- Dumped from database version 12.14
-- Dumped by pg_dump version 12.14

-- Started on 2025-10-01 21:56:16

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 2866 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 18624)
-- Name: permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.permissions (
    id character varying(36) NOT NULL,
    parent_menu character varying(20),
    parent_id character varying(36),
    name character varying(20) NOT NULL,
    alias character varying(20),
    url character varying(50),
    icon character varying(20),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.permissions OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 18665)
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refresh_tokens (
    id character varying(36) NOT NULL,
    revoked boolean,
    expired_at timestamp with time zone,
    user_id character varying(36) NOT NULL
);


ALTER TABLE public.refresh_tokens OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 18631)
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_permissions (
    role_id character varying(36) NOT NULL,
    permission_id character varying(36) NOT NULL
);


ALTER TABLE public.role_permissions OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 18615)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id character varying(36) NOT NULL,
    name character varying(100) NOT NULL,
    desctiption character varying(255) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 18646)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(36) NOT NULL,
    name character varying(100) NOT NULL,
    username character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255),
    social_id character varying(100),
    provider character varying(100),
    avatar character varying(200),
    is_superadmin boolean DEFAULT false,
    role_id character varying(36),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 2857 (class 0 OID 18624)
-- Dependencies: 203
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.permissions (id, parent_menu, parent_id, name, alias, url, icon, created_at, updated_at, deleted_at) FROM stdin;
3f2b44b7-5778-435e-abe7-6efcccd509c5			role-delete	Delete Role			\N	\N	\N
05c717c6-6375-406f-95bb-b78e76217a79			user-update	Edit User			\N	\N	\N
a67219b9-6c36-43e7-9166-5d0fc52e1508			user-delete	Delete User			\N	\N	\N
fdc62d2a-fde5-439c-8814-0a31edee4f37			user-index	User	user		\N	\N	\N
04e36e80-9e32-4292-80ed-27f3b3cda7db			role-show	Show Role			\N	\N	\N
a6bb9a24-efff-4b5f-ad89-7313a9a9e95b			role-update	Edit Role			\N	\N	\N
6392eb2c-3a06-45cd-b1d2-c710238a931f			user-show	Show User			\N	\N	\N
fe0b848b-0b7e-438c-a9f2-903127cebe3e			role-index	Role	role		\N	\N	\N
5a153313-4d93-4722-b5d1-857c7512de12			role-store	Create Role			\N	\N	\N
292a923e-65f2-4ab1-89aa-b0d96f435c8b			user-store	Create User			\N	\N	\N
125f06ec-6445-40fa-86a2-62f82f63442a			test	test			2025-09-29 15:54:27.703713+07	2025-09-29 15:54:27.703713+07	\N
\.


--
-- TOC entry 2860 (class 0 OID 18665)
-- Dependencies: 206
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refresh_tokens (id, revoked, expired_at, user_id) FROM stdin;
\.


--
-- TOC entry 2858 (class 0 OID 18631)
-- Dependencies: 204
-- Data for Name: role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.role_permissions (role_id, permission_id) FROM stdin;
6f55ddd2-aa15-473b-bc67-a48a26e25f01	3f2b44b7-5778-435e-abe7-6efcccd509c5
6f55ddd2-aa15-473b-bc67-a48a26e25f01	05c717c6-6375-406f-95bb-b78e76217a79
6f55ddd2-aa15-473b-bc67-a48a26e25f01	a67219b9-6c36-43e7-9166-5d0fc52e1508
6f55ddd2-aa15-473b-bc67-a48a26e25f01	fdc62d2a-fde5-439c-8814-0a31edee4f37
6f55ddd2-aa15-473b-bc67-a48a26e25f01	04e36e80-9e32-4292-80ed-27f3b3cda7db
6f55ddd2-aa15-473b-bc67-a48a26e25f01	a6bb9a24-efff-4b5f-ad89-7313a9a9e95b
6f55ddd2-aa15-473b-bc67-a48a26e25f01	6392eb2c-3a06-45cd-b1d2-c710238a931f
6f55ddd2-aa15-473b-bc67-a48a26e25f01	fe0b848b-0b7e-438c-a9f2-903127cebe3e
6f55ddd2-aa15-473b-bc67-a48a26e25f01	5a153313-4d93-4722-b5d1-857c7512de12
6f55ddd2-aa15-473b-bc67-a48a26e25f01	292a923e-65f2-4ab1-89aa-b0d96f435c8b
\.


--
-- TOC entry 2856 (class 0 OID 18615)
-- Dependencies: 202
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, name, desctiption, created_at, updated_at, deleted_at) FROM stdin;
6f55ddd2-aa15-473b-bc67-a48a26e25f01	Internal Admin		2025-09-29 14:09:40.509757+07	2025-09-29 14:16:19.074692+07	\N
d4889a8a-d651-424e-8b10-e1a9bca9dfb8	Test		2025-09-29 15:08:00.943867+07	2025-09-29 15:08:00.943867+07	2025-09-29 15:09:12.403724+07
\.


--
-- TOC entry 2859 (class 0 OID 18646)
-- Dependencies: 205
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, username, email, password, social_id, provider, avatar, is_superadmin, role_id, created_at, updated_at, deleted_at) FROM stdin;
031d38b5-4ed7-4e01-b112-a7ce08ecc593	Super Admin	superadmin	superadmin@mail.com	$2a$04$vAazCfjre6j6q3lZK10/Uu6tYXw/uGphi00o43trEhTYoq/2IpwLK				t	\N	2025-09-29 14:04:44.216684+07	2025-09-29 14:04:44.216684+07	\N
\.


--
-- TOC entry 2712 (class 2606 OID 18628)
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 2725 (class 2606 OID 18669)
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 2714 (class 2606 OID 18635)
-- Name: role_permissions role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_pkey PRIMARY KEY (role_id, permission_id);


--
-- TOC entry 2707 (class 2606 OID 18619)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 2721 (class 2606 OID 18654)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2708 (class 1259 OID 18677)
-- Name: idx_permissions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_permissions_deleted_at ON public.permissions USING btree (deleted_at);


--
-- TOC entry 2709 (class 1259 OID 18630)
-- Name: idx_permissions_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_permissions_id ON public.permissions USING btree (id);


--
-- TOC entry 2710 (class 1259 OID 18629)
-- Name: idx_permissions_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_permissions_name ON public.permissions USING btree (name);


--
-- TOC entry 2722 (class 1259 OID 18675)
-- Name: idx_refresh_tokens_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_refresh_tokens_id ON public.refresh_tokens USING btree (id);


--
-- TOC entry 2723 (class 1259 OID 18676)
-- Name: idx_refresh_tokens_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);


--
-- TOC entry 2703 (class 1259 OID 18623)
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- TOC entry 2704 (class 1259 OID 18622)
-- Name: idx_roles_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_roles_id ON public.roles USING btree (id);


--
-- TOC entry 2705 (class 1259 OID 18621)
-- Name: idx_roles_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_roles_name ON public.roles USING btree (name);


--
-- TOC entry 2715 (class 1259 OID 18660)
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- TOC entry 2716 (class 1259 OID 18662)
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- TOC entry 2717 (class 1259 OID 18664)
-- Name: idx_users_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_id ON public.users USING btree (id);


--
-- TOC entry 2718 (class 1259 OID 18661)
-- Name: idx_users_role_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_role_id ON public.users USING btree (role_id);


--
-- TOC entry 2719 (class 1259 OID 18663)
-- Name: idx_users_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);


--
-- TOC entry 2729 (class 2606 OID 18670)
-- Name: refresh_tokens fk_refresh_tokens_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT fk_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2727 (class 2606 OID 18641)
-- Name: role_permissions fk_role_permissions_permission; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES public.permissions(id);


--
-- TOC entry 2726 (class 2606 OID 18636)
-- Name: role_permissions fk_role_permissions_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 2728 (class 2606 OID 18655)
-- Name: users fk_users_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE SET NULL;


-- Completed on 2025-10-01 21:56:17

--
-- PostgreSQL database dump complete
--


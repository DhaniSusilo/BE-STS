--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

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
-- Name: members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.members (
    nik character varying(16) NOT NULL,
    nama character varying(100) NOT NULL,
    no_hp character varying(20) NOT NULL,
    provinsi character varying(100) NOT NULL,
    kabupaten character varying(100) NOT NULL,
    kecamatan character varying(100) NOT NULL,
    kelurahan character varying(100) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.members OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(255) NOT NULL,
    username text,
    first_name text,
    last_name text,
    password text,
    level text,
    "for" text,
    enabled bigint DEFAULT 1,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: members; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.members (nik, nama, no_hp, provinsi, kabupaten, kecamatan, kelurahan, created_at) FROM stdin;
3201010101010001	Ahmad Santoso	081234567890	Jawa Barat	Bandung	Coblong	Dago	2025-07-04 22:47:12.857498+07
3201010101010002	Ahmad Santoso 2	081234567891	Jawa Barat	Bandung	Coblong	Lebakgede	2025-07-04 22:47:12.857498+07
3201010101010003	Ahmad Santoso 3	081234567892	Jawa Barat	Bandung	Coblong	Cicadas	2025-07-04 22:47:12.857498+07
3201010101010004	Ahmad Santoso 4	081234567893	Jawa Barat	Bandung	Coblong	Sukajadi	2025-07-04 22:47:12.857498+07
3201010101010005	Ahmad Santoso 5	081234567894	Jawa Barat	Bandung	Coblong	Cidadap	2025-07-04 22:47:12.857498+07
3201010101010006	Ahmad Santoso 6	081234567895	Jawa Barat	Bandung	Sukasari	Cicendo	2025-07-04 22:47:12.857498+07
3201010101010007	Ahmad Santoso 7	081234567896	Jawa Barat	Bandung	Sukasari	Sukamiskin	2025-07-04 22:47:12.857498+07
3201010101010008	Ahmad Santoso 8	081234567897	Jawa Barat	Bandung	Sukasari	Kebonwaruk	2025-07-04 22:47:12.857498+07
3201010101010009	Ahmad Santoso 9	081234567898	Jawa Barat	Bandung	Sukasari	Mekarwangi	2025-07-04 22:47:12.857498+07
3201010101010010	Ahmad Santoso 10	081234567899	Jawa Barat	Bandung	Sukasari	Pasir Kaliki	2025-07-04 22:47:12.857498+07
3201010101010011	Ahmad Santoso 11	081234567900	Jawa Barat	Bandung	Andir	Cijerah	2025-07-04 22:47:12.857498+07
3201010101010012	Ahmad Santoso 12	081234567901	Jawa Barat	Bandung	Andir	Kebon Jeruk	2025-07-04 22:47:12.857498+07
3201010101010013	Ahmad Santoso 13	081234567902	Jawa Barat	Bandung	Andir	Cikutra	2025-07-04 22:47:12.857498+07
3201010101010014	Ahmad Santoso 14	081234567903	Jawa Barat	Bandung	Andir	Padasuka	2025-07-04 22:47:12.857498+07
3201010101010015	Ahmad Santoso 15	081234567904	Jawa Barat	Bandung	Andir	Sukamiskin	2025-07-04 22:47:12.857498+07
3201010101010016	Ahmad Santoso 16	081234567905	Jawa Barat	Bandung	Batununggal	Batununggal	2025-07-04 22:47:12.857498+07
3201010101010017	Ahmad Santoso 17	081234567906	Jawa Barat	Bandung	Batununggal	Cijawura	2025-07-04 22:47:12.857498+07
3201010101010018	Ahmad Santoso 18	081234567907	Jawa Barat	Bandung	Batununggal	Cibeunying	2025-07-04 22:47:12.857498+07
3201010101010019	Ahmad Santoso 19	081234567908	Jawa Barat	Bandung	Batununggal	Ujung Berung	2025-07-04 22:47:12.857498+07
3201010101010020	Ahmad Santoso 20	081234567909	Jawa Barat	Bandung	Batununggal	Sukaluyu	2025-07-04 22:47:12.857498+07
3201010101010021	Ahmad Santoso 21	081234567910	Jawa Barat	Bandung	Mandalajati	Cibiru	2025-07-04 22:47:12.857498+07
3201010101010022	Ahmad Santoso 22	081234567911	Jawa Barat	Bandung	Mandalajati	Mandalajati	2025-07-04 22:47:12.857498+07
3201010101010023	Ahmad Santoso 23	081234567912	Jawa Barat	Bandung	Mandalajati	Cipadung	2025-07-04 22:47:12.857498+07
3201010101010024	Ahmad Santoso 24	081234567913	Jawa Barat	Bandung	Mandalajati	Panyileukan	2025-07-04 22:47:12.857498+07
3201010101010025	Ahmad Santoso 25	081234567914	Jawa Barat	Bandung	Mandalajati	Cisaranten Kulon	2025-07-04 22:47:12.857498+07
3201010101030001	Sudidik	081234567890	Jawa Barat	Bandung	Coblong	Dago	2025-07-04 22:49:43.861886+07
3604111805010002	dhanisusilo	0895336545219	Jawa Tengah	Semarang	Candisari	Gunung Pati	2025-07-04 23:23:02.389963+07
3604111805010004	dhanisusilo	0895336545219	Jawa Barat	Bandung	Sukasari	Pasir Kaliki	2025-07-04 23:23:52.829369+07
3604111805010102	dhanisusilo	0895336545219	Jawa Barat	Bogor	Cibinong	Gunung Putri	2025-07-07 06:59:23.469639+07
1111111111111111	user	0895336545219	Aceh	Aceh Besar	Kuta Baro	Lampasi	2025-07-07 07:08:54.344238+07
1111111111111112	user2	0895336545219	Sumatera Utara	Deli Serdang	Lubuk Pakam	Bakaran Batu	2025-07-07 07:09:15.132642+07
1111111111111113	user3	0895336545219	DKI Jakarta	Kepulauan Seribu	Pulau Kelapa	Pulau Kelapa	2025-07-07 07:09:40.286372+07
1111111111111114	user4	0895336545219	Kalimantan Timur	Kutai Kartanegara	Tenggarong	Mangkurawang	2025-07-07 07:10:01.651908+07
1111111111111115	user5	0895336545219	Papua	Jayapura	Abepura	Vim	2025-07-07 07:10:18.265702+07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, first_name, last_name, password, level, "for", enabled, created_at, updated_at) FROM stdin;
11682f9c-9bdd-4e60-90ee-955a3ec58d04	didikSudidik	Didik	Sudidik	$2a$10$r1q3EI9s7SUYWx4BUxtOAuEx8IWdoltXHEKCfV8f1yWwPzOobcoYq	Provinsi	Jawa Barat	1	2025-07-04 21:07:52.171426+07	2025-07-04 21:07:52.171426+07
7ab7a2e0-dd24-4ee9-89f7-2e339390aab2	AdminPusat	Didik	Sudidik	$2a$10$Ojc8nL7vxy3njfUy8AIpYeF4dDlxHNh1AuN518hY.BJbEW0y6uTz.	Pusat		1	2025-07-04 22:52:43.793086+07	2025-07-04 22:52:43.793086+07
f1771e46-5aa7-47d7-8abc-fb55a060390a	AdminProvinsi	Didik	Sudidik	$2a$10$mPvdyuJtyDA0yRd6mSwpKuQgm418/mLfcnse/tIJtet8gzMKSNJ5C	Provinsi	Jawa Barat	1	2025-07-04 22:53:24.119412+07	2025-07-04 22:53:24.119412+07
43ac8f78-e825-4c3c-b2ef-9f0d509afeca	AdminKabupaten	Didik	Sudidik	$2a$10$7TDCH1MTauLaCTuPp/ARG.nCAP0utLzH/ILroXht4RcsAaF6soZEa	Kabupaten	Bandung	1	2025-07-04 22:54:04.388686+07	2025-07-04 22:54:04.388686+07
4a6e49b4-cfee-49b7-989c-5483d82d19e0	AdminKecamatan	Didik	Sudidik	$2a$10$CeFUZphWGYHk0xoZ4gUtGuA2sUBpJjilIkQBhpM1Mh1wLb6w19Vn2	Kecamatan	Coblong	1	2025-07-04 22:54:25.874129+07	2025-07-04 22:54:25.874129+07
5b340d41-cd48-4604-9a3b-fd15631cc23a	AdminKelurahan	Didik	Sudidik	$2a$10$u02UXiXxVcbSbdwunF4hD.wE0h/ZnuAw/mdvkqhGuiUd0aDJL2/ZK	Kelurahan	Dago	1	2025-07-04 22:54:42.594311+07	2025-07-04 22:54:42.594311+07
916f79a8-299b-4add-822e-c4eedf6ef495	admin1	admin	admin	$2a$10$aCQ36VuZfecOJUXzW.LLXejz59CHgHOFGeI8M.x8WU.kMCFfvFPhW	Pusat		1	2025-07-07 00:16:12.663066+07	2025-07-07 00:16:33.471798+07
\.


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_members_nik; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_members_nik ON public.members USING btree (nik);


--
-- Name: idx_users_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);


--
-- PostgreSQL database dump complete
--


--
-- PostgreSQL database dump
--

-- Dumped from database version 9.3.13
-- Dumped by pg_dump version 9.5.3

-- Started on 2016-06-15 01:10:28 BRT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 11756)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 1943 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 171 (class 1259 OID 16387)
-- Name: wiki; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE wiki (
    name character varying(30) NOT NULL,
    description text
);


ALTER TABLE wiki OWNER TO root;

--
-- TOC entry 1935 (class 0 OID 16387)
-- Dependencies: 171
-- Data for Name: wiki; Type: TABLE DATA; Schema: public; Owner: root
--

COPY wiki (name, description) FROM stdin;
\.


--
-- TOC entry 1827 (class 2606 OID 16394)
-- Name: name; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY wiki
    ADD CONSTRAINT name PRIMARY KEY (name);


--
-- TOC entry 1942 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2016-06-15 01:10:28 BRT

--
-- PostgreSQL database dump complete
--


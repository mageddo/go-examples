CREATE TABLE wiki (
    name character varying(30) NOT NULL,
    description text
);


INSERT INTO wiki ("name","description") VALUES ('name1', 'name1 test');

SELECT * FROM wiki;
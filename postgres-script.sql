CREATE TABLE IF NOT EXISTS public."Cars"
(
    "Make" text COLLATE pg_catalog."default" NOT NULL,
    "Model" text COLLATE pg_catalog."default" NOT NULL,
    "Color" text COLLATE pg_catalog."default" NOT NULL,
    "Power" text COLLATE pg_catalog."default" NOT NULL,
    "Year" text COLLATE pg_catalog."default" NOT NULL,
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    CONSTRAINT "Cars_pkey" PRIMARY KEY (id),
    CONSTRAINT "Cars_id_id1_key" UNIQUE (id)
        INCLUDE(id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Cars"
    OWNER to postgres;

INSERT INTO public."Cars"(
	"Make", "Model", "Color", "Power", "Year")
	VALUES ('BMW', 'M3 GTR', 'Silver', '444 BHP', '2005'),
	('Lamborghini', 'Gallardo', 'Orange', '493 BHP', '2005'),
	('McLaren', 'F1', 'Yellow', '618 BHP', '2004'),
	('Mercedes-Benz', 'SLR McLaren', 'Black', '617 BHP', '2006');
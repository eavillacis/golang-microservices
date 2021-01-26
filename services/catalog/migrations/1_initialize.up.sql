-- Table: public.brands
CREATE TABLE public.brands
(
	"id" uuid NOT NULL,
	"name" character varying(45) NOT NULL,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	"deleted_at" timestamp with time zone,
	CONSTRAINT "brands_pkey" PRIMARY KEY ("id")
)
WITH (OIDS = FALSE)
TABLESPACE pg_default;

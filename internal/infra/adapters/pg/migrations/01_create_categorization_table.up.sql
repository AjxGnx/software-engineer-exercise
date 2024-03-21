CREATE TABLE IF NOT EXISTS public.categorization
(
    id          serial       NOT NULL,
    "number"    bigint       NOT NULL,
    category    varchar(10)  NOT NULL,

    CONSTRAINT categorization_pkey PRIMARY KEY (id),
    CONSTRAINT unique_categorization UNIQUE ("number",category)
);

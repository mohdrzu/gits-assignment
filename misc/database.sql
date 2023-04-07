-- DROP TABLE public.authors;

CREATE TABLE public.authors (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NULL,
	email text NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_authors_deleted_at ON public.authors USING btree (deleted_at);


-- DROP TABLE public.books;

CREATE TABLE public.books (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	title text NULL,
	"year" int8 NULL,
	author_id int8 NULL,
	publisher_id int8 NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_books_deleted_at ON public.books USING btree (deleted_at);


-- DROP TABLE public.publishers;

CREATE TABLE public.publishers (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NULL,
	"location" text NULL,
	book_id int8 NULL,
	CONSTRAINT idx_publishers_book_id UNIQUE (book_id),
	CONSTRAINT publishers_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_publishers_deleted_at ON public.publishers USING btree (deleted_at);



-- DROP TABLE public.users;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	email text NULL,
	"password" text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

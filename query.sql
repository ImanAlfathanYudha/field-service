CREATE TABLE public.field (
    id bigint PRIMARY KEY,
    uuid text NOT NULL,
    code VARCHAR(15) NOT NULL,
    name VARCHAR(100) NOT NULL,
    price_per_hour INT NOT NULL,
    images TEXT[] NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

 field_id INT NOT NULL,
    time_id INT NOT NULL,
    date DATE NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ


CREATE TABLE public.time (
      id bigint PRIMARY KEY,
    uuid text NOT NULL,
    start_time TIME WITHOUT TIME ZONE NOT NULL,
    end_time TIME WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE public.field_schedule (
    id bigint PRIMARY KEY,
    uuid text NOT NULL,
    field_id INT NOT NULL,
    time_id INT NOT NULL,
    date DATE NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
)
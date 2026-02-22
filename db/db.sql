-- schema prototype

create table salon.orders (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    customer_name text not null,
    subtotal numeric(12,0) not null,
    tax numeric(12,0) not null,
    tax_percentage numeric(5,2) not null,
    discount numeric(12,0) not null,
    total numeric(12,0) not null,
    payment_method_id int foreign key not null,
    constraint fk_orders_payment_methods foreign key (payment_method_id) references salon.payment_methods
);

create table salon.order_details (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    order_id int foreign key not null,
    service_id int foreign key not null,
    worker_id int foreign key not null,
    service_price numeric(12,0) not null,
    constraint fk_order_details_orders foreign key (order_id) references salon.orders,
    constraint fk_order_details_services foreign key (service_id) references salon.services,
    constraint fk_order_details_workers foreign key (worker_id) references salon.workers
);

create table salon.commisions (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    order_detail_id int foreign key not null,
    fee numeric(12,0) not null,
    constraint fk_commisions_order_details foreign key (order_detail_id) references salon.order_details
);

create table salon.services (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    name text not null,
    price numeric(12,0) not null,
    category_id text not null,
    is_active bool not null,
    constraint fk_services_service_categories foreign key (category_id) references salon.service_categories(id)
);

create table salon.service_categories (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    name text not null,
    is_active bool not null
);

create table salon.workers (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    name text not null,
    is_active bool not null
);

create table salon.payment_methods (
    id serial primary key not null,
    created_at timestamptz default now() not null,
    created_by text not null,
    updated_at timestamptz null,
    updated_by text null,
    name text not null,
    is_active bool not null
);

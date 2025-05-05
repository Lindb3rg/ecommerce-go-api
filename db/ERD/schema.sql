// Northwind Database Schema in DBML format for dbdiagram.io

Table categories {
  category_id smallint [pk, increment]
  category_name varchar(15) [not null]
  description text
  picture bytea
}

Table customer_demographics {
  customer_type_id char [pk]
  customer_desc text
}

Table customers {
  customer_id char [pk]
  company_name varchar(40) [not null]
  contact_name varchar(60)
  contact_title varchar(60)
  address varchar(60)
  city varchar(60)
  region varchar(60)
  postal_code varchar(10)
  country varchar(60)
  phone varchar(24)
  fax varchar(24)
  created_at timestamp [default: `NOW()`]
  active boolean [default: true]
}

Table customer_customer_demo {
  customer_id char [not null, ref: > customers.customer_id]
  customer_type_id char [not null, ref: > customer_demographics.customer_type_id]
  
  indexes {
    (customer_id, customer_type_id) [pk]
  }
}

Table employees {
  employee_id smallint [pk, increment]
  last_name varchar(60) [not null]
  first_name varchar(20) [not null]
  title varchar(60)
  title_of_courtesy varchar(25)
  birth_date date
  hire_date date
  address varchar(60)
  city varchar(60)
  region varchar(60)
  postal_code varchar(10)
  country varchar(60)
  home_phone varchar(24)
  extension varchar(4)
  photo bytea
  notes text
  reports_to smallint [ref: > employees.employee_id]
  photo_path varchar(255)
  created_at timestamp [default: `NOW()`]
  active boolean [default: true]
}

Table suppliers {
  supplier_id smallint [pk, increment]
  company_name varchar(40) [not null]
  contact_name varchar(60)
  contact_title varchar(60)
  address varchar(60)
  city varchar(60)
  region varchar(60)
  postal_code varchar(10)
  country varchar(60)
  phone varchar(24)
  fax varchar(24)
  homepage text
}

Table products {
  product_id smallint [pk, increment]
  product_name varchar(40) [not null]
  supplier_id smallint [ref: > suppliers.supplier_id]
  category_id smallint [ref: > categories.category_id]
  quantity_per_unit varchar(20)
  unit_price real
  units_in_stock smallint
  units_on_order smallint
  reorder_level smallint
  discontinued integer [not null]
}

Table region {
  region_id smallint [pk, increment]
  region_description char [not null]
}

Table shippers {
  shipper_id smallint [pk, increment]
  company_name varchar(40) [not null]
  phone varchar(24)
}

Table orders {
  order_id smallint [pk, increment]
  customer_id char [ref: > customers.customer_id]
  employee_id smallint [ref: > employees.employee_id]
  order_date date
  required_date date
  shipped_date date
  ship_via smallint [ref: > shippers.shipper_id]
  freight real
  ship_name varchar(40)
  ship_address varchar(60)
  ship_city varchar(30)
  ship_region varchar(30)
  ship_postal_code varchar(10)
  ship_country varchar(30)
}

Table territories {
  territory_id varchar(20) [pk]
  territory_description char [not null]
  region_id smallint [not null, ref: > region.region_id]
}

Table employee_territories {
  employee_id smallint [not null, ref: > employees.employee_id]
  territory_id varchar(20) [not null, ref: > territories.territory_id]
  
  indexes {
    (employee_id, territory_id) [pk]
  }
}

TABLE us_states {
    state_id smallint [pk, increment]
    state_name varchar(100)
    state_abbr varchar(2)
    state_region varchar(50)
}

Table order_details {
  order_id smallint [not null, ref: > orders.order_id]
  product_id smallint [not null, ref: > products.product_id]
  unit_price real [not null]
  quantity smallint [not null]
  discount real [not null]
  
  indexes {
    (order_id, product_id) [pk]
  }
}

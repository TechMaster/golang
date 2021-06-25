# Sử dụng GORM và MySQL

**Các quy ước mặc định của GORM**
- `ID` mặc định là cột primary key
- Tên bảng là danh từ tiếng Anh viết Hoa, kiểu Snake Case (ký tự đầu của từ viết hoa)
- Tên cột cũng viết kiểu Snake Case



Sửa lại tên bảng và cột theo chuẩn GORM
```
Table products {
  ID int [pk, increment]
  name varchar [not null]
  description varchar [not null]
  madein varchar(2) [ref: > countries.code]
  price int
  manufacturer int [ref: > manufacturers.ID]
}

Enum Property_Type {
  0 [note: "string"]
  1 [note: "integer"]
  2 [note: "float"]
  3 [note: "bool"]
  4 [note: "array"]
}

Table product_properies {
  ID int [pk, increment]
  product_id int [ref: > products.ID]
  key varchar [not null]
  value varchar [not null]
  type Property_Type
}

Table countries {
  code varchar(2) [pk, not null]
  name varchar [not null]
}

Table product_prices {
  ID int [pk, increment]
  product_id int [ref: > products.ID]
  price int
  created_at datetime [default: `now()`]
}

Table product_medias {
  ID int [pk, increment]
  product_id int [ref: > products.ID]
  uri varchar [not null]
  media_type media_types 
}

Enum media_types {
  photo
  vIDeo
  PDF
}

Table manufacturers {
  ID int [pk, increment]
  name varchar
  country varchar[2] [ref: > countries.code]
}

Table categories {
  ID int [pk, increment]
  name varchar [not null]
  parent_id int [ref: > categories.ID]
}

Table relate_products {
  product_id int [ref: > products.ID]
  relate_id int [ref: > products.ID]
  relation relate_types [not null]
}

Enum relate_types {
  oldversion
  newversion
  similar
  recommend
}

Table users {
  ID int [pk, increment]
  email varchar [not null, unique]
  mobile varchar [unique]
  password varchar [not null]
}

Table customers {
  ID int [pk, increment]
  user_id int [ref: - users.ID]
}

Table addresses {
  ID int [pk, increment]
  customer_id int [not null, ref: > customers.ID]
  addr varchar [not null]
  city_id int [ref: > cities.ID]
}

Table cities {
  ID int [pk, increment]
  name varchar [not null]
}
```
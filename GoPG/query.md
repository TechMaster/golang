# Thử nghiệm truy vấn

## 1. Liệt kê user và tất cả các role của user đó
User có bao nhiêu role thì sẽ có ngần đó dòng trả về. Nhìn không được gọn mắt lắm.

```sql
SELECT u."name", r.id, r."name" FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
```
|name|id|name|
|----|--|----|
|Rachel Cremin|4|SALE|
|Rachel Cremin|6|AUTHOR|
|Rachel Cremin|7|EDITOR|
|Sofia Keebler|7|EDITOR|
|Sofia Keebler|2|STUDENT|
|Sofia Keebler|1|ADMIN|


## 2. Gom các role của một user thành một mảng
Sử dụng hàm `array_agg` hàm này luôn phải đi cùng lệnh `GROUP BY`.
Nhược điểm của câu lệnh SELECT luôn chỉ trả về 2 cột: cột `group by` và cột `aggregate`
```sql
SELECT u."name", array_agg(r."name") FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
GROUP BY u."name";
```
|name|array_agg|
|----|---------|
|Verlie Donnelly|{TRAINER,MAINTAINER,ADMIN}|
|Taryn Reichert|{MAINTAINER,AUTHOR,ADMIN}|
|Stefanie Walter|{EMPLOYER,ADMIN,SALE}|
|Flavie VonRueden|{SALE,EDITOR,MAINTAINER}|
|Genevieve Mraz|{SALE,ADMIN,TRAINER}|
|Norma Crona|{STUDENT,SALE,MAINTAINER}|
|Ole Turner|{STUDENT,MAINTAINER,SALE}|
|Emilia Spinka|{EDITOR,STUDENT,TRAINER}|


## 3. Lấy được user, roles và nhiều cột khác

```sql
SELECT du.id, du.name, result.enum_roles FROM demo.users du INNER JOIN
(SELECT u.id, array_agg(r."name") enum_roles FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
GROUP BY u.id) result ON du.id = result.id
```

|id|name|enum_roles|
|--|----|----------|
|2G0gfhsK|Rachel Cremin|{SALE,AUTHOR,EDITOR}|
|W4hmrqRO|Sofia Keebler|{EDITOR,STUDENT,ADMIN}|
|PeZcWXcb|Howard Wilderman|{EDITOR,EMPLOYER,ADMIN}|
|s8mTS_NC|Aditya Fadel|{TRAINER,EMPLOYER,AUTHOR}|

## 4. Tìm kiếm user by id
Khi không dùng cột array việc join bảng khá là dài dòng
- Cách 1:
  ```sql
  EXPLAIN ANALYZE SELECT du.id, du.name, result.enum_roles FROM demo.users du 
  INNER JOIN
  (SELECT u.id, array_agg(r."name") enum_roles FROM 
  demo.users u 
  INNER JOIN demo.user_role ur ON u.id = ur.user_id 
  INNER JOIN demo.roles r ON ur.role_id = r.id
  GROUP BY u.id) result ON du.id = result.id
  WHERE du.id = 'W4hmrqRO'
  ```
- Cách 2:
  ```sql
  -- Lệnh này tham số trong cùng nhìn hơi khó
  EXPLAIN ANALYZE SELECT du.id, du.name, result.enum_roles FROM demo.users du 
  INNER JOIN
  (SELECT u.id, array_agg(r."name") enum_roles FROM 
  demo.users u 
  INNER JOIN demo.user_role ur ON u.id = ur.user_id 
  INNER JOIN demo.roles r ON ur.role_id = r.id
  WHERE u.id = 'W4hmrqRO'
  GROUP BY u.id) result ON du.id = result.id
  ```


Ngược lại dùng sẵn cột array. Câu lệnh SELECT đơn giản vãi !
```sql
EXPLAIN ANALYZE SELECT * FROM demo.users u ;
```
Tốc độ nhanh hơn 10 lần.

## 5. Chọn user theo role nào đó

### 5.1 Sử dụng lệnh JOIN
```sql
SELECT u."name", r.id, r."name" FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
WHERE r.id IN 2
```

|name|id|name|
|----|--|----|
|Sofia Keebler|2|STUDENT|
|Alyson Dicki|2|STUDENT|
|Royal Volkman|2|STUDENT|
|Olen Kerluke|2|STUDENT|
|Federico Sporer|2|STUDENT|
|Ole Turner|2|STUDENT|
|Susanna Grady|2|STUDENT|
|Giuseppe Hauck|2|STUDENT|
|Vivianne Schimmel|2|STUDENT|

## 5.2 Sử dụng truy vấn trong array
```sql
SELECT * FROM users u WHERE 2=ANY(int_roles)
```
|id|name|int_roles|enum_roles|
|--|----|---------|----------|
|W4hmrqRO|Sofia Keebler|{7,2,1}|{EDITOR,STUDENT,ADMIN}|
|ZimSGk4J|Alyson Dicki|{2,7,5}|{STUDENT,EDITOR,EMPLOYER}|
|sg_qMoJP|Royal Volkman|{2,1,6}|{STUDENT,ADMIN,AUTHOR}|
|nZOIcaKf|Olen Kerluke|{3,2,5}|{TRAINER,STUDENT,EMPLOYER}|
|KSO1gcEW|Federico Sporer|{2,6,5}|{STUDENT,AUTHOR,EMPLOYER}|
|o2FtoQCV|Ole Turner|{2,8,4}|{STUDENT,MAINTAINER,SALE}|
|BpoglBOh|Susanna Grady|{3,2,6}|{TRAINER,STUDENT,AUTHOR}|
|_61Lq2Vr|Giuseppe Hauck|{5,6,2}|{EMPLOYER,AUTHOR,STUDENT}|

## 6. Liệt kê danh sách user có đúng role ROLE `2` và `3`

## 6.1 Nếu dùng cột array thì quá đơn giản
```sql
SELECT * FROM users u WHERE 2=ANY(int_roles) and 3=ANY(int_roles)
```

|id|name|int_roles|enum_roles|
|--|----|---------|----------|
|nZOIcaKf|Olen Kerluke|{3,2,5}|{TRAINER,STUDENT,EMPLOYER}|
|BpoglBOh|Susanna Grady|{3,2,6}|{TRAINER,STUDENT,AUTHOR}|
|_goFSoLK|Uriah Russel|{3,6,2}|{TRAINER,AUTHOR,STUDENT}|
|onWWgLXw|Emilia Spinka|{7,2,3}|{EDITOR,STUDENT,TRAINER}|
|9fVbfM1A|Gregoria Mraz|{1,3,2}|{ADMIN,TRAINER,STUDENT}|
|M2XStLps|Larue Harvey|{2,3,4}|{STUDENT,TRAINER,SALE}|
|zwrKD5HT|Thora Stoltenberg|{3,6,2}|{TRAINER,AUTHOR,STUDENT}|
|TyRStF1z|Melvin McLaughlin|{2,5,3}|{STUDENT,EMPLOYER,TRAINER}|

## 6.2 Còn join bảng thì tôi chưa nghĩ ra !
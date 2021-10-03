SELECT u."name", r.id, r."name" FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id



SELECT u."name", array_agg(r."name") FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
GROUP BY u."name";


-- Lấy được cả id, name, roles
SELECT du.id, du.name, result.enum_roles FROM demo.users du INNER JOIN
(SELECT u.id, array_agg(r."name") enum_roles FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
GROUP BY u.id) result ON du.id = result.id


-- Hai lệnh này chạy gần như sau
EXPLAIN ANALYZE SELECT du.id, du.name, result.enum_roles FROM demo.users du 
INNER JOIN
(SELECT u.id, array_agg(r."name") enum_roles FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
GROUP BY u.id) result ON du.id = result.id
WHERE du.id = 'W4hmrqRO'

-- Lệnh này tham số trong cùng nhìn hơi khó
EXPLAIN ANALYZE SELECT du.id, du.name, result.enum_roles FROM demo.users du 
INNER JOIN
(SELECT u.id, array_agg(r."name") enum_roles FROM 
demo.users u 
INNER JOIN demo.user_role ur ON u.id = ur.user_id 
INNER JOIN demo.roles r ON ur.role_id = r.id
WHERE u.id = 'W4hmrqRO'
GROUP BY u.id) result ON du.id = result.id


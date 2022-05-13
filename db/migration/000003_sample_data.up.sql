INSERT INTO 
    "warehouses" ("name")
VALUES 
    ('main'),
    ('north'),
    ('west'),
    ('south');

INSERT INTO 
    "items" ("title","price","WarehouseID")
VALUES 
    ('apple', 89.5, lastval()),
    ('banana', 77, lastval()),
    ('orange', 120, lastval()),
    ('tomato', 300, lastval());
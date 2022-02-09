CREATE TABLE IF NOT EXISTS PERMISSIONS 
(   id              serial PRIMARY key, 
    name      CHARACTER VARYING(30) NOT NULL ,
    description         TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS permissions_employees
(   permissions_id            INTEGER REFERENCES PERMISSIONS(id) ON DELETE CASCADE NOT NULL  ,
    employee_id            INTEGER REFERENCES EMPLOYEE(id) ON DELETE CASCADE NOT NULL     
);

INSERT INTO PERMISSIONS (name, description) VALUES
('read_user','ability to read a user'), 
('creat_user','ability to create a user'),
('delete_user','ability to delete a user'),
('update_user','ability to update a user'),
('read_hotel','ability to read a hotel'), 
('creat_hotel','ability to create a hotel'),
('delete_hotel','ability to delete a hotel'),
('update_hotel','ability to update a hotel');

INSERT INTO permissions_employees (permissions_id, employee_id) VALUES 
(1,3),
(2,3),
(3,3),
(4,3),
(5,3),
(6,3),
(7,3),
(8,3);
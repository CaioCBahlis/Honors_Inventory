INSERT INTO locations (room_name, building_type) VALUES
  ('Honors Warehouse','Warehouse'),
  ('Honors 401A',    'Classroom'),
  ('Honors Kitchen', 'Lab'),
  ('Office 1',       'Office'),
  ('HON 305E',       'Classroom'),
  ('HON 418E',       'Classroom'),
  ('Office 2',       'Office'),
  ('HON 411B',       'Classroom');

INSERT INTO audit(message) VALUES
  ('Rocky Created the Database'),
  ('Rocky Added 50 New Items');


INSERT INTO equipment (model, equipment_type, equipment_status, location_id) VALUES

  ('Dell XPS 13',     'Laptop',   'Working',           1),
  ('MacBook Pro 16',  'Laptop',   'Needs Maintenance', 5),


  ('HP EliteDesk 800','Desktop',  'Working',           3),
  ('Dell OptiPlex 7090','Desktop','Broken',            7),


  ('iPad Air (5th)',  'Tablet',   'Missing',           2),
  ('Samsung Galaxy Tab S8','Tablet','Working',        8),


  ('Samsung 24"',     'Monitor',  'Working',           4),
  ('LG UltraFine 27"','Monitor',  'Needs Maintenance', 6),


  ('Logitech K380',   'Keyboard', 'Working',           1),


  ('Logitech M185',   'Mouse',    'Broken',            3),


  ('Canon PIXMA TS5320','Printer','Working',           4),


  ('Epson PowerLite X49','Projector','Needs Maintenance',2);


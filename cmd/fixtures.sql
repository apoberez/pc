-- WRITE CATEGORIES
TRUNCATE category;
INSERT INTO category (uid, title, slug, image) VALUES
  ('b902c4b5-e820-43ff-84f7-9d4bf11d497f', 'CPU', 'cpu', '/static/images/nav-cpu.png'),
  ('0222a037-361b-42fb-bc25-d1d4eee0dc44', 'CPU Coooler', 'cpu-cooler', '/static/images/nav-cpucooler.png'),
  ('db2b65b9-66a4-4e4a-ad88-2db669a3b51a', 'Motherboard', 'motherboard', '/static/images/nav-motherboard.png'),
  ('7e8ffbdb-bebb-4d7a-80df-60864344ed96', 'Memory', 'memory', '/static/images/nav-memory.png'),
  ('c23c83ed-7f57-4cc3-81c4-67d387c0c7e1', 'Video Card', 'video-card', '/static/images/nav-videocard.png'),
  ('2ef7b4a9-bac8-465d-a797-79d9029df6d1', 'Storage', 'storage', '/static/images/nav-ssd.png'),
  ('f7d6dab5-9821-43ba-b2a4-a881e51b5ec1', 'Power', 'power', '/static/images/nav-powersupply.png'),
  ('21e2e756-a26f-4669-9e72-e43920e5ef8f', 'Case', 'case', '/static/images/nav-case.png');


TRUNCATE manufacturer CASCADE;
INSERT INTO manufacturer (uid, title) VALUES
  ('8d15f587-3717-4213-b748-3f8cd4fdb9cc', 'AMD'),
  ('687e634a-1836-4e30-8219-119d84ff7767', 'Intel');

INSERT INTO product (uid, title, manufacturer, specification) VALUES
  (
    'f07561cd-4fa8-4346-bc7f-009810519d25',
    'AMD FX-8350 4.0GHz 8-Core Processor',
    '8d15f587-3717-4213-b748-3f8cd4fdb9cc',
    '{
      "socket": "AM3+",
      "operating frequency": 4000,
      "cores": 8,
      "rating": 4
    }'
  ), (
    '7d30a2b3-1660-4c73-8db2-7c63ae41ca5e',
    'Intel Core i5-6600 3.3GHz Quad-Core Processor',
    '687e634a-1836-4e30-8219-119d84ff7767',
    '{
      "socket": "LGA1151",
      "operating frequency": 3900,
      "cores": 4,
      "rating": 5
    }'
  );
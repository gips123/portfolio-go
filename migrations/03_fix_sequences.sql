-- Fix PostgreSQL Sequences
-- Jalankan script ini setelah seed data untuk memperbaiki sequence
-- Sequence harus di-set ke nilai yang lebih tinggi dari ID terakhir yang ada

-- Fix projects sequence
SELECT setval('projects_id_seq', (SELECT MAX(id) FROM projects));

-- Fix project_images sequence
SELECT setval('project_images_id_seq', (SELECT MAX(id) FROM project_images));

-- Fix skill_categories sequence
SELECT setval('skill_categories_id_seq', (SELECT MAX(id) FROM skill_categories));

-- Fix contact_data sequence
SELECT setval('contact_data_id_seq', (SELECT MAX(id) FROM contact_data));

-- Fix contact_messages sequence (jika ada)
SELECT setval('contact_messages_id_seq', COALESCE((SELECT MAX(id) FROM contact_messages), 1));

-- Verify sequences
SELECT 'projects' as table_name, last_value as current_sequence_value 
FROM projects_id_seq
UNION ALL
SELECT 'project_images', last_value FROM project_images_id_seq
UNION ALL
SELECT 'skill_categories', last_value FROM skill_categories_id_seq
UNION ALL
SELECT 'contact_data', last_value FROM contact_data_id_seq;


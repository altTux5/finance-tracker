INSERT INTO categories (id, name)
VALUES 
    (gen_random_uuid(), 'Food'),
    (gen_random_uuid(), 'Snacks'),
    (gen_random_uuid(), 'Groceries'),
    (gen_random_uuid(), 'Transport'),
    (gen_random_uuid(), 'Entertainment'),
    (gen_random_uuid(), 'Utilities'),
    (gen_random_uuid(), 'Shopping'),
    (gen_random_uuid(), 'Healthcare'),
    (gen_random_uuid(), 'Other')
ON CONFLICT (name) DO NOTHING;
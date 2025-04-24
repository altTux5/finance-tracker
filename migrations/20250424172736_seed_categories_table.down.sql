DELETE FROM categories
WHERE name IN (
    'Food',
    'Snacks',
    'Groceries',
    'Transport',
    'Entertainment',
    'Utilities',
    'Shopping',
    'Healthcare',
    'Other'
);
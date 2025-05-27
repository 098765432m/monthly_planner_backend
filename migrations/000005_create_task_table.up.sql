DROP TABLE IF EXISTS tasks;

-- Drop the enum type if it exists
DROP TYPE IF EXISTS task_status_enum;

-- Create the enum type
CREATE TYPE task_status_enum AS ENUM('NOT_DONE', 'DONE');

-- Create the tasks table that uses the enum type
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    description TEXT,
    status task_status_enum NOT NULL DEFAULT 'NOT_DONE',
    time_start TIME,
    time_end TIME,
    day_id UUID NOT NULL REFERENCES days(id) ON DELETE CASCADE,
    task_category_id UUID REFERENCES task_categories(id) ON DELETE SET NULL
);

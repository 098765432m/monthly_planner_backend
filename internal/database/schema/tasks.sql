CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    description TEXT,
    status ENUM('NOT_DONE', 'DONE') NOT NULL DEFAULT 'NOT_DONE',
    time_start TIME,
    time_end TIME,
    day_id UUID NOT NULL REFERENCES days(id) ON DELETE CASCADE
)
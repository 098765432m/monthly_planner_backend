CREATE TABLE days (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date DATE NOT NULL UNIQUE,
    day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
    month_id UUID NOT NULL REFERENCES months(id) ON DELETE CASCADE
)
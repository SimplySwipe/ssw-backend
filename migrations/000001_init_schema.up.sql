-- USERS
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    google_id TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    phone TEXT,
    photo_url TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- USER PREFERENCES
CREATE TABLE user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    regions TEXT[],
    job_types TEXT[],
    salary_min INTEGER,
    remote_only BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- JOBS
CREATE TABLE jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    location TEXT,
    salary_min INTEGER,
    salary_max INTEGER,
    job_type TEXT,
    company TEXT,
    url TEXT,
    posted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- SAVED JOBS
CREATE TABLE saved_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    job_id UUID REFERENCES jobs(id) ON DELETE CASCADE,
    saved_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(user_id, job_id)
);

-- SWIPES
CREATE TABLE swipes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    job_id UUID REFERENCES jobs(id) ON DELETE CASCADE,
    direction TEXT CHECK (direction IN ('like', 'dislike')),
    swiped_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(user_id, job_id)
);

-- APPLICATIONS
CREATE TABLE applications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    job_id UUID REFERENCES jobs(id) ON DELETE CASCADE,
    status TEXT,
    applied_at TIMESTAMPTZ DEFAULT now()
);

-- USER ACTIVITY
CREATE TABLE user_activity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    action TEXT,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);
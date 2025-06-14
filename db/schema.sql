-- Create task_status table
CREATE TABLE task_status (
                             status_id SERIAL PRIMARY KEY,
                             status_name VARCHAR(50) UNIQUE NOT NULL
);

-- Create task table
CREATE TABLE task (
                      task_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      parent_task_id UUID REFERENCES task(task_id),
                      task_order INTEGER NOT NULL,
                      name VARCHAR(255) NOT NULL,
                      description TEXT,
                      start_time TIMESTAMP,
                      end_time TIMESTAMP,
                      status_id INTEGER NOT NULL REFERENCES task_status(status_id),
                      updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_user VARCHAR(100) NOT NULL,
                      last_modified_process VARCHAR(100),
                      last_modified_app VARCHAR(100),
                      last_request_id VARCHAR(100) NOT NULL,
                      last_action_id VARCHAR(100),
                      CONSTRAINT unique_task_order_per_parent UNIQUE (parent_task_id, task_order)
);

-- Create task_type table
CREATE TABLE task_type (
                           task_type_id SERIAL PRIMARY KEY,
                           task_id UUID UNIQUE NOT NULL REFERENCES task(task_id),
                           type VARCHAR(50) NOT NULL
);

-- Create redjade_task table
CREATE TABLE redjade_task (
                              redjade_task_id SERIAL PRIMARY KEY,
                              task_type_id INTEGER NOT NULL REFERENCES task_type(task_type_id),
                              redjade_link TEXT
);

-- Create sams_survey_task table
CREATE TABLE sams_survey_task (
                                  task_id SERIAL PRIMARY KEY,
                                  task_type_id INTEGER NOT NULL REFERENCES task_type(task_type_id),
                                  survey_id TEXT
);

-- Create task_keyword table
CREATE TABLE task_keyword (
                              keyword_id SERIAL PRIMARY KEY,
                              task_id UUID NOT NULL REFERENCES task(task_id),
                              keyword VARCHAR(25) NOT NULL
);

-- Create indexes for better performance
CREATE INDEX idx_task_parent_id ON task(parent_task_id);
CREATE INDEX idx_task_status_id ON task(status_id);
CREATE INDEX idx_task_updated_ts ON task(updated_ts);
CREATE INDEX idx_task_type_task_id ON task_type(task_id);
CREATE INDEX idx_task_keyword_task_id ON task_keyword(task_id);
CREATE INDEX idx_task_keyword_keyword ON task_keyword(keyword);

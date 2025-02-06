DROP TABLE IF EXISTS rating_cards;

CREATE TABLE IF NOT EXISTS rating_cards (
                                            id VARCHAR(36) PRIMARY KEY,
    question TEXT NOT NULL,
    category VARCHAR(255) NOT NULL,
    order_id INT NOT NULL
    );

CREATE TABLE Rating (
                        Id INT AUTO_INCREMENT PRIMARY KEY,
                        UserId INT NOT NULL,
                        TimeStamp DATETIME NOT NULL,
                        RatingCardId INT NOT NULL,
                        RatingCandidate INT NOT NULL,
                        TextResponseCandidate TEXT,
                        RatingEmployer VARCHAR(255),
                        TextResponseEmployer TEXT
);

CREATE TABLE TotalRatings (
                              RatingCardId INT PRIMARY KEY,
                              AverageRating FLOAT,
                              TotalSubmissions INT
);

START TRANSACTION;

-- Insert questions for Performance category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('1', 'Would a customer see you as a senior?', 'Performance', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('2', 'Do you deliver high-quality work consistently?', 'Performance', 2);

-- Insert questions for Technical Skillset category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('3', 'How do you rate your proficiency related to your role description?', 'Technical Skillset', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('4', 'Do you keep your technical knowledge up to date?', 'Technical Skillset', 2);

-- Insert questions for Technical Predispositions category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('5', 'How comfortable are you in solving complex technical problems?', 'Technical Predispositions', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('6', 'Do you approach tasks with a problem-solving mindset?', 'Technical Predispositions', 2);

-- Insert questions for Sales category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('7', 'How do you approach sales challenges?', 'Sales', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('8', 'Can you handle objections effectively during a sales pitch?', 'Sales', 2);

-- Insert questions for Recruiting category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('9', 'How do you assess candidates for a job role?', 'Recruiting', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('10', 'Do you build strong relationships with candidates?', 'Recruiting', 2);

-- Insert questions for Teamwork category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('11', 'How well do you collaborate with your team?', 'Teamwork', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('12', 'Can you resolve conflicts within the team?', 'Teamwork', 2);

-- Insert questions for Coaching category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('13', 'Do you provide constructive feedback to your peers?', 'Coaching', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('14', 'Are you an effective mentor to others?', 'Coaching', 2);

-- Insert questions for Prodyna Insights category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('15', 'How do you contribute to Prodyna Insights?', 'Prodyna Insights', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('16', 'Do you actively share your knowledge within the team?', 'Prodyna Insights', 2);

-- Insert questions for Overall category
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('17', 'Do you consistently meet or exceed expectations?', 'Overall', 1);
INSERT INTO rating_cards (id, question, category, order_id) VALUES ('18', 'How would you rate your overall performance?', 'Overall', 2);

COMMIT;
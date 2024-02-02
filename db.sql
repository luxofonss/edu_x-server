create type gender as ENUM('male', 'female', 'other');
create type role as ENUM('student','teacher','parent');
create type level as ENUM('beginer','elementary','intermediate','upper_intermediate','advanced','proficiency');
create type course_info as ENUM('intend', 'who','requirement','welcome_msg','congrat_msg');
create type question_level as ENUM('easy','medium','hard','super_hard');
create type assignment_type as ENUM('mid','final','regular');
create type currency as ENUM('vnd', 'usd');
create type question_type as ENUM('single_choice','multi_choice','short_answer','long_answers', 'super_question');
create type learner_type as ENUM('student','pupil','other');
create type supervisor_relationship as ENUM('mother', 'father', 'sister','brother','other');
create type auth_types as ENUM('email_pwd', 'facebook', 'google');
create type enroll_status as enum ('PENDING', 'ACTIVE', 'INACTIVE');
create type feedback_types as enum ('GOOD','BAD');

alter type level add value 'beginner'

alter type question_type add value 'multi_choice';


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP table IF exists "users" CASCADE;
CREATE TABLE "users" (
  "id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" varchar(30) unique not NULL,
  "email" varchar(100) unique not NULL,
  "verified" boolean default false,
  "phone_number" varchar(20) unique not NULL,
  "first_name" varchar(60) not null,
  "last_name" varchar(60) not null,
  "dob" timestamp not NULL,
  "gender" gender not NULL,
  "role" role not NULL,
  "avatar" varchar(100),
  "deleted_at" timestamp default NULL,
  "updated_at" timestamp not null default CURRENT_TIMESTAMP,
  "created_at" timestamp not null default CURRENT_TIMESTAMP
);

DROP table IF exists "addresses" cascade;
CREATE TABLE "addresses" (
  "id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid not null, 
  "province" varchar(50) not null,
  "district" varchar(50) not null,
  "ward" varchar(50) not null,
  "house_number" varchar(50),
  "country" varchar(100),
  "updated_at" timestamp default CURRENT_TIMESTAMP,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "deleted_at" timestamp default null,
  
  constraint fk_user_id foreign key (user_id) references users(id) on delete cascade
);

drop table if exists "auths" cascade;
create table "auths" (
	"user_id" uuid not null PRIMARY KEY,
	"auth_type" auth_types not null default 'email_pwd',
	"salt" text not null,
	"email" varchar(100) unique not null,
	"password" text not null,
	"service_id" varchar(50),
 	"updated_at" timestamp default CURRENT_TIMESTAMP,
    "created_at" timestamp default CURRENT_TIMESTAMP,
    "deleted_at" timestamp default null,
    
    constraint fk_user_id foreign key (user_id) references users(id) on delete cascade
);

drop table if exists "learner_info";
create table "learner_info" (
	"id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
	"user_id" uuid not null,
	"type" learner_type,
	"grade" int,
	"school" varchar(300),
	"deleted_at" timestamp default NULL,
    "updated_at" timestamp not null default CURRENT_TIMESTAMP,
    "created_at" timestamp not null default CURRENT_TIMESTAMP,
    
    CONSTRAINT check_student_info CHECK (
        ("type" <> 'student') OR ("type" = 'student' AND "grade" IS NOT NULL AND "school" IS NOT NULL)
    ),
    constraint fk_user_id foreign key (user_id) references users(id) on delete cascade
);


drop table if exists "learner_supervisors";
create table "learner_supervisors" (
	"learner_id" uuid not null,
	"supervisor_id" uuid not null,
	"relationship" supervisor_relationship not null,
	"deleted_at" timestamp default NULL,
    "updated_at" timestamp not null default CURRENT_TIMESTAMP,
    "created_at" timestamp not null default CURRENT_TIMESTAMP,
    
    foreign key ("learner_id") references users ("id"),
    foreign key ("supervisor_id") references users ("id"),
    primary key ("learner_id", "supervisor_id")
);

DROP table IF exists "edu_qualifications" CASCADE;
create table edu_qualifications (
	"id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" varchar(200) not null,
	"description" varchar(500) not null,
	"updated_at" timestamp default CURRENT_TIMESTAMP,
  	"created_at" timestamp default CURRENT_TIMESTAMP,
  	"deleted_at" timestamp default NULL
);

INSERT INTO edu_qualifications ("id", "name", "description")
VALUES 
    (DEFAULT, 'High School Diploma', 'Typically obtained after completing secondary education.'),
    (DEFAULT, 'Associate''s Degree', 'A two-year undergraduate degree often offered by community colleges.'),
    (DEFAULT, 'Bachelor''s Degree', 'A four-year undergraduate degree earned at a college or university.'),
    (DEFAULT, 'Master''s Degree', 'An advanced degree usually requiring 1-2 years beyond a bachelor''s degree.'),
    (DEFAULT, 'Doctorate or Ph.D.', 'The highest academic degree typically requiring several years of research and study.'),
    (DEFAULT, 'Diploma in Technical Education', 'A technical diploma providing specialized training in technical fields.'),
    (DEFAULT, 'Certificate Programs', 'Short-term programs offering specialized training or certification.'),
    (DEFAULT, 'Professional Certifications', 'Certifications indicating expertise in a specific profession or field.'),
    (DEFAULT, 'Postgraduate Diploma', 'An advanced diploma following completion of a bachelor''s degree.'),
    (DEFAULT, 'Specialized Training Programs', 'Focused training programs in specific fields or industries.'),
    (DEFAULT, 'Honorary Degrees', 'Degrees conferred honorarily for significant contributions to society or fields, not earned through academic study.'),
    (DEFAULT, 'Vocational Education', 'Training focusing on specific skills and trades.'),
    (DEFAULT, 'Online Course Certifications', 'Certifications obtained through completion of online courses.'),
    (DEFAULT, 'Continuing Education Credits', 'Credits earned for professional development or ongoing education.'),
    (DEFAULT, 'Apprenticeships', 'Training programs combining on-the-job training with classroom instruction.'),
    (DEFAULT, 'Language Proficiency Certificates', 'Certifications demonstrating proficiency in a specific language.');

    
drop table if exists "school_positions" cascade;
create table school_positions(
	"id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
	"name" varchar(50) not null,
	"description" varchar(300),
	"updated_at" timestamp default CURRENT_TIMESTAMP,
  	"created_at" timestamp default CURRENT_TIMESTAMP,
  	"deleted_at" timestamp default NULL
);

insert into school_positions ("id", "name","description")
VALUES 
    (DEFAULT,'Teacher', 'Responsible for teaching and guiding students in the learning process.'),
    (DEFAULT,'Principal', 'Head responsible for managing academic and administrative activities.'),
    (DEFAULT,'Content Manager', 'Responsible for developing and managing educational content.'),
    (DEFAULT,'Technical Support', 'Provides technical assistance for the online learning platform.');

drop table if exists "teacher_infos" cascade;
create table teacher_infos (
    "id" uuid not null PRIMARY KEY DEFAULT uuid_generate_v4(),
	"user_id" uuid not null,
--	"edu_qualification" uuid not null,
	"biography" text not null,
	"updated_at" timestamp default CURRENT_TIMESTAMP,
  	"created_at" timestamp default CURRENT_TIMESTAMP,
  	"deleted_at" timestamp default null,
  	CONSTRAINT fk_teacher_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE cascade,
--  	constraint fk_edu_qualification foreign key (edu_qualification) references edu_qualifications(id) 
);

alter table teacher_infos add column edu_qualification varchar(400);

drop table if exists "schools" cascade;
create table "schools" (
	"id" uuid not null PRIMARY KEY,
	"name" varchar(200) not null,
	"description" text not null,
	"address" uuid not null references addresses(id),
	"logo" varchar(200),
	"background" varchar(200),
	"verified" boolean default false,
	"deleted_at" timestamp default NULL,
  	"created_at" timestamp default CURRENT_TIMESTAMP,
  	"updated_at" timestamp default CURRENT_TIMESTAMP
);

drop table if exists "teacher_school" cascade;
create table "teacher_school" (
	"teacher_id" uuid not null references users("id") on delete cascade,
	"school_id" uuid not null references schools("id") on delete cascade,
	"start_work_time" timestamp,
	"salary" int,
	"position" uuid not null references school_positions(id),
	"deleted_at" timestamp default NULL,
  	"created_at" timestamp default CURRENT_TIMESTAMP,
  	"updated_at" timestamp default CURRENT_TIMESTAMP
);

DROP table IF exists "subjects" CASCADE;
CREATE TABLE "subjects" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "name" varchar(50) not null,
  "description" text,
  "thumbnail_url" varchar(100),
  "deleted_at" timestamp default NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "updated_at" timestamp default CURRENT_TIMESTAMP
);

insert into subjects ("id","name", "description","thumbnail_url")
values (default, 'English', 'English', 'thumbnail');

DROP table IF exists "courses" CASCADE;
CREATE TABLE "courses" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "name" varchar(200) not NULL,
  "description" text not NULL,
  "background_img" text,
  "thumbnail" text,
  "start_date" timestamp,
  "end_date" timestamp,
  "price" decimal(12,2),
  "currency" currency default 'vnd',
  "level" level not NULL,
  "is_verified" boolean default false,
  "subject_id" uuid not NULL,
  "grade" int,
  "code" varchar(10) not NULL,
  "teacher_id" uuid references users(id) on delete cascade,
  "school_id" uuid references schools(id) on delete cascade,
  "updated_at" timestamp default CURRENT_TIMESTAMP,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "deleted_at" timestamp default null,
  FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id"),
  constraint check_only_one_owner check (
	(teacher_id is not null and school_id is null) or 
	(teacher_id is null and school_id is not null)
  )
);

alter table courses add column thumbnail text;

drop table if exists "user_enroll_course" cascade;
create table "user_enroll_course" (
	"id" uuid not NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid not NULL,
	"course_id" uuid not null,
	"price" decimal(12,2),
	"student_id" int,
	"updated_at" timestamp default CURRENT_TIMESTAMP,
	"created_at" timestamp default CURRENT_TIMESTAMP,
	"deleted_at" timestamp default null,
	PRIMARY KEY ("user_id", "course_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id")
);
alter table user_enroll_course add column status enroll_status default 'PENDING';

DROP table IF exists "course_infos" CASCADE;
CREATE TABLE "course_infos" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "course_id" uuid not NULL references courses(id) on delete cascade,
  "content" text not NULL,
  "type" course_info not NULL,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "updated_at" timestamp default CURRENT_TIMESTAMP
);

DROP table IF exists "sections" CASCADE;
CREATE TABLE "sections" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "name" varchar(100) not NULL,
  "description" text not NULL,
  "course_id" uuid not null references courses(id) on delete CASCADE,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "updated_at" timestamp default CURRENT_TIMESTAMP
);

alter table sections drop column description;
alter table sections add column description text;


DROP table IF exists "lectures" CASCADE;
CREATE TABLE "lectures" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "name" varchar(100) not NULL,
  "description" text not NULL,	
  "background" varchar(100),
  "video_url" varchar(100),
  "section_id" uuid references sections(id) on delete cascade,
  "updated_at" timestamp default CURRENT_TIMESTAMP,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "deleted_at" timestamp default NULL
);

alter table lectures drop column description;
alter table lectures add column description text;

DROP table IF exists "documents" CASCADE;
CREATE TABLE "documents" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "title" varchar(200) not NULL,
  "url" varchar(200) not null,
  "description" text not null,
  "user_id" uuid references users(id) on delete cascade,
  "school_id" uuid references schools(id) on delete cascade,
  "updated_at" timestamp default CURRENT_TIMESTAMP,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  "deleted_at" timestamp default NULL,
  constraint check_only_one_owner check (
	(user_id is not null and school_id is null) or 
	(user_id is null and school_id is not null)
  )
);

DROP table IF exists "document_placement" CASCADE;
CREATE TABLE document_placement (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "document_id" uuid NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
  "course_id" uuid references courses(id) on delete cascade,
  "lecture_id" uuid references lectures(id) on delete cascade,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT check_placement CHECK ((course_id IS NOT NULL AND lecture_id IS NULL) OR (course_id IS NULL AND lecture_id IS NOT NULL))
);

DROP table IF exists "assignments" CASCADE;
CREATE TABLE "assignments" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "title" varchar(200) not null,
  "description" text not null,
  "total_point" integer not null,
  "multiple_attempt" boolean,
  "type" assignment_type not null,
  "start_time" timestamp,
  "end_time" timestamp,
  "placement_id" uuid not null,
  "subject_id" uuid references subjects(id),
  "teacher_id" uuid references users(id) on delete cascade,
  "school_id" uuid references schools(id) on delete cascade,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
alter table assignments add column time int; 

--
--DROP table IF exists "assignment_placement";
--CREATE TABLE "assignment_placement" (
--  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
--  "assignment_id" uuid references assignments("id") not null,
--  "type" assignment_type,
--  "start_time" timestamp,
--  "multiple_attempts" boolean,
--  "end_time" timestamp,
--  "course_id" uuid references courses(id) on delete cascade,
--  "lecture_id" uuid references lectures(id) on delete cascade,
--  "section_id" uuid references sections(id) on delete cascade,
--  "deleted_at" timestamp default NULL,
--  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
--  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
--  CONSTRAINT check_placement CHECK (
--  (course_id IS NOT NULL and section_id IS NULL and lecture_id IS NULL ) OR 
--  (course_id IS NOT NULL and section_id IS NOT NULL and lecture_id IS NULL ) OR
--  (course_id IS NOT NULL AND section_id IS NOT NULL and lecture_id IS NOT NULL)
--  )
--);

DROP table IF exists "assignment_attempts" CASCADE;
CREATE TABLE "assignment_attempts" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "assignment_id" uuid references assignments(id),
  "user_id" uuid not null references users(id),
  "point" integer,
  "teacher_comment" text,
  "finished_at" timestamp default NULL,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  constraint check_assignment check (assignment_id is not null)
);
alter table assignment_attempts add column assignment_time_millis integer;

DROP table IF exists "questions" CASCADE;
CREATE TABLE "questions" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "parent_id" uuid references questions(id),
  "assignment_id" uuid references assignments(id),
  "title" text not null,
  "image" json,
  "audio_url" text,
  "point" integer,
  "type" question_type not NULL,
  "level" question_level not NULL,
  "answer_explain" text,
  "teacher_id" uuid references users(id) on delete cascade,
  "school_id" uuid references schools(id) on delete cascade,
  "subject_id" uuid references subjects(id),
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  constraint check_only_one_owner check (
	(teacher_id is not null and school_id is null) or 
	(teacher_id is null and school_id is not null)
  )
);

alter table questions add column point integer;

DROP table IF exists "question_choices" CASCADE;
CREATE TABLE "question_choices" (
  "id" uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
  "content" text not null,
  "order" int not null,
  "is_correct" boolean not null,
  "question_id" uuid not null references questions(id),
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
alter table question_choices drop column answer_explain;

DROP table IF exists "question_answers" CASCADE;
create table "question_answers" (
	id uuid not NULL PRIMARY key DEFAULT uuid_generate_v4(),
    user_id uuid not null references users(id),
    question_id uuid not null references questions(id),
    assignment_attempt_id uuid not null references assignment_attempts("id"),
    selected_option_id uuid references question_choices(id), -- Store the ID of the selected option for multiple/single choice questions
    text_answer TEXT, -- For short/long answer questions
    constraint check_only_one_answer check (
    	(selected_option_id is not null and text_answer is null)  or 
    	(selected_option_id is null and text_answer is not null)
    ),
    "deleted_at" timestamp default NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
alter table question_answers  add constraint unique_question_attempt_user 
        UNIQUE (question_id, assignment_attempt_id, user_id);
alter table question_answers drop constraint check_only_one_answer;
alter table question_answers add column score int; -- score for long answer

-- feedback for long answer
DROP table IF exists "feedbacks" CASCADE;
CREATE TABLE "feedbacks" (
  "id" uuid not null primary key,
  "message" text not null,
  "user_id" uuid not null references users(id),
  "type" feedback_types not null,
  "question_answer_id" uuid  not null references question_answers(id),
  "feedback_id" uuid references feedbacks(id) default NULL, -- target feedback
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

drop table if exists "correct_answers" cascade;
create table "correct_answers" (
	question_id uuid not NULL PRIMARY key references questions("id") on delete cascade DEFAULT uuid_generate_v4(),
	correct_text_answer text,
	answer_explain text,
	"deleted_at" timestamp default NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

--End question

DROP table IF exists "conversations" CASCADE;
CREATE TABLE "conversations" (
  "id" uuid not null primary key DEFAULT uuid_generate_v4(),
  "content" text not null,
  "owner_id" uuid not null references users(id) on delete cascade,
  "placement_id" uuid,
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP table IF exists "comments" CASCADE;
CREATE TABLE "comments" (
  "id" uuid not null primary key default uuid_generate_v4(),
  "message" text not null,
  "user_id" uuid not null references users(id),
  "conversation_id" uuid  not null references conversations(id),
  "deleted_at" timestamp default NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

drop table if exists "reacts" cascade;
create table "reacts" (
	"id" uuid not null primary key default uuid_generate_v4(),
	"text" varchar(50) not null,
	deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP table IF exists "react_conversations" CASCADE;
create table "react_conversations" (
	user_id uuid not null default uuid_generate_v4(),
	conversation_id uuid not null,
	type uuid not null, -- The type of reaction
	deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    foreign key ("type") references "reacts" ("id"),
    foreign key ("user_id") references "users" ("id"),
    foreign key ("conversation_id") references "conversations" ("id"),
    primary key ("user_id", "conversation_id")
);

DROP table IF exists "react_comments" CASCADE;
create table "react_comments" (
	user_id uuid not null primary key default uuid_generate_v4(),
	comment_id uuid not null,
	type uuid not null, -- The type of reaction
	deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    foreign key ("type") references "reacts" ("id"),
    foreign key ("user_id") references "users" ("id"),
    foreign key ("comment_id") references "comments" ("id")
);